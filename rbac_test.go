package rbac

import (
	"testing"
	"log"

)

func TestInitDB(t *testing.T) {
	InitDB()

}

func TestCreatePermission(t *testing.T) {
	db := InitDB()
	p, err := db.NewPermission("RO")
	if err != nil {
		t.Error()
	}
	defer db.CloseDB();
	log.Println(p)
}

func TestCreateRole(t *testing.T) {
	db := InitDB()
	r, err := db.NewRole("Admin")
	if err != nil {
		t.Error()
	}
	defer db.CloseDB();

	log.Println(r)
}

func TestCreateUser(t *testing.T) {

	db := InitDB()
	
	u, err := db.NewUser("testuser")

	if err != nil {
		t.Error()
	}
	defer db.CloseDB();
	log.Println(u)
}

func TestUserWithRoleWithPermission(t *testing.T) {
	db := InitDB()
	
	pro, err := db.NewPermission("RO")
	if err != nil {
		t.Error()
	}

	prw, err := db.NewPermission("RW")
	if err != nil {
		t.Error()
	}

	r, err := db.NewRole("Admin")
	if err != nil {
		t.Error()
	}
	
	
	u, err := db.NewUser("Zarul")
	if err != nil {
		t.Error()
	}

	r, err = db.SetPermission(pro, r)
	if err != nil {
		t.Error()
	}
	r, err = db.SetPermission(prw,r)
	if err != nil {
		t.Error()
	}
	u, err = db.SetRole(r,u)
	if err != nil {
		t.Error()
	}
	defer db.CloseDB();
	log.Println(r)
}