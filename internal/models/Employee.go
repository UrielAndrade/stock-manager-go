package models

type Employee struct {
	Id       int  `json:"id" gorm:"primaryKey;autoIncrement"`
	IdUserFk User `json:"id_user" gorm:"not null;unique;foreignKey:IdUserFk;references:Id"`
}
