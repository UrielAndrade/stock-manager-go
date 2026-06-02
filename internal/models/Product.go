package models

type Product struct {
	ID       int     `json:"id" gorm:"primaryKey;autoIncrement"`
	BrandID  int     `json:"brand_id" gorm:"not null"`
	Brand    Brand   `json:"brand" gorm:"foreignKey:BrandID"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}
