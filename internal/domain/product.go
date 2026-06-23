package domain

// Product representa a entidade de produto no contexto de domínio de pedidos.
type Product struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Quantity int
}
