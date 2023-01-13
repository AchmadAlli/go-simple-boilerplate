package routes

import (
	"github.com/achmadAlli/go-simple-boilerplate/adapters/api/v1/injector"
	"github.com/labstack/echo/v4"
)

func Routes(g *echo.Group) {
	NewCommonRoute(
		injector.ProvideSwaggerHandler(),
	).Register(g)
}
