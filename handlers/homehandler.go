package handlers

import (
	"ToDoApplication/database"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ToDosHandler(c echo.Context) error {
	// Retrieve todos from the database
	todos := database.RetrieveToDo(c)

	// Data naar de HTML sturen.
	cookies := c.Cookies()
	data := map[string]interface{}{
		"Cookies": cookies,
		"Todos":   todos,
	}

	// HTML pagina renderen inclusief cookies erbij.
	return c.Render(http.StatusOK, "todos.html", data)

}
