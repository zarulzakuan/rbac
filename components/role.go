package components

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type DBEnv struct {
	*gorm.DB
}

// Role is given to a single user, but can have multiple permissions
type Role struct {
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
	ID          uuid.UUID     `gorm:"primary_key"`
	Name        string        `gorm:"unique;not null"`
	Permissions []*Permission `gorm:"many2many:role_permission;"`
}

// SetPermission sets single or multiple permission IDs to Role
func (dbenv *DBEnv) SetPermission(p *Permission, r *Role) (*Role, error) {
	r.Permissions = append(r.Permissions, p)
	if dbObj := dbenv.Save(r); dbObj.Error != nil {
		return nil, dbObj.Error
	}
	return r, nil
}
