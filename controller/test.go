package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

// Test ...
func Test(c echo.Context) error {

	return c.String(http.StatusOK, "GET Test")
}
