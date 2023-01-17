package auth

import (
	"github.com/labstack/echo/v4"
)

func AuthController(e *echo.Group) {
	group := e.Group("/auth")
	group.POST("/login", handleLogin)
	group.POST("/register", handleRegister)
	group.POST("/refresh", handleRefresh)
	group.GET("/:code", handleVerifyUser)
}
