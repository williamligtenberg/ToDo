package database

import "ToDoApplication/models"

func CreateUser(Users models.User) {
	DB().Create(&Users)
}
