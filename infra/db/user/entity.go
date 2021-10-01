package user

import (
	"gorm.io/gorm"
)

type Entity struct {
	Id      uint
	Name    string
	Surname string
	Age     uint
	gorm.Model
}

func (Entity) TableName() string {
	return "User"
}
