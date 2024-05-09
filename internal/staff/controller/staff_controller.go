package controller

import "net/http"

type StaffController interface {
	Register(w http.ResponseWriter, r *http.Request)
}
