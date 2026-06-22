package models

type Employee struct {
	Id     int  `json:"id" gorm:"primaryKey;autoIncrement"`
	UserId int  `json:"userId" gorm:"not null;unique"`
	User   User `json:"user" gorm:"foreignKey:UserId;references:Id"`
}
