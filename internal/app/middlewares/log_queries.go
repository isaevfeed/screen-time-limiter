package middlewares

import (
	"context"
	"fmt"
	"net/http"
)

type (
	LogQueries struct {
		log logger
	}

	logger interface {
		InfoContext(ctx context.Context, msg string, args ...any)
		Error(msg string, args ...any)
	}
)

func NewLogQueries(logger logger) *LogQueries {
	return &LogQueries{
		log: logger,
	}
}

func (l *LogQueries) ServeHTTP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l.log.InfoContext(r.Context(), fmt.Sprintf("%s %s", r.Method, r.URL.Path))

		next.ServeHTTP(w, r)
	})
}
