package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"screen-time-limiter/internal/domain/model"
	"screen-time-limiter/internal/domain/request"
	"screen-time-limiter/internal/utils/response"
	"time"
)

type (
	AddTimeLimitHandler struct {
		limitRepo limitRepo
	}
)

func NewAddTimeLimitHandler(limitRepo limitRepo) *AddTimeLimitHandler {
	return &AddTimeLimitHandler{
		limitRepo: limitRepo,
	}
}

func (h *AddTimeLimitHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		response.MakeResponse(w, http.StatusBadRequest, "body is required")
		return
	}

	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.MakeResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	var addLimitReq request.AddLimit
	err = json.Unmarshal(body, &addLimitReq)
	if err != nil {
		response.MakeResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	amount, err := time.ParseDuration(addLimitReq.Amount)
	if err != nil {
		response.MakeResponse(w, http.StatusBadRequest, amountDurationErr)
		return
	}

	errs := validateAddLimitReq(addLimitReq)
	if errs.Err() != nil {
		errs.MakeResponse(w)
	}

	err = h.limitRepo.Create(r.Context(), model.Limit{
		UserID: addLimitReq.UserID,
		Amount: int32(amount.Seconds()),
	})
	if err != nil {
		response.MakeResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.MakeResponse(w, http.StatusCreated, "Лимиты успешно добавлены")
}

func validateAddLimitReq(addLimitReq request.AddLimit) *response.ValidationError {
	validErrs := response.NewValidationError()

	if addLimitReq.UserID == 0 {
		validErrs.AddDetail("user_id", notEmptyErr)
	}
	if addLimitReq.Amount == "" {
		validErrs.AddDetail("amount", notEmptyErr)
	}

	return validErrs
}
