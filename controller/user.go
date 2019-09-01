package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

// UsuarioTestePOST ...
func UsuarioTestePOST(c echo.Context) error {

	return c.String(http.StatusOK, "User route")
}

// UsuarioTesteGET ...
func UsuarioTesteGET(c echo.Context) error {

	return c.String(http.StatusOK, "GET User route")
}

// UsuarioTestePUT ...
func UsuarioTestePUT(c echo.Context) error {

	return c.String(http.StatusOK, "PUT User route")
}
