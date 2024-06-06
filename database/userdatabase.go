package database

import (
	"ToDoApplication/models"
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var database *gorm.DB

func connect() error {
	var err error

	database, err = gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}

	if err := database.AutoMigrate(&models.User{}); err != nil {
		// Error afhandelen.
		return fmt.Errorf("failed to migrate User model: %v", err)
	}
	if err := database.AutoMigrate(&models.ToDo{}); err != nil {
		// Error afhandelen.
		return fmt.Errorf("failed to migrate ToDo model: %v", err)
	}

	return nil
}

func DB() *gorm.DB {
	if database == nil {
		if err := connect(); err != nil {
			fmt.Errorf("error connecting database connection: %v", err)
			return nil
		}
	}
	return database
}
