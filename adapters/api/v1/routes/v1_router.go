package routes

import (
	"github.com/labstack/echo/v4"
)

func Routes(g *echo.Group) {
	NewCommonRoute().Register(g)
}
