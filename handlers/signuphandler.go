package handlers

import (
	"ToDoApplication/database"
	"ToDoApplication/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func SignupPage(c echo.Context) error {
	// Signup pagina renderen.
	return c.Render(http.StatusOK, "signup.html", nil)
}

func CreateUserHandler(c echo.Context) error {
	// Username en password uit de form halen.
	username := c.FormValue("username")
	password := c.FormValue("password")

	// User variable maken na het model van User.
	user := models.User{
		Username: username,
		Password: password,
	}

	// User aanmaken.
	err := database.CreateUser(user)
	if err != nil {
		// Error afhandelen.
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create user")
	}

	if AuthenticateUser(username, password) {
		// Cookie instellen.
		c.SetCookie(&http.Cookie{
			Name:     "user",
			Value:    username,
			Expires:  time.Now().Add(time.Minute * 60),
			Secure:   true,
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
		})

		// Gebruiker doorsturen naar ToDos.
		return c.Redirect(http.StatusSeeOther, "/todos")
	}

	// Error afhandelen.
	return echo.NewHTTPError(http.StatusUnauthorized, "Failed to authenticate user")
}
