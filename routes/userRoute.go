package routes

import (
	"teste/controller"

	"github.com/labstack/echo"
)

// UserRoute ...
func UserRoute(g *echo.Group) {

	g.POST("/", controller.UsuarioTestePOST)

	g.GET("/get", controller.UsuarioTesteGET)

	g.PUT("/put", controller.UsuarioTestePUT)
}
