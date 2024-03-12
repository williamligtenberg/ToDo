package models

import "gorm.io/gorm"

type ToDo struct {
	gorm.Model
	Title       string
	Description string
	UserID      uint
	User        User `gorm:"foreignkey:UserID"`
}
