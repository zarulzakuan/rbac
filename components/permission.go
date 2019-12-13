package components

import (
	"time"

	"github.com/google/uuid"
)

// Permission is given to single or multiple roles
type Permission struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	ID        uuid.UUID `gorm:"primary_key"`
	Name      string    `gorm:"unique;not null"`
}
