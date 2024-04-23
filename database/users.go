package database

import (
	"ToDoApplication/models"
	"github.com/labstack/gommon/log"
)

func CreateUser(user models.User) error {
	// User aanmaken en error afhandelen.
	if err := DB().Create(&user).Error; err != nil {
		log.Errorf("Error creating user")
	}
	return nil
}
