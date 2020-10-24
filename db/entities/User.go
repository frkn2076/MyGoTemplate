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

// func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	
// 	return
//   }
  
// func (u *User) AfterCreate(tx *gorm.DB) (err error) {
// 	tx.Model(u).Update("created_at", time.Now())
// 	return
// }
