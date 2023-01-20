package routes

import (
	"github.com/achmadAlli/go-simple-boilerplate/adapters/api/v1/handler"
	"github.com/labstack/echo/v4"
)

type userRoute struct {
	userHandler *handler.UserHandler
}

func NewUserRoute(handler *handler.UserHandler) *userRoute {
	return &userRoute{
		userHandler: handler,
	}
}

func (r *userRoute) Register(g *echo.Group) {
	g.POST("/users", r.userHandler.Store)
	g.GET("/users", r.userHandler.Fetch)
	g.GET("/users/:id", r.userHandler.Find)
	g.PUT("/user/:id", r.userHandler.Update)
	g.DELETE("/users/:id", r.userHandler.Destroy)
}
