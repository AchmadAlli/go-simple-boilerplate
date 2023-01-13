package routes

import (
	"github.com/achmadAlli/go-simple-boilerplate/adapters/api/v1/handler"
	"github.com/labstack/echo/v4"
)

type CommonRoute struct {
	handler *handler.SwaggerHandler
}

func NewCommonRoute() *CommonRoute {
	return &CommonRoute{
		handler: handler.NewSwaggerHandler(),
	}
}

func (r CommonRoute) Register(g *echo.Group) {
	g.GET("/swagger", r.handler.Swagger)
}
