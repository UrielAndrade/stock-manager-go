package models

type Manufacturer struct {
	Id            int    `json:"id" gorm:"primaryKey"`
	Name          string `json:"name"`
	Country       string `json:"country"`
	Email         string `json:"email"`
	FundationYear int    `json:"fundation_year"`
}
