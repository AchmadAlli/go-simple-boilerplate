package api

import (
	"net/http"

	"github.com/achmadAlli/go-simple-boilerplate/adapters/api/utils"
	v1 "github.com/achmadAlli/go-simple-boilerplate/adapters/api/v1/routes"
	"github.com/labstack/echo/v4"
)

func RegisterRoute(e *echo.Echo) {
	api := e.Group("/api")
	gv1 := api.Group("/v1")

	api.GET("/check", func(c echo.Context) error {
		return c.JSON(http.StatusOK, utils.NewResponse(nil, "server is up"))
	})

	api.Static("/docs", "docs/swagger-ui")

	v1.Routes(gv1)
}
