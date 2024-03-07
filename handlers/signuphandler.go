package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func SignupPage(c echo.Context) error {
	return c.Render(http.StatusOK, "signup.html", nil)
}
