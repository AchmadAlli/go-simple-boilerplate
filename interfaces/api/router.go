package api

import (
	"net/http"

	"github.com/achmadAlli/go-simple-boilerplate/interfaces/api/utils"
	"github.com/labstack/echo/v4"
)

func RegisterRoute(e *echo.Echo) {
	api := e.Group("/api")
	v1 := api.Group("/v1")

	api.GET("/check", func(c echo.Context) error {
		return c.JSON(http.StatusOK, utils.NewResponse(nil, "server is up"))
	})

	v1.GET("/check", func(c echo.Context) error {
		return c.JSON(http.StatusOK, utils.NewResponse(nil, "api v1 is up"))
	})
}
