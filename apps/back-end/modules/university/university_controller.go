package university

import (
	"github.com/labstack/echo/v4"
)

func UniController(e *echo.Group, jwtMiddleware echo.MiddlewareFunc) {
	e.GET("/university", handleGetUniversityData, jwtMiddleware)
}
