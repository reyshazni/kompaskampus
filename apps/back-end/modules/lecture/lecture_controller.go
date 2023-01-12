package lecture

import "github.com/labstack/echo/v4"

func LectureController(e *echo.Group, jwtMiddleware echo.MiddlewareFunc) {
	group := e.Group("/lecture")
	group.POST("", handleAddLecture, jwtMiddleware)
}
