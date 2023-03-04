package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID                int    `json:"ID"`
	Email             string `json:"Email"`
	Name              string `json:"Name"`
	SeccondName       string `json:"SeccondName"`
	Password          string `json:"Password,omitempty"`
	Isadmin           bool   `json:"Isadmin"`
	EncryptedPassword string `json:"-"`
}

func (u *User) BeforeCreate() error {

	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return err
		}

		u.EncryptedPassword = enc
	}

	return nil
}

func (u *User) Sanitize() {
	u.Password = ""
}

func (u *User) ComparePassword(password string) bool {

	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password)) == nil
}

func encryptString(s string) (string, error) {

	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
