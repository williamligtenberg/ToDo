package handlers

import (
	"ToDoApplication/database"
	"ToDoApplication/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func SignupPage(c echo.Context) error {
	return c.Render(http.StatusOK, "signup.html", nil)
}

func CreateUserHandler(c echo.Context) error {
	database.CreateUser(models.User{
		Username: c.FormValue("username"),
		Password: c.FormValue("password"),
	})
	return c.Redirect(http.StatusFound, "/")
}
