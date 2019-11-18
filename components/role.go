package components

import "github.com/google/uuid"

// Role is given to a single user, but can have multiple permissions
type Role struct {
	ID           uuid.UUID
	Name         string
	PermissionID []int
}

// SetPermission sets single or multiple permission IDs to Role
func (r *Role) SetPermission(p *Permission) (*Role, error) {
	return r, nil
}