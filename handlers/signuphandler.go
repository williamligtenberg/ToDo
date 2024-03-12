package handlers

import (
	"ToDoApplication/database"
	"ToDoApplication/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func SignupPage(c echo.Context) error {
	// Signup pagina renderen.
	return c.Render(http.StatusOK, "signup.html", nil)
}

func CreateUserHandler(c echo.Context) error {
	//CreateUser functie aanroepen en username en password uit de form halen.
	database.CreateUser(models.User{
		Username: c.FormValue("username"),
		Password: c.FormValue("password"),
	})
	// Doorsturen naar de login pagina.
	return c.Redirect(http.StatusFound, "/")
}
