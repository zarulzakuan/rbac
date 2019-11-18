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