package handlers

import (
	"ToDoApplication/database"
	"ToDoApplication/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func NewToDoPageHandler(c echo.Context) error {
	cookies := c.Cookies()
	data := map[string]interface{}{
		"Cookies": cookies,
	}
	return c.Render(http.StatusOK, "newtodo.html", data)
}

func CreateToDoHandler(c echo.Context) error {
	userCookie, err := c.Cookie("user")
	if err != nil {
		fmt.Println("error while getting username cookie")
	}
	username := userCookie.Value
	database.CreateToDo(models.ToDo{
		Title:       c.FormValue("title"),
		Description: c.FormValue("description"),
		Creator:     username,
	})
	return c.Redirect(http.StatusFound, "/todos")
}
