package rbac

import (
	"testing"
	"log"

)

func TestCreatePermission(t *testing.T) {
	p, err := NewPermission("RO")
	if err != nil {
		t.Error()
	}
	log.Println(p)
}

func TestCreateRole(t *testing.T) {
	r, err := NewRole("Admin")
	if err != nil {
		t.Error()
	}
	log.Println(r)
}

func TestCreateUser(t *testing.T) {
	u, err := NewUser("Zarul")
	if err != nil {
		t.Error()
	}
	log.Println(u)
}

func TestUserWithRoleWithPermission(t *testing.T) {
	pro, err := NewPermission("RO")
	if err != nil {
		t.Error()
	}

	prw, err := NewPermission("RW")
	if err != nil {
		t.Error()
	}

	r, err := NewRole("Admin")
	if err != nil {
		t.Error()
	}
	u, err := NewUser("Zarul")
	if err != nil {
		t.Error()
	}

	r, err = r.SetPermission(pro)
	if err != nil {
		t.Error()
	}
	r, err = r.SetPermission(prw)
	if err != nil {
		t.Error()
	}
	u, err = u.SetRole(r)
	if err != nil {
		t.Error()
	}

	log.Println(r)
}