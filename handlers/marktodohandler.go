package handlers

import (
	"ToDoApplication/database"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

func MarkToDoHandler(c echo.Context) error {
	// todoID maken.
	todoID, err := strconv.Atoi(c.Param("id"))
	// Error afhandelen.
	if err != nil {
		log.Printf("Failed to get ID")
		return err
	}

	// ToDo done in database updaten.
	err = database.MarkToDoAsDone(todoID)
	// Error afhandelen.
	if err != nil {
		log.Printf("Failed to update ToDo status: %v", err)
		return err
	}
	return c.Redirect(http.StatusFound, "/todos")
}
