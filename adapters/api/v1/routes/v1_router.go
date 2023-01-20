package routes

import (
	"net/http"

	"github.com/achmadAlli/go-simple-boilerplate/adapters/api/utils"
	"github.com/achmadAlli/go-simple-boilerplate/adapters/api/v1/injector"
	"github.com/labstack/echo/v4"
)

func Routes(g *echo.Group) {

	g.GET("/check-health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, utils.NewResponse(nil, "server is up"))
	})

	NewCommonRoute(
		injector.ProvideSwaggerHandler(),
	).Register(g)

	NewUserRoute(
		injector.ProvideUserHandler(),
	).Register(g)
}
