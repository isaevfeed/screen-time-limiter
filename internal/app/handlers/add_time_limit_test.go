package handlers

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"screen-time-limiter/internal/domain/model"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTimeLimitHandler_ServeHTTP(t *testing.T) {
	t.Parallel()

	userID := int32(1)
	amount := int32(3600)
	bodyRaw := "{\"user_id\":1,\"amount\":\"1h\"}"

	tests := []struct {
		name   string
		status int
		body   io.Reader
		before func(ctx context.Context, limitRepoMock *LimitRepoMock)
	}{
		{
			name:   "no body",
			status: http.StatusBadRequest,
		},
		{
			name:   "validation error",
			status: http.StatusBadRequest,
			body:   strings.NewReader("{\"user_id\":0,\"amount\":\"\"}"),
		},
		{
			name:   "Create error",
			status: http.StatusInternalServerError,
			body:   strings.NewReader(bodyRaw),
			before: func(ctx context.Context, limitRepoMock *LimitRepoMock) {
				limitRepoMock.CreateMock.Expect(ctx, model.Limit{
					UserID: userID,
					Amount: amount,
				}).Return(errors.New("error"))
			},
		},
		{
			name:   "ok",
			status: http.StatusCreated,
			body:   strings.NewReader(bodyRaw),
			before: func(ctx context.Context, limitRepoMock *LimitRepoMock) {
				limitRepoMock.CreateMock.Expect(ctx, model.Limit{
					UserID: userID,
					Amount: amount,
				}).Return(nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			limitRepoMock := NewLimitRepoMock(t)
			handler := NewAddTimeLimitHandler(limitRepoMock)

			a := assert.New(t)

			req, err := http.NewRequest(http.MethodPost, "/v1/limit", tt.body)
			a.NoError(err)

			// Создаем ResponseRecorder для записи ответа
			rr := httptest.NewRecorder()

			if tt.before != nil {
				tt.before(req.Context(), limitRepoMock)
			}

			// Вызываем ServeHTTP
			handler.ServeHTTP(rr, req)

			a.Equal(tt.status, rr.Code)
		})
	}
}
