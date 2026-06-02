package models

type Brand struct {
	ID             int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name           string `json:"name" gorm:"index"`
	Country        string `json:"country" gorm:"index"`
	Email          string `json:"email" gorm:"unique"`
	FoundationYear int    `json:"foundation_year" gorm:"index"`
}
