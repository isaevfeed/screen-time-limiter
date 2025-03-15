package handlers

import (
	"context"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"screen-time-limiter/internal/domain/model"
	"strings"
	"testing"
)

func TestAddUser_ServeHTTP(t *testing.T) {
	t.Parallel()

	firstName := "Michael"
	lastName := "Isaev"
	bodyRaw := fmt.Sprintf(`{"first_name": "%s", "last_name": "%s"}`, firstName, lastName)

	tests := []struct {
		name   string
		status int
		body   io.Reader
		before func(ctx context.Context, userRepoMock *UserRepoMock)
	}{
		{
			name:   "no body",
			status: http.StatusBadRequest,
		},
		{
			name:   "validation error",
			status: http.StatusBadRequest,
			body:   strings.NewReader("{\"first_name\":\"\",\"last_name\":\"\"}"),
		},
		{
			name:   "Add error",
			status: http.StatusInternalServerError,
			body:   strings.NewReader(bodyRaw),
			before: func(ctx context.Context, userRepoMock *UserRepoMock) {
				userRepoMock.AddMock.Expect(ctx, model.User{
					FirstName: firstName,
					LastName:  lastName,
				}).Return(errors.New("error"))
			},
		},
		{
			name:   "ok",
			status: http.StatusCreated,
			body:   strings.NewReader(bodyRaw),
			before: func(ctx context.Context, userRepoMock *UserRepoMock) {
				userRepoMock.AddMock.Expect(ctx, model.User{
					FirstName: firstName,
					LastName:  lastName,
				}).Return(nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userRepoMock := NewUserRepoMock(t)
			handler := NewAddUserHandler(userRepoMock)

			a := assert.New(t)

			req, err := http.NewRequest(http.MethodPost, "/v1/limit", tt.body)
			a.NoError(err)

			// Создаем ResponseRecorder для записи ответа
			rr := httptest.NewRecorder()

			if tt.before != nil {
				tt.before(req.Context(), userRepoMock)
			}

			// Вызываем ServeHTTP
			handler.ServeHTTP(rr, req)

			a.Equal(tt.status, rr.Code)
		})
	}
}
