package entities

type Setting struct {
	Id       uint   `json:"id"`
	Key 	 string `json:"key"`
	Value	 string `json:"value"`
}

func (Setting) TableName() string {
	return "Setting"
}
