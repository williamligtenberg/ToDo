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

	//SQLite database openen.
	database, err = gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}

	// User model migreren.
	if err := database.AutoMigrate(&models.User{}); err != nil {
		// Error afhandelen.
		return fmt.Errorf("failed to migrate User model: %v", err)
	}

	// ToDo model migreren.
	if err := database.AutoMigrate(&models.ToDo{}); err != nil {
		// Error afhandelen.
		return fmt.Errorf("failed to migrate ToDo model: %v", err)
	}

	return nil
}

func DB() *gorm.DB {
	if database == nil {
		// Database connectie maken.
		if err := connect(); err != nil {
			// Error afhandelen.
			fmt.Errorf("error connecting database connection: %v", err)
			return nil
		}
	}
	return database
}
