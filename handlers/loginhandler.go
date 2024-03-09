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
	return c.Render(http.StatusOK, "index.html", nil)
}

func LoginHandler(c echo.Context) error {
	if c.Request().Method == http.MethodPost {
		username := c.FormValue("username")
		password := c.FormValue("password")

		// Check if the username and password match using GORM
		if AuthenticateUser(username, password) {
			// Set the cookie
			c.SetCookie(&http.Cookie{
				Name:     "username",
				Value:    username,
				Expires:  time.Now().Add(time.Minute * 60),
				Secure:   true,
				HttpOnly: true,
				SameSite: 1, //houd de cookie alleen op de site, speelt niet naar derden
			})

			// Redirect to /home
			return c.Redirect(http.StatusSeeOther, "/home")
		}
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid username or password")
	}

	return c.NoContent(http.StatusBadRequest)
}

func AuthenticateUser(username, password string) bool {
	var user models.User
	// Find the user by username and password using GORM
	err := database.DB().Where("username = ? AND password = ?", username, password).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return false
	} else if err != nil {
		log.Fatal(err)
	}

	return true
}
