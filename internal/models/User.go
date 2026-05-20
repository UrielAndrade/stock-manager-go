package models

type User struct {
	Id       int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Birthday string `json:"birthday"`
	Address  string `json:"address"`
	Cpf      string `json:"cpf"`
}
