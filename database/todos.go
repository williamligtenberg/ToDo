package database

import (
	"ToDoApplication/models"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
)

func DeleteToDo(id int) error {
	// Variable voor ToDo item.
	var result models.ToDo
	// Items in de database zoeken en error afhandelen.
	if err := DB().Where("id = ?", id).First(&result).Error; err != nil {
		return fmt.Errorf("error looking for item: %v", err)
	}

	// Item verwijderen en error afhandelen.
	if err := DB().Delete(&result).Error; err != nil {
		return fmt.Errorf("error deleting item: %v", err)
	}
	return nil
}

func CreateToDo(ToDo models.ToDo) error {
	// ToDo aanmaken en error afhandelen.
	if err := DB().Create(&ToDo).Error; err != nil {
		fmt.Errorf("Error while creating ToDo: %v")
	}
	return nil
}

func RetrieveToDo(c echo.Context) ([]models.ToDo, error) {
	// User cookie ophalen.
	userCookie, err := c.Cookie("user")
	// Error afhandelen.
	if err != nil {
		return nil, fmt.Errorf("error while getting username cookie: %v", err)
	}

	username := userCookie.Value

	// Slice voor ToDo items maken.
	var result []models.ToDo
	DB().Where("Creator = ?", username).Find(&result)
	return result, nil
}

func MarkToDoAsDone(todoID int) error {
	// Database connectie.
	db := DB()
	// Error afhandelen.
	if db == nil {
		return errors.New("database connection is nil")
	}

	// Database done van ToDo updaten.
	result := db.Model(&models.ToDo{}).Where("id = ?", todoID).Update("done", true)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
