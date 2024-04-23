package handlers

import (
	"ToDoApplication/database"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

func DeleteToDoHandler(c echo.Context) error {
	// todoID maken.
	todoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Failed to get ID")
		return err
	}
	//De DeleteToDo function aanroepen en todoID meegeven. IDK waarom hij DeleteToDo soms niet vindt maar hij werkt wel (?)
	err = database.DeleteToDo(todoID)
	if err != nil {
		log.Printf("Failed to reach the DeleteToDo function passing the todoID: %v", err)
		return err
	}
	return c.Redirect(http.StatusFound, "/todos")
}
