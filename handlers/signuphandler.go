package handlers

import (
	"ToDoApplication/database"
	"ToDoApplication/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func SignupPage(c echo.Context) error {
	return c.Render(http.StatusOK, "signup.html", nil)
}

func CreateUserHandler(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	user := models.User{
		Username: username,
		Password: password,
	}

	err := database.CreateUser(user)
	if err != nil {
		// Error afhandelen.
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create user")
	}

	if AuthenticateUser(username, password) {
		c.SetCookie(&http.Cookie{
			Name:     "user",
			Value:    username,
			Expires:  time.Now().Add(time.Minute * 60),
			Secure:   true,
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
		})
		return c.Redirect(http.StatusSeeOther, "/todos")
	}
	return echo.NewHTTPError(http.StatusUnauthorized, "Failed to authenticate user")
}
