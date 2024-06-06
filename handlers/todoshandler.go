package handlers

import (
	"ToDoApplication/database"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ToDosHandler(c echo.Context) error {
	todos, err := database.RetrieveToDo(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve ToDo items")
	}

	cookies := c.Cookies()
	data := map[string]interface{}{
		"Cookies": cookies,
		"Todos":   todos,
	}
	return c.Render(http.StatusOK, "todos.html", data)
}
