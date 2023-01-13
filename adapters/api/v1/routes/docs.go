package routes

import (
	"github.com/achmadAlli/go-simple-boilerplate/adapters/api/v1/handler"
	"github.com/labstack/echo/v4"
)

type CommonRoute struct {
	swaggerHandler *handler.SwaggerHandler
}

func NewCommonRoute(
	swaggerHandler *handler.SwaggerHandler,
) *CommonRoute {
	return &CommonRoute{
		swaggerHandler: swaggerHandler,
	}
}

func (r CommonRoute) Register(g *echo.Group) {
	g.GET("/swagger", r.swaggerHandler.SpecDocs)
}
