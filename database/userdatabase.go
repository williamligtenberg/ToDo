package database

import (
	"ToDoApplication/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"log"
)

var database *gorm.DB

func connect() {
	var err error

	// SQLite database openen.
	database, err = gorm.Open(sqlite.Open("users.db"))
	if err != nil {
		panic(err)
	}

	// User en Recipe models migreren.
	err = database.AutoMigrate(&models.User{}, &models.ToDo{})
	if err != nil {
		log.Fatal(err)
	}
}

func DB() *gorm.DB {
	// Database connectie.
	if database == nil {
		connect()
	}
	return database
}
