package port

import (
    "context"
    "github.com/google/uuid"
    "estoque-go/internal/domain"
)

type OrderRepository interface {
    Create(ctx context.Context, order *domain.Order) error
    FindByID(ctx context.Context, id uuid.UUID) (*domain.Order, error)
    UpdateStatus(ctx context.Context, id uuid.UUID, status domain.OrderStatus) error
    List(ctx context.Context, filter OrderFilter) ([]*domain.Order, error)
}

type OrderFilter struct {
    UserID    *uuid.UUID
    ProductID *uuid.UUID
    Status    *domain.OrderStatus
    Limit     int
    Offset    int
}

type AuditRepository interface {
    Save(ctx context.Context, audit *domain.OrderAudit) error
}
