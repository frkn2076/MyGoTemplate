package entities

import(
	"gorm.io/gorm"
)

type Login struct {
	Id       uint   
	UserName string 
	Email	 string 
	Password string 
	gorm.Model
}

func (Login) TableName() string {
	return "Login"
}
