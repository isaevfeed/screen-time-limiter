package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"screen-time-limiter/internal/domain/model"
	"screen-time-limiter/internal/domain/request"
	"screen-time-limiter/internal/utils/response"
)

type (
	AddUser struct {
		userRepo userRepo
	}
)

func NewAddUserHandler(userRepo userRepo) *AddUser {
	return &AddUser{
		userRepo: userRepo,
	}
}

func (h *AddUser) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	var addUserReq request.AddUser
	err = json.Unmarshal(body, &addUserReq)
	if err != nil {
		response.MakeResponse(w, http.StatusInternalServerError, err.Error())
	}

	errs := validateAddUserReq(addUserReq)
	if errs.Err() != nil {
		errs.MakeResponse(w)
		return
	}

	err = h.userRepo.Add(r.Context(), model.User{
		FirstName: addUserReq.FirstName,
		LastName:  addUserReq.LastName,
	})
	if err != nil {
		response.MakeResponse(w, http.StatusInternalServerError, err.Error())
	}

	response.MakeResponse(w, http.StatusCreated, "Пользователь успешно создан")
}

func validateAddUserReq(addUserReq request.AddUser) *response.ValidationError {
	validErrs := response.NewValidationError()

	if addUserReq.FirstName == "" {
		validErrs.AddDetail("first_name", notEmptyErr)
	}
	if addUserReq.LastName == "" {
		validErrs.AddDetail("last_name", notEmptyErr)
	}

	return validErrs
}
