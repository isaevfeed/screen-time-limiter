package handlers

import "net/http"

type (
	AddUser struct{}
)

func NewAddUserHandler() *AddUser {
	return &AddUser{}
}

func (h *AddUser) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
