package components

import (
	"time"

	"github.com/google/uuid"
)

// User is well, user. 1 user can only have one role
type User struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	ID        uuid.UUID `gorm:"primary_key"`
	Name      string    `gorm:"unique;not null"`
	RoleID    uuid.UUID
}

// SetRole sets role id to user
func (dbenv *DBEnv) SetRole(r *Role, u *User) (*User, error) {
	u.RoleID = r.ID
	if dbObj := dbenv.Save(u); dbObj.Error != nil {
		return nil, dbObj.Error
	}
	return u, nil
}
