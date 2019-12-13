package rbac

import (
	
	"github.com/zarulzakuan/rbac/components"
	
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	
)

type DBEnv struct {
	*gorm.DB
}

// InitDB creates tables and their schemas
func InitDB() *DBEnv {
	db, err := gorm.Open("sqlite3", "rbac.db")
	if err != nil {
		panic("failed to connect database")
	}

	if !db.HasTable(&components.User{}) {
		db.CreateTable(&components.User{})
	}
	if !db.HasTable(&components.Role{}) {
		db.CreateTable(&components.Role{})
	}
	if !db.HasTable(&components.Permission{}) {
		db.CreateTable(&components.Permission{})
	}
	return &DBEnv{db}
}

// NewPermission creates new permission
func (dbenv *DBEnv) NewPermission(name string) (*components.Permission, error) {
	var p components.Permission
	p.ID = uuid.New()
	p.Name = name
	if dbObj := dbenv.Create(&components.Permission{ID: uuid.New(), Name: name}); dbObj.Error != nil {
		return nil, dbObj.Error
	}
	return &p, nil
}

// NewRole creates new role (duh)
func (dbenv *DBEnv) NewRole(name string) (*components.Role, error) {
	var r components.Role
	r.ID = uuid.New()
	r.Name = name
	if dbObj := dbenv.Create(&components.Role{ID: uuid.New(), Name: name}); dbObj.Error != nil {
		return nil, dbObj.Error
	}
	return &r, nil
}

// NewUser creates new user
func (dbenv *DBEnv) NewUser(name string) (*components.User, error) {
	var u components.User
	u.ID = uuid.New()
	u.Name = name

	if dbObj := dbenv.Create(&components.User{ID: uuid.New(), Name: name}); dbObj.Error != nil {
		return nil, dbObj.Error
	}

	return &u, nil
}

func (dbenv *DBEnv) CloseDB() {
	dbenv.Close()
}
