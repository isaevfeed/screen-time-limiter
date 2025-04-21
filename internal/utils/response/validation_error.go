package response

import (
	"encoding/json"
	"errors"
	"net/http"
)

type (
	ValidationError struct {
		Error   string             `json:"error"`
		Details []validationDetail `json:"details"`
	}
	validationDetail struct {
		Filed   string `json:"filed"`
		Message string `json:"message"`
	}
)

func NewValidationError() *ValidationError {
	return &ValidationError{
		Error:   "Ошибка валидации",
		Details: []validationDetail{},
	}
}

func (e *ValidationError) AddDetail(field string, message string) {
	e.Details = append(e.Details, validationDetail{
		Filed:   field,
		Message: message,
	})
}

func (e *ValidationError) MakeResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	err := json.NewEncoder(w).Encode(e)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (e *ValidationError) Err() error {
	if len(e.Details) > 0 {
		return errors.New("Ошибка валидации")
	}

	return nil
}
