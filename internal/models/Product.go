package models

type Product struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	ManufacturerID int     `json:"manufacturer_id"`
	Price          float64 `json:"price"`
	Brand          string  `json:"brand"`
	Quantity       int     `json:"quantity"`
}
