package components

import "github.com/google/uuid"

// User is well, user. 1 user can only have one role
type User struct {
	ID     uuid.UUID
	Name   string
	RoleID uuid.UUID
}

// SetRole sets role id to user
func (u *User) SetRole(r *Role) (*User, error) {
	u.RoleID = r.ID
	return u, nil
}