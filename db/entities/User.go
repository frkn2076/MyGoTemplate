package entities

import(
	// "time"

	"gorm.io/gorm"
)

type User struct {
	Id      uint   
	Name    string 
	Surname string 
	Age     uint  
	gorm.Model
}

func (User) TableName() string {
	return "User"
}
