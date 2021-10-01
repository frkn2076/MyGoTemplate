package localization

import (
	"gorm.io/gorm"
)

type Entity struct {
	Id       uint   `gorm:"uniqueIndex;autoIncrement:true"`
	Resource string `gorm:"primaryKey"`
	Message  string
	Language string `gorm:"primaryKey"`
	gorm.Model
}

func (Entity) TableName() string {
	return "Localization"
}
