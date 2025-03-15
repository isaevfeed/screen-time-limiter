package handlers

import (
	"net/http"
	responseDomain "screen-time-limiter/internal/domain/response"
	"screen-time-limiter/internal/utils/response"
	"strconv"
)

type (
	GetTotalHistoryTimesHandler struct {
		limitHistoryRepo limitHistoryRepo
		limitRepo        limitRepo

		timer timer
	}
)

func NewGetTotalHistoryTimesHandler(limitHistoryRepo limitHistoryRepo, limitRepo limitRepo, timer timer) *GetTotalHistoryTimesHandler {
	return &GetTotalHistoryTimesHandler{
		limitHistoryRepo: limitHistoryRepo,
		limitRepo:        limitRepo,
		timer:            timer,
	}
}

func (h *GetTotalHistoryTimesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	limtiIDRaw := r.URL.Query().Get("limit_id")

	errs := response.NewValidationError()
	if limtiIDRaw == "" {
		errs.AddDetail("limit_id", notEmptyErr)
	}
	if errs.Err() != nil {
		errs.MakeResponse(w)
		return
	}

	limitID, err := strconv.Atoi(limtiIDRaw)
	if err != nil {
		response.MakeResponse(w, http.StatusInternalServerError, wrongLimitIDErr)
		return
	}

	ctx := r.Context()
	sum, err := h.limitHistoryRepo.Sum(ctx, int32(limitID), h.timer.Now())
	if err != nil {
		response.MakeResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	limit, err := h.limitRepo.GetByLimitID(ctx, int32(limitID))
	if err != nil {
		response.MakeResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.MakeRespWithData(w, http.StatusOK, responseDomain.GetTotalLimitTimesResp{
		Sum:     sum,
		Expired: limit.Expired(int32(sum)),
	})
}
