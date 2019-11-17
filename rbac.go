package rbac

import (
	"github.com/zarulzakuan/rbac/components"
	"github.com/google/uuid"

)
// NewPermission creates new permission
func NewPermission(name string) (*components.Permission, error) {
	var p components.Permission
	p.ID = uuid.New()
	p.Name = name
	return &p, nil
}

// NewRole creates new role (duh)
func NewRole(name string) (*components.Role, error) {
	var r components.Role
	r.ID = uuid.New()
	r.Name = name
	return &r, nil
}

// NewUser creates new user
func NewUser(name string) (*components.User, error) {
	var u components.User
	u.ID = uuid.New()
	u.Name = name
	return &u, nil
}




