package domain

import (
	"github.com/google/uuid"
)

// Product representa a entidade de produto no contexto de domínio de pedidos.
// Nota: Existe uma discrepância entre esta entidade (que usa UUID) e models.Product (que usa int).
type Product struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name     string
	Quantity int
}
