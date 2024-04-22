package handlers

import (
	"ToDoApplication/database"
	"ToDoApplication/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func NewToDoPageHandler(c echo.Context) error {
	// NewToDo pagina renderen.
	// Cookies naar de HTML sturen.
	cookies := c.Cookies()
	data := map[string]interface{}{
		"Cookies": cookies,
	}
	return c.Render(http.StatusOK, "newtodo.html", data)
}

func CreateToDoHandler(c echo.Context) error {
	//Check for user who's logged in
	userCookie, err := c.Cookie("user")
	if err != nil {
		fmt.Println("Error while getting username cookie")
	}
	username := userCookie.Value
	//CreateToDo functie aanroepen, title en description uit de form halen en username van de cookie.
	database.CreateToDo(models.ToDo{
		Title:       c.FormValue("title"),
		Description: c.FormValue("description"),
		Creator:     username,
	})
	// Doorsturen naar de home pagina.
	return c.Redirect(http.StatusFound, "/todos")
}
