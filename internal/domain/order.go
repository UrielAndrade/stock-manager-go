package domain

import (
    "fmt"
    "time"
    "github.com/google/uuid"
)

type OrderType string

type OrderStatus string

const (
    OrderTypeBuy  OrderType = "buy"
    OrderTypeSell OrderType = "sell"

    OrderStatusPending   OrderStatus = "pending"
    OrderStatusExecuted  OrderStatus = "executed"
    OrderStatusCanceled  OrderStatus = "canceled"
)

type Order struct {
    ID          uuid.UUID   `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
    UserID      uuid.UUID   `gorm:"type:uuid;not null" json:"user_id"`
    ProductID   uuid.UUID   `gorm:"type:uuid;not null" json:"product_id"`
    Quantity    int         `gorm:"not null" json:"quantity"`
    Price       float64     `gorm:"not null" json:"price"`
    Type        OrderType   `gorm:"type:varchar(4);not null" json:"type"`
    Status      OrderStatus `gorm:"type:varchar(10);not null" json:"status"`
    CreatedAt   time.Time   `json:"created_at"`
    UpdatedAt   time.Time   `json:"updated_at"`
}

func NewOrder(userID, productID uuid.UUID, quantity int, price float64, orderType OrderType) (*Order, error) {
    if quantity <= 0 {
        return nil, fmt.Errorf("quantity must be greater than zero")
    }
    if price <= 0 {
        return nil, fmt.Errorf("price must be greater than zero")
    }
    if orderType != OrderTypeBuy && orderType != OrderTypeSell {
        return nil, fmt.Errorf("invalid order type")
    }
    return &Order{
        ID:        uuid.New(),
        UserID:    userID,
        ProductID: productID,
        Quantity:  quantity,
        Price:     price,
        Type:      orderType,
        Status:    OrderStatusPending,
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }, nil
}
