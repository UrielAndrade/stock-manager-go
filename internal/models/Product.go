package models

type Product struct {
	ID               int          `json:"id" gorm:"primaryKey;autoIncrement"`
	IdManufacturerFk int          `json:"manufacturer_id" gorm:"not null;column:id_manufacturer_fk"`
	Manufacturer     Brand 		  `json:"manufacturer" gorm:"foreignKey:IdManufacturerFk;references:ID"`
	Name             string       `json:"name"`
	Price            float64      `json:"price"`
	Brand            string       `json:"brand"`
	Quantity         int          `json:"quantity"`
}
