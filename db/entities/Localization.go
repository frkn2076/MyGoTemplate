package entities

import(
	"gorm.io/gorm"
)

type Localization struct {
	Id       uint	`gorm:"auto_increment"`
	Resource string `gorm:"primaryKey"`
	Message  string
	Language string `gorm:"primaryKey"`
	gorm.Model
}

func (Localization) TableName() string {
	return "Localization"
}
