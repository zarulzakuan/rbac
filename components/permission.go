package components

import "github.com/google/uuid"

// Permission is given to single or multiple roles
type Permission struct {
	ID   uuid.UUID
	Name string
}