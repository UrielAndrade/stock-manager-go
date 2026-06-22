package postgres

import (
    "context"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "stock-manager-go/internal/domain"
    "stock-manager-go/internal/port"
    "stock-manager-go/internal/database"
    "gorm.io/gorm"
)

// orderRepository is a PostgreSQL implementation of port.OrderRepository using GORM.
type orderRepository struct {
    db *gorm.DB
}

// NewOrderRepository creates a new OrderRepository backed by the global GORM DB.
func NewOrderRepository() port.OrderRepository {
    return &orderRepository{db: database.DB}
}

// Create inserts a new order into the database.
func (r *orderRepository) Create(ctx context.Context, order *domain.Order) error {
    if order == nil {
        return errors.New("order is nil")
    }
    // Use the context for traceability.
    result := r.db.WithContext(ctx).Create(order)
    return result.Error
}

// FindByID retrieves an order by its UUID.
func (r *orderRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Order, error) {
    var order domain.Order
    result := r.db.WithContext(ctx).First(&order, "id = ?", id)
    if result.Error != nil {
        return nil, result.Error
    }
    return &order, nil
}

// UpdateStatus updates only the status field of an order.
func (r *orderRepository) UpdateStatus(ctx context.Context, id uuid.UUID, status domain.OrderStatus) error {
    result := r.db.WithContext(ctx).Model(&domain.Order{}).Where("id = ?", id).Update("status", status)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return fmt.Errorf("no order found with id %s", id)
    }
    return nil
}

// List returns orders filtered by the provided criteria.
func (r *orderRepository) List(ctx context.Context, filter port.OrderFilter) ([]*domain.Order, error) {
    var orders []*domain.Order
    query := r.db.WithContext(ctx).Model(&domain.Order{})
    if filter.UserID != nil {
        query = query.Where("user_id = ?", *filter.UserID)
    }
    if filter.ProductID != nil {
        query = query.Where("product_id = ?", *filter.ProductID)
    }
    if filter.Status != nil {
        query = query.Where("status = ?", *filter.Status)
    }
    if filter.Limit > 0 {
        query = query.Limit(filter.Limit)
    }
    if filter.Offset > 0 {
        query = query.Offset(filter.Offset)
    }
    result := query.Find(&orders)
    return orders, result.Error
}

// Ensure orderRepository satisfies the interface at compile time.
var _ port.OrderRepository = (*orderRepository)(nil)
