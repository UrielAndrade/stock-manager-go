package dtos

import (
    "github.com/google/uuid"
    "estoque-go/internal/domain"
    "time"
)

type CreateOrderDTO struct {
    UserID    uuid.UUID   `json:"user_id"`
    ProductID uuid.UUID   `json:"product_id"`
    Quantity  int         `json:"quantity"`
    Price     float64     `json:"price"`
    Type      domain.OrderType `json:"type"`
}

type OrderResponseDTO struct {
    ID        uuid.UUID   `json:"id"`
    UserID    uuid.UUID   `json:"user_id"`
    ProductID uuid.UUID   `json:"product_id"`
    Quantity  int         `json:"quantity"`
    Price     float64     `json:"price"`
    Type      domain.OrderType `json:"type"`
    Status    domain.OrderStatus `json:"status"`
    CreatedAt time.Time   `json:"created_at"`
    UpdatedAt time.Time   `json:"updated_at"`
}

// ToOrderResponse maps a domain.Order to its DTO representation.
func ToOrderResponse(o *domain.Order) OrderResponseDTO {
    return OrderResponseDTO{
        ID:        o.ID,
        UserID:    o.UserID,
        ProductID: o.ProductID,
        Quantity:  o.Quantity,
        Price:     o.Price,
        Type:      o.Type,
        Status:    o.Status,
        CreatedAt: o.CreatedAt,
        UpdatedAt: o.UpdatedAt,
    }
}
