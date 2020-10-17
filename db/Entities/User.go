package entities

type User struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     uint   `json:"age"`
}

func (User) TableName() string {
	return "User"
}
