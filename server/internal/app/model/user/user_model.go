package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID                int    `json:"id"`
	Email             string `json:"email"`
	Name              string `json:"name"`
	SeccondName       string `json:"seccondName"`
	Password          string `json:"password,omitempty"`
	Isadmin           bool   `json:"isadmin"`
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
