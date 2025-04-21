package handlers

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"screen-time-limiter/internal/domain/model"
	"screen-time-limiter/internal/utils"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPushCurrentTimeHandler_ServeHTTP(t *testing.T) {
	t.Parallel()

	limitID := int32(1)
	userID := int32(2)
	now := time.Now()
	year, month, day := now.Date()
	date, _ := time.Parse(time.DateOnly, fmt.Sprintf("%d-%s-%s", year, utils.FixDate(int(month)), utils.FixDate(day)))
	bodyRaw := fmt.Sprintf(`{"limit_id": %d, "time_amount": "%s"}`, limitID, "1h")

	tests := []struct {
		name   string
		status int
		body   io.Reader
		before func(ctx context.Context, limitHistoryRepoMock *LimitHistoryRepoMock, limitRepoMock *LimitRepoMock, timerMock *TimerMock)
	}{
		{
			name:   "no body",
			status: http.StatusBadRequest,
		},
		{
			name:   "validation error",
			status: http.StatusBadRequest,
			body:   strings.NewReader("{\"limit_id\":0,\"time_amount\":\"\"}"),
		},
		{
			name:   "GetByLimitID is empty",
			status: http.StatusNotFound,
			body:   strings.NewReader(bodyRaw),
			before: func(ctx context.Context, limitHistoryRepoMock *LimitHistoryRepoMock, limitRepoMock *LimitRepoMock, timerMock *TimerMock) {
				limitRepoMock.GetByLimitIDMock.Expect(ctx, limitID).
					Return(nil, nil)
			},
		},
		{
			name:   "GetByLimitID error",
			status: http.StatusInternalServerError,
			body:   strings.NewReader(bodyRaw),
			before: func(ctx context.Context, limitHistoryRepoMock *LimitHistoryRepoMock, limitRepoMock *LimitRepoMock, timerMock *TimerMock) {
				limitRepoMock.GetByLimitIDMock.Expect(ctx, limitID).
					Return(nil, errors.New("error"))
			},
		},
		{
			name:   "Push error",
			status: http.StatusInternalServerError,
			body:   strings.NewReader(bodyRaw),
			before: func(ctx context.Context, limitHistoryRepoMock *LimitHistoryRepoMock, limitRepoMock *LimitRepoMock, timerMock *TimerMock) {
				limitRepoMock.GetByLimitIDMock.Expect(ctx, limitID).
					Return(&model.Limit{
						UserID: userID,
					}, nil)
				timerMock.NowMock.Expect().Return(now)
				limitHistoryRepoMock.PushMock.Expect(ctx, model.LimitHistory{
					LimitID:    limitID,
					TimeAmount: 3600,
					SentAt:     now,
					LimitDate:  date,
				}).Return(errors.New("error"))
			},
		},
		{
			name:   "ok",
			status: http.StatusOK,
			body:   strings.NewReader(bodyRaw),
			before: func(ctx context.Context, limitHistoryRepoMock *LimitHistoryRepoMock, limitRepoMock *LimitRepoMock, timerMock *TimerMock) {
				limitRepoMock.GetByLimitIDMock.Expect(ctx, limitID).
					Return(&model.Limit{
						UserID: userID,
					}, nil)
				timerMock.NowMock.Return(now)
				limitHistoryRepoMock.PushMock.Expect(ctx, model.LimitHistory{
					LimitID:    limitID,
					TimeAmount: 3600,
					SentAt:     now,
					LimitDate:  date,
				}).Return(nil)
				limitHistoryRepoMock.SumMock.Expect(ctx, limitID, now).
					Return(100, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			limitRepoMock := NewLimitRepoMock(t)
			limitHistoryRepoMock := NewLimitHistoryRepoMock(t)
			timerMock := NewTimerMock(t)
			handler := NewPushCurrentTimeHandler(limitHistoryRepoMock, limitRepoMock, timerMock)

			a := assert.New(t)

			req, err := http.NewRequest(http.MethodPost, "/v1/limit/history", tt.body)
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
