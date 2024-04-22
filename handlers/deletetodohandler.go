package handlers

import (
	"ToDoApplication/database"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func DeleteToDoHandler(c echo.Context) error {
	fmt.Println("DeleteToDoHandler called") // Check if the handler is being called
	todoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// Handle error
		return err
	}
	fmt.Println("Todo ID:", todoID) // Check if todoID is correctly extracted
	database.DeleteToDo(todoID)     // Call DeleteToDo function passing the todoID
	return c.Redirect(http.StatusFound, "/todos")
}
