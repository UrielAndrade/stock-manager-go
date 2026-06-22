package postgres

import (
    "context"
    "stock-manager-go/internal/domain"
    "stock-manager-go/internal/port"
    "stock-manager-go/internal/database"
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
