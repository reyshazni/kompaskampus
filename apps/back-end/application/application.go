package application

import (
	"FindMyDosen/config"
	"FindMyDosen/modules/auth"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ApplicationDelegate() {
	e := echo.New()
	v1 := e.Group("/v1")
	auth.AuthController(v1)
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"msg": "ok",
		})
	})
	port := config.GetServerPort()
	e.Logger.Fatal(e.Start(port))
}
