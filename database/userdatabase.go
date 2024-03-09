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
	database, err = gorm.Open(sqlite.Open("users.db"))
	if err != nil {
		panic(err)
	}
	err = database.AutoMigrate(models.User{})
	if err != nil {
		log.Fatal(err)
	}
}

func DB() *gorm.DB {
	if database == nil {
		connect()
	}
	return database
}
