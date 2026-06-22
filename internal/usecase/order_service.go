package usecase

import (
    "context"
    "encoding/json"
    "fmt"
    "time"
    "github.com/google/uuid"
    "estoque-go/internal/domain"
    "estoque-go/internal/port"
    "estoque-go/internal/database"
    "gorm.io/gorm"
)

type OrderService struct {
    orderRepo port.OrderRepository
    auditRepo port.AuditRepository
}

func NewOrderService(or port.OrderRepository, ar port.AuditRepository) *OrderService {
    return &OrderService{orderRepo: or, auditRepo: ar}
}

// CreateOrder creates a new order and records an audit entry.
func (s *OrderService) CreateOrder(ctx context.Context, userID, productID uuid.UUID, quantity int, price float64, orderType domain.OrderType) (*domain.Order, error) {
    order, err := domain.NewOrder(userID, productID, quantity, price, orderType)
    if err != nil {
        return nil, err
    }
    // Persist order
    if err := s.orderRepo.Create(ctx, order); err != nil {
        return nil, err
    }
    // Record audit
    audit := &domain.OrderAudit{
        UserID:   userID,
        Action:   "CREATE",
        Entity:   "order",
        EntityID: order.ID,
        NewData:  toJSON(order),
        CreatedAt: time.Now(),
    }
    _ = s.auditRepo.Save(ctx, audit) // ignore error for now
    return order, nil
}

// ExecuteOrder marks an order as executed, updates product stock, and audits.
func (s *OrderService) ExecuteOrder(ctx context.Context, orderID uuid.UUID) (*domain.Order, error) {
    order, err := s.orderRepo.FindByID(ctx, orderID)
    if err != nil {
        return nil, err
    }
    if order.Status != domain.OrderStatusPending {
        return nil, fmt.Errorf("order not pending")
    }
    // Adjust product quantity atomically
    tx := database.DB.Begin()
    var product domain.Product // assume domain.Product mirrors models.Product
    if err := tx.WithContext(ctx).First(&product, "id = ?", order.ProductID).Error; err != nil {
        tx.Rollback()
        return nil, err
    }
    if order.Type == domain.OrderTypeBuy {
        if product.Quantity < order.Quantity {
            tx.Rollback()
            return nil, fmt.Errorf("insufficient stock for product %s", product.ID)
        }
        product.Quantity -= order.Quantity
    } else { // sell
        product.Quantity += order.Quantity
    }
    if err := tx.WithContext(ctx).Save(&product).Error; err != nil {
        tx.Rollback()
        return nil, err
    }
    // Update order status
    if err := s.orderRepo.UpdateStatus(ctx, orderID, domain.OrderStatusExecuted); err != nil {
        tx.Rollback()
        return nil, err
    }
    // Commit transaction
    if err := tx.Commit().Error; err != nil {
        return nil, err
    }
    // Refresh order after status update
    order, err = s.orderRepo.FindByID(ctx, orderID)
    if err != nil {
        return nil, err
    }
    // Audit
    audit := &domain.OrderAudit{
        UserID:    order.UserID,
        Action:    "EXECUTE",
        Entity:    "order",
        EntityID:  order.ID,
        OldData:   toJSON(order), // simplistic, could store previous state
        NewData:   toJSON(order),
        CreatedAt: time.Now(),
    }
    _ = s.auditRepo.Save(ctx, audit)
    return order, nil
}

// CancelOrder marks an order as cancelled and audits.
func (s *OrderService) CancelOrder(ctx context.Context, orderID uuid.UUID) (*domain.Order, error) {
    // Update status
    if err := s.orderRepo.UpdateStatus(ctx, orderID, domain.OrderStatusCanceled); err != nil {
        return nil, err
    }
    order, err := s.orderRepo.FindByID(ctx, orderID)
    if err != nil {
        return nil, err
    }
    audit := &domain.OrderAudit{
        UserID:    order.UserID,
        Action:    "CANCEL",
        Entity:    "order",
        EntityID:  order.ID,
        NewData:   toJSON(order),
        CreatedAt: time.Now(),
    }
    _ = s.auditRepo.Save(ctx, audit)
    return order, nil
}

func (s *OrderService) ListOrders(ctx context.Context, filter port.OrderFilter) ([]*domain.Order, error) {
    return s.orderRepo.List(ctx, filter)
}

// Helper to marshal any struct to JSON string.
func toJSON(v interface{}) string {
    b, _ := json.Marshal(v)
    return string(b)
}
