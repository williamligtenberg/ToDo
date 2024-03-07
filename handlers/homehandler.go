package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HomePage(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}
