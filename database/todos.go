package database

import "ToDoApplication/models"

func CreateToDo(ToDo models.ToDo) {
	DB().Create(&ToDo)
}
