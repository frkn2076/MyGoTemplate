package login

import (
	"gorm.io/gorm"
)

type Entity struct {
	Id       uint   `gorm:"uniqueIndex;autoIncrement:true"`
	UserName string `gorm:"primaryKey"`
	Email    string `gorm:"primaryKey"`
	Password string
	gorm.Model
}

func (Entity) TableName() string {
	return "Login"
}
