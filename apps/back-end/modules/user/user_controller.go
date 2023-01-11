package user

import "github.com/labstack/echo/v4"

func UserController(e *echo.Group, jwtMiddleware echo.MiddlewareFunc) {
	group := e.Group("/user")
	group.GET("", handleGetUserInformation, jwtMiddleware)
	group.PUT("", handleUpdateUserInformation, jwtMiddleware)
	group.PUT("/password", handleChangePassword, jwtMiddleware)
}
