package store

import (
	"database/sql"

	"github.com/t67y110v/software-engineering/internal/app/model"
	"github.com/t67y110v/software-engineering/internal/app/store"
)

type AllStoreRepository struct {
	store *Store
}

func (r *AllStoreRepository) Create(u *model.User) error {
	if err := u.BeforeCreate(); err != nil {
		return err
	}
	return r.store.db.QueryRow(
		"INSERT INTO users (email, encrypted_password, userName, seccondName) VALUES ($1, $2,$3,$4) RETURNING id",
		u.Email,
		u.EncryptedPassword,
		u.Name,
		u.SeccondName,
	).Scan(&u.ID)
}

func (r *AllStoreRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password, isadmin, username, seccondname FROM users WHERE email = $1",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
		&u.Isadmin,
		&u.Name,
		&u.SeccondName,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}
	return u, nil

}

func (r *AllStoreRepository) FindByID(ID string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password, isadmin, username, seccondname FROM users WHERE id = $1",
		ID,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
		&u.Isadmin,
		&u.Name,
		&u.SeccondName,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}
	return u, nil

}
