package modules

import (
	"FindMyDosen/config"
	"FindMyDosen/model/entity"
	"FindMyDosen/modules/auth"
	"FindMyDosen/modules/lecture"
	"FindMyDosen/modules/university"
	"FindMyDosen/modules/user"
	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func MainController(e *echo.Echo) {
	jwtMiddleware := getJWTMiddleware()
	v1 := e.Group("/v1")
	auth.AuthController(v1)
	university.UniController(v1, jwtMiddleware)
	user.UserController(v1, jwtMiddleware)
	lecture.LectureController(v1, jwtMiddleware)
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
