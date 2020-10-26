package entities

import(
	"gorm.io/gorm"
)

type Login struct {
	Id		 uint   `gorm:"uniqueIndex;autoIncrement:true"`
	UserName string `gorm:"primaryKey"`
	Email	 string `gorm:"primaryKey"`
	Password string 
	gorm.Model
}

func (Login) TableName() string {
	return "Login"
}
