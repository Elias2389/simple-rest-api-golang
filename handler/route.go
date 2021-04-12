package handler

import (
	"github.com/Elias2389/api-rest/middleware"
	"net/http"
)

func RoutePerson(mux *http.ServeMux, storage Storage) {
	h := newPerson(storage)
	mux.HandleFunc("/v1/persons/create", middleware.Log(middleware.Authentication(h.create)))
	mux.HandleFunc("/v1/persons/", middleware.Log(h.getAll))
	mux.HandleFunc("/v1/persons/get", h.getByID)
}

// RouteLogin .
func RouteLogin(mux *http.ServeMux, storage Storage) {
	h := newLogin(storage)

	mux.HandleFunc("/v1/login", h.login)
}
