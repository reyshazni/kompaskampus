package application

import (
	"FindMyDosen/config"
	"FindMyDosen/modules"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"time"
)

func ApplicationDelegate() {
	e := echo.New()
	// Middleware
	setupMiddlewares(e)
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"msg": "ok",
		})
	})
	modules.MainController(e)
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
