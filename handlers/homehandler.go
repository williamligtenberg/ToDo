package handlers

import (
	"ToDoApplication/database"
	"ToDoApplication/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func HomeHandler(c echo.Context) error {
	// Cookies naar de HTML sturen.
	cookies := c.Cookies()
	data := map[string]interface{}{
		"Cookies": cookies,
	}

	// HTML pagina renderen inclusief cookies erbij.
	return c.Render(http.StatusOK, "home.jet.html", data)

}

func CreateToDoHandler(c echo.Context) error {
	//CreateToDo functie aanroepen en title en description uit de form halen.
	database.CreateToDo(models.ToDo{
		Title:       c.FormValue("title"),
		Description: c.FormValue("description"),
	})
	// Doorsturen naar de home pagina.
	return c.Redirect(http.StatusFound, "/home")
}
