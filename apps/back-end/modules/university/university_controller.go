package university

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func UniController(e *echo.Group, jwtMiddleware echo.MiddlewareFunc) {
	e.GET("/university", func(c echo.Context) error {
		return c.String(http.StatusOK, "uni")
	}, jwtMiddleware)
}
