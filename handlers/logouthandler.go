package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func LogoutHandler(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "user"
	cookie.Value = ""
	cookie.MaxAge = -1
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteStrictMode
	cookie.Path = "/" // Set the same path as when the cookie was set

	c.SetCookie(cookie)

	return c.Redirect(http.StatusSeeOther, "/")
}
