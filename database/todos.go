package database

import (
	"ToDoApplication/models"
	"fmt"
	"github.com/labstack/echo/v4"
)

func CreateToDo(ToDo models.ToDo) {
	DB().Create(&ToDo)
}

func RetrieveToDo(c echo.Context) []models.ToDo {
	userCookie, err := c.Cookie("user")
	if err != nil {
		fmt.Println("Error while getting username cookie in RetrieveToDo()")
	}

	username := userCookie.Value

	var result []models.ToDo
	DB().Where("Creator = ?", username).Find(&result)
	return result
}

func DeleteToDo(id int) {
	var result models.ToDo
	if err := DB().Where("id = ?", id).First(&result).Error; err != nil {
		fmt.Println("error")
		return
	}

	if err := DB().Delete(&result).Error; err != nil {
		fmt.Println("error")
		return
	}
}
