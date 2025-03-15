package response

import (
	"encoding/json"
	"net/http"
)

type (
	Response struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}
)

func MakeResponse(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(&Response{
		Status:  status,
		Message: message,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
