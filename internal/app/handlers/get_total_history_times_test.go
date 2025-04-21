package handlers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"screen-time-limiter/internal/domain/model"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetTotalHistoryTimesHandler_ServeHTTP(t *testing.T) {
	t.Parallel()

	limitID := "1"
	now := time.Now()

	tests := []struct {
		name    string
		status  int
		limitID string
		before  func(ctx context.Context, limitHistoryRepoMock *LimitHistoryRepoMock, limitRepoMock *LimitRepoMock, timerMock *TimerMock)
	}{
		{
			name:    "validation error",
			status:  http.StatusBadRequest,
			limitID: "",
		},
		{
			name:    "ok",
			status:  http.StatusInternalServerError,
			limitID: "invalid",
		},
		{
			name:    "Sum error",
			status:  http.StatusInternalServerError,
			limitID: limitID,
			before: func(ctx context.Context, limitHistoryRepoMock *LimitHistoryRepoMock, limitRepoMock *LimitRepoMock, timerMock *TimerMock) {
				timerMock.NowMock.Return(now)

				limitIDInt, _ := strconv.Atoi(limitID)

				limitHistoryRepoMock.SumMock.Expect(ctx, int32(limitIDInt), now).
					Return(0, errors.New("Sum error"))
			},
		},
		{
			name:    "ok",
			status:  http.StatusOK,
			limitID: limitID,
			before: func(ctx context.Context, limitHistoryRepoMock *LimitHistoryRepoMock, limitRepoMock *LimitRepoMock, timerMock *TimerMock) {
				timerMock.NowMock.Return(now)

				limitIDInt, _ := strconv.Atoi(limitID)

				limitHistoryRepoMock.SumMock.Expect(ctx, int32(limitIDInt), now).
					Return(100, nil)
				limitRepoMock.GetByLimitIDMock.Expect(ctx, int32(limitIDInt)).
					Return(&model.Limit{}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			limitHistoryRepoMock := NewLimitHistoryRepoMock(t)
			limitRepoMock := NewLimitRepoMock(t)
			timerMock := NewTimerMock(t)
			handler := NewGetTotalHistoryTimesHandler(limitHistoryRepoMock, limitRepoMock, timerMock)

			a := assert.New(t)

			req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/v1/limit/history/total?limit_id=%s", tt.limitID), nil)
			a.NoError(err)

			// Создаем ResponseRecorder для записи ответа
			rr := httptest.NewRecorder()

			if tt.before != nil {
				tt.before(req.Context(), limitHistoryRepoMock, limitRepoMock, timerMock)
			}

			// Вызываем ServeHTTP
			handler.ServeHTTP(rr, req)

			a.Equal(tt.status, rr.Code)
		})
	}
}
