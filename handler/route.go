package handler

import (
	"github.com/Elias2389/api-rest/middleware"
	"github.com/labstack/echo/v4"
)

// Router of person
func RoutePerson(e *echo.Echo, storage Storage) {
	h := newPerson(storage)

	person := e.Group("/v1/persons")
	person.Use(middleware.Authentication)

	person.POST("", h.create)
	person.GET("", h.getAll)
	person.GET("/:id", h.getByID)
}

// RouteLogin .
func RouteLogin(e *echo.Echo, storage Storage) {
	h := newLogin(storage)

	e.POST("/v1/login", h.login)
}
