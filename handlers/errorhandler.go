package handlers

import (
	"github.com/labstack/echo/v4"
)

func HTTPErrorHandler(err error, c echo.Context) {
	// Render error page.
	template := "error.html"
	c.Render(echo.ErrNotFound.Code, template, nil)
}
