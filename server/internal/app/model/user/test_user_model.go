package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Email:       "userTest@test.org",
		Password:    "qwerty",
		Name:        "test",
		SeccondName: "testSeccondName",
		Isadmin:     false,
	}
}
