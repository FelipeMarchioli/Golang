package routes

import (
	"teste/controller"

	"github.com/labstack/echo"
)

// TestRoute ...
func TestRoute(testRoute *echo.Group) {

	testRoute.GET("/", controller.Test)
}
