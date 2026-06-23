package domain

import (
	"time"

	"github.com/google/uuid"
)

type OrderAudit struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UserID    int       `gorm:"not null" json:"user_id"`
	Action    string    `json:"action"` // CREATE, UPDATE, EXECUTE, CANCEL
	Entity    string    `json:"entity"` // "order"
	EntityID  uuid.UUID `gorm:"type:uuid" json:"entity_id"`
	OldData   string    `json:"old_data,omitempty"` // JSON snapshot
	NewData   string    `json:"new_data,omitempty"`
	Diff      string    `json:"diff,omitempty"`
}
