package models

type Employee struct {
	Id int `json:"id" gorm:"primaryKey"`
	*User
}
