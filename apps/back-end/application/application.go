package application

import (
	"FindMyDosen/config"
	"FindMyDosen/model/entity"
	"FindMyDosen/modules/auth"
	"FindMyDosen/modules/university"
	"FindMyDosen/modules/user"
	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"time"
)

func ApplicationDelegate() {
	e := echo.New()
	// Middleware
	setupMiddlewares(e)
	jwtMiddleware := getJWTMiddleware()
	e.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	}, jwtMiddleware)
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"msg": "ok",
		})
	})

	v1 := e.Group("/v1")
	auth.AuthController(v1)
	university.UniController(v1, jwtMiddleware)
	user.UserController(v1, jwtMiddleware)

	port := config.GetServerPort()
	e.Logger.Fatal(e.Start(port))
}

func setupMiddlewares(e *echo.Echo) {
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		ErrorMessage: "Request Timeout!",
		Timeout:      30 * time.Second,
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Gzip())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
	e.Use(middleware.Secure())
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
