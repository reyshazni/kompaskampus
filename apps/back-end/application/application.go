package application

import (
	"FindMyDosen/config"
	"FindMyDosen/model/entity"
	"FindMyDosen/modules/auth"
	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func ApplicationDelegate() {
	e := echo.New()
	v1 := e.Group("/v1")
	auth.AuthController(v1)
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	}, getJWTMiddleware())
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"msg": "ok",
		})
	})
	port := config.GetServerPort()
	e.Logger.Fatal(e.Start(port))
}

func getJWTMiddleware() echo.MiddlewareFunc {
	jwtConf := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(entity.JwtClaims)
		},
		SigningKey: []byte(config.GetJWTSecret()),
	}
	return echojwt.WithConfig(jwtConf)
}
