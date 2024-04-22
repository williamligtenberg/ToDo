package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func LogoutHandler(c echo.Context) error {
	// De cookies "resetten" zodat je uitgelogd word.
	c.SetCookie(&http.Cookie{
		Name:     "user",
		Value:    "",
		MaxAge:   -1,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
	//Terugsturen naar de inlogpagina.
	return c.Redirect(http.StatusSeeOther, "/login")
}
