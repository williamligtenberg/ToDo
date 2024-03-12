package handlers

import (
	"github.com/labstack/echo/v4"
)

func HTTPErrorHandler(err error, c echo.Context) {
	// Error pagina renderen.
	template := "error.html"
	c.Render(echo.ErrNotFound.Code, template, nil)
}
