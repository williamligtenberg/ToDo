package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func LogoutHandler(c echo.Context) error {
	// De cookies "resetten" zodat je uitgelogd word.
	cookie := new(http.Cookie)
	cookie.Name = "user"
	cookie.Value = ""
	cookie.MaxAge = -1
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteStrictMode
	cookie.Path = "/"

	c.SetCookie(cookie)
	//Terugsturen naar de inlogpagina.
	return c.Redirect(http.StatusSeeOther, "/")
}
