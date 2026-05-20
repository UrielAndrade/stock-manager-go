package models

type Employee struct {
	Id int `json:"id"`
	*User	
}
