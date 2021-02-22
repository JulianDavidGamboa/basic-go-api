package handler

import (
	"github.com/JulianDavidGamboa/basic-go-api/middleware"
	"github.com/labstack/echo/v4"
)

// RoutePerson .
func RoutePerson(e *echo.Echo, storage Storage) {
	h := newPerson(storage)
	person := e.Group("/v1/persons")
	person.Use(middleware.Authentication)

	person.POST("", h.create)
	person.PUT("/:id", h.update)
	e.DELETE("/:id", h.delete)
	person.GET("", h.getAll)

}

// RouteLogin .
func RouteLogin(e *echo.Echo, storage Storage) {
	h := newLogin(storage)

	e.POST("/v1/login", h.login)
}
