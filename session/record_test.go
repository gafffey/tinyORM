package session

import "testing"

func TestSession_Find(t *testing.T) {
	s := NewSession().Model(&User{})
	users := make([]*User, 0)
	if err := s.Find(&users); err != nil || len(users) != 2 {
		t.Fatal("failed to query all")
	}

	for _, u := range users {
		t.Log(*u)
	}
}

func TestSession(t *testing.T) {
	s := NewSession().Model(&User{})
	users := make([]*User, 0)
	if err := s.Where("Name like ?", "a%").Limit(2).OrderBy("Id DESC").Find(&users); err != nil || len(users) != 2 {
		t.Fatal("failed to query all")
	}

	for _, u := range users {
		t.Log(*u)
	}
}

func TestSession_Limit(t *testing.T) {
	s := NewSession().Model(&User{})
	var users []*User
	if err := s.Limit(1).Find(&users); err != nil {
		t.Fatal("failed to query all")
	}

	t.Log(len(users))
	t.Log(*users[0])
}

func TestSession_Where(t *testing.T) {
	s := NewSession().Model(&User{})
	var users []*User
	if err := s.Where("Name = ?", "Tom").Find(&users); err != nil {
		t.Fatal("failed to query all")
	}

	t.Log(len(users))
	t.Log(*users[0])
}

func TestSession_OrderBy(t *testing.T) {
	s := NewSession().Model(&User{})
	var users []*User
	if err := s.OrderBy("Id DESC").Find(&users); err != nil {
		t.Fatal("failed to query all")
	}

	t.Log(len(users))
	t.Log(*users[0])
	t.Log(*users[1])
}
