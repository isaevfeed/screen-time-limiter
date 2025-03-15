package response

import (
	"encoding/json"
	"net/http"
)

type (
	RespWithData[T any] struct {
		Status int `json:"status"`
		Data   T   `json:"data"`
	}
)

func MakeRespWithData[T any](w http.ResponseWriter, status int, data T) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(&RespWithData[T]{
		Status: status,
		Data:   data,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
