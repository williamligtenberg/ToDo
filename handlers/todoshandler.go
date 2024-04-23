package handlers

import (
	"ToDoApplication/database"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ToDosHandler(c echo.Context) error {
	// Todos uit de database halen.
	todos, err := database.RetrieveToDo(c)
	// Error afhandelen.
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve ToDo items")
	}

	// Data naar de HTML sturen.
	cookies := c.Cookies()
	data := map[string]interface{}{
		"Cookies": cookies,
		"Todos":   todos,
	}

	// HTML pagina renderen inclusief data erbij.
	return c.Render(http.StatusOK, "todos.html", data)
}
