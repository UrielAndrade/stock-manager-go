package handlers

import (
    "context"
    "net/http"
    "strconv"

    "stock-manager-go/internal/dtos"
    "stock-manager-go/internal/usecase"
    "stock-manager-go/internal/port"
    "github.com/go-fuego/fuego"
)

// OrderHandler groups the use‑case service.
type OrderHandler struct {
    svc *usecase.OrderService
}

// NewOrderHandler creates a handler with required dependencies.
func NewOrderHandler(or port.OrderRepository, ar port.AuditRepository) *OrderHandler {
    return &OrderHandler{svc: usecase.NewOrderService(or, ar)}
}

// CreateOrder creates a new order.
func (h *OrderHandler) CreateOrder(c fuego.ContextWithBody[dtos.CreateOrderDTO]) (dtos.OrderResponseDTO, error) {
    input, err := c.Body()
    if err != nil {
        return dtos.OrderResponseDTO{}, fuego.BadRequestError{Err: err}
    }
    order, err := h.svc.CreateOrder(c.Context(), input.UserID, input.ProductID, input.Quantity, input.Price, input.Type)
    if err != nil {
        return dtos.OrderResponseDTO{}, err
    }
    return dtos.ToOrderResponse(order), nil
}

// ExecuteOrder marks an order as executed.
func (h *OrderHandler) ExecuteOrder(c fuego.ContextNoBody) (dtos.OrderResponseDTO, error) {
    idStr := c.PathParam("id")
    id, err := uuid.Parse(idStr)
    if err != nil {
        return dtos.OrderResponseDTO{}, fuego.BadRequestError{Err: err}
    }
    order, err := h.svc.ExecuteOrder(c.Context(), id)
    if err != nil {
        return dtos.OrderResponseDTO{}, err
    }
    return dtos.ToOrderResponse(order), nil
}

// CancelOrder marks an order as canceled.
func (h *OrderHandler) CancelOrder(c fuego.ContextNoBody) (dtos.OrderResponseDTO, error) {
    idStr := c.PathParam("id")
    id, err := uuid.Parse(idStr)
    if err != nil {
        return dtos.OrderResponseDTO{}, fuego.BadRequestError{Err: err}
    }
    order, err := h.svc.CancelOrder(c.Context(), id)
    if err != nil {
        return dtos.OrderResponseDTO{}, err
    }
    return dtos.ToOrderResponse(order), nil
}

// GetOrders returns a list of orders.
func (h *OrderHandler) GetOrders(c fuego.ContextNoBody) ([]dtos.OrderResponseDTO, error) {
    // Simple filter: none for now.
    orders, err := h.svc.ListOrders(c.Context(), port.OrderFilter{})
    if err != nil {
        return nil, err
    }
    var resp []dtos.OrderResponseDTO
    for _, o := range orders {
        resp = append(resp, dtos.ToOrderResponse(o))
    }
    return resp, nil
}

// GetOrder returns a single order by ID.
func (h *OrderHandler) GetOrder(c fuego.ContextNoBody) (dtos.OrderResponseDTO, error) {
    idStr := c.PathParam("id")
    id, err := uuid.Parse(idStr)
    if err != nil {
        return dtos.OrderResponseDTO{}, fuego.BadRequestError{Err: err}
    }
    order, err := h.svc.orderRepo.FindByID(c.Context(), id)
    if err != nil {
        return dtos.OrderResponseDTO{}, fuego.NotFoundError{Err: err}
    }
    return dtos.ToOrderResponse(order), nil
}
