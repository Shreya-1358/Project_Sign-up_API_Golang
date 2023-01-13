package model

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	MobileNo string `json:"mobile"`
	Address  string `json:"address"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (b *User) TableName() string {
	return "user"
}
