package routes

import (
	"net/http"

	"github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
)

// Index ...
func Index() *echo.Echo {

	e := echo.New()

	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {

		return c.String(http.StatusOK, "DEU CERTO!!!")
	})

	UserRoute(e.Group("/users"))
	TestRoute(e.Group("/test"))

	return e
}
