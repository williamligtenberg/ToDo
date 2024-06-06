package handlers

import (
	"ToDoApplication/database"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

func MarkToDoHandler(c echo.Context) error {
	todoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Failed to get ID")
		return err
	}

	err = database.MarkToDoAsDone(todoID)
	if err != nil {
		log.Printf("Failed to update ToDo status: %v", err)
		return err
	}
	return c.Redirect(http.StatusFound, "/todos")
}
