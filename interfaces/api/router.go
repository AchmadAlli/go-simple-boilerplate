package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterRoute(e *echo.Echo) {
	api := e.Group("/api")
	v1 := api.Group("/v1")

	api.GET("/check", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "server is running")
	})

	v1.GET("/check", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "api v1 is up")
	})
}
