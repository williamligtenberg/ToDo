package database

import (
	"ToDoApplication/models"
	"github.com/labstack/gommon/log"
)

func CreateUser(user models.User) error {
	if err := DB().Create(&user).Error; err != nil {
		log.Errorf("Error creating user")
	}
	return nil
}
