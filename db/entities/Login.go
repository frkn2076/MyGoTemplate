package entities

import(
	"gorm.io/gorm"
)

type Login struct {
	Id       uint   `gorm:"auto_increment"`
	UserName string `gorm:"primaryKey"`
	Email	 string 
	Password string 
	gorm.Model
}

func (Login) TableName() string {
	return "Login"
}
