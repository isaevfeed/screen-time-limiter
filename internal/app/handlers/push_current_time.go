package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"screen-time-limiter/internal/domain/model"
	"screen-time-limiter/internal/domain/request"
	responseDomain "screen-time-limiter/internal/domain/response"
	"screen-time-limiter/internal/utils"
	"screen-time-limiter/internal/utils/response"
	"time"
)

type (
	PushCurrentTimeHandler struct {
		limitHistoryRepo limitHistoryRepo
		limitRepo        limitRepo

		timer timer
	}
)

func NewPushCurrentTimeHandler(limitHistoryRepo limitHistoryRepo, limitRepo limitRepo, timer timer) *PushCurrentTimeHandler {
	return &PushCurrentTimeHandler{
		limitHistoryRepo: limitHistoryRepo,
		limitRepo:        limitRepo,
		timer:            timer,
	}
}

func (h *PushCurrentTimeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	var pushCurrentTimeReq request.PushCurrentTime
	err = json.Unmarshal(body, &pushCurrentTimeReq)
	if err != nil {
		response.MakeResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	errs := validatePushCurrentTimeReq(pushCurrentTimeReq)
	if errs.Err() != nil {
		errs.MakeResponse(w)
		return
	}

	ctx := r.Context()
	limit, err := h.limitRepo.GetByLimitID(ctx, pushCurrentTimeReq.LimitID)
	if err != nil {
		response.MakeResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	if limit == nil {
		response.MakeResponse(w, http.StatusNotFound, "limit not found")
		return
	}

	now := h.timer.Now()
	year, month, day := now.Date()
	date, err := time.Parse(time.DateOnly, fmt.Sprintf("%d-%s-%s", year, utils.FixDate(int(month)), utils.FixDate(day)))
	if err != nil {
		response.MakeResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	amount, err := time.ParseDuration(pushCurrentTimeReq.TimeAmount)
	if err != nil {
		response.MakeResponse(w, http.StatusInternalServerError, amountDurationErr)
		return
	}

	err = h.limitHistoryRepo.Push(ctx, model.LimitHistory{
		LimitID:    pushCurrentTimeReq.LimitID,
		TimeAmount: int32(amount.Seconds()),
		SentAt:     now,
		LimitDate:  date,
	})
	if err != nil {
		response.MakeResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	historyTotalSum, err := h.limitHistoryRepo.Sum(ctx, pushCurrentTimeReq.LimitID, now)
	if err != nil {
		response.MakeResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.MakeRespWithData(w, http.StatusOK, responseDomain.PushCurrentTimeResp{
		Data: responseDomain.PushCurrentTimeRespData{
			Message:     "Текущее время учтено",
			Expired:     limit.Expired(int32(historyTotalSum)),
			TimeBalance: limit.ApplyBalance(int32(historyTotalSum)),
		},
	})
}

func validatePushCurrentTimeReq(req request.PushCurrentTime) *response.ValidationError {
	errs := response.NewValidationError()

	if req.LimitID == 0 {
		errs.AddDetail("limit_id", notEmptyErr)
	}
	if req.TimeAmount == "" {
		errs.AddDetail("time_amount", notEmptyErr)
	}

	return errs
}
