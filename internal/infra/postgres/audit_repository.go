package postgres

import (
    "context"
    "fmt"
    "estoque-go/internal/domain"
    "estoque-go/internal/port"
    "estoque-go/internal/database"
    "gorm.io/gorm"
)

// auditRepository implements port.AuditRepository using GORM.
type auditRepository struct {
    db *gorm.DB
}

func NewAuditRepository() port.AuditRepository {
    return &auditRepository{db: database.DB}
}

func (r *auditRepository) Save(ctx context.Context, audit *domain.OrderAudit) error {
    if audit == nil {
        return fmt.Errorf("audit is nil")
    }
    return r.db.WithContext(ctx).Create(audit).Error
}

var _ port.AuditRepository = (*auditRepository)(nil)
