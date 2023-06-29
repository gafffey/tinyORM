package session

import "testing"

type User struct {
	Id   int
	Name string
}

func TestSession_Model(t *testing.T) {
	s := NewSession().Model(&User{})
	table := s.RefTable()
	if table.Name != "User" {
		t.Fatal("failed to call Model")
	}
}

func TestSession_DropTable(t *testing.T) {
	s := NewSession().Model(&User{})
	if err := s.DropTable(); err != nil {
		t.Fatal(err)
	}
}
