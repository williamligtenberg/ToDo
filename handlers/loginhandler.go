package handlers

import (
	"ToDoApplication/database"
	"ToDoApplication/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

func LoginPage(c echo.Context) error {
	// Kijken of de username cookie er is.
	_, err := c.Cookie("username")
	if err != nil {
		// De loginpagina weer laten zien omdat de gebruiker niet juist is ingelogd.
		return c.Render(http.StatusOK, "login.html", nil)
	}

	// Alle cookies ophalen en gebruiker naar todos.html sturen.
	cookies := c.Cookies()
	data := map[string]interface{}{
		"Cookies": cookies,
	}

	c.Render(http.StatusOK, "todos.html", data)
	return nil
}

func LoginHandler(c echo.Context) error {
	if c.Request().Method == http.MethodPost {
		username := c.FormValue("username")
		password := c.FormValue("password")

		if AuthenticateUser(username, password) {
			c.SetCookie(&http.Cookie{
				Name:     "user",
				Value:    username,
				Expires:  time.Now().Add(time.Minute * 60),
				Secure:   true,
				HttpOnly: true,
				SameSite: 1,
			})

			return c.Redirect(http.StatusSeeOther, "/todos")
		}
		return c.Redirect(http.StatusSeeOther, "/login?error=incorrect")
	}
	return c.NoContent(http.StatusBadRequest)
}

func AuthenticateUser(username, password string) bool {
	var user models.User
	err := database.DB().Where("username = ? AND password = ?", username, password).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return false
	} else if err != nil {
		log.Fatal(err)
	}
	return true
}
