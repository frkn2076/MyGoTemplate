package entities

type Login struct {
	Id       uint   `json:"id"`
	UserName string `json:"username"`
	Email	 string `json:"email"`
	Password string `json:"password"`
	// User     User   `json:"user" gorm:"foreignKey:Id"`
}

func (Login) TableName() string {
	return "Login"
}
