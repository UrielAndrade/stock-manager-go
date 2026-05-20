package models

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
	Address  string `json:"address"`
	Cpf      string `json:"cpf"`
	
}
