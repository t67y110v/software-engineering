package teststore

import (
	"errors"
	"strconv"

	model "github.com/t67y110v/software-engineering/internal/app/model/user"
)

type PostgresStoreRepository struct {
	pgStore *pgStore

	usersEmail map[string]*model.User
	usersIDs   map[int]*model.User
}

func (r *PostgresStoreRepository) Create(u *model.User) error {
	if err := u.BeforeCreate(); err != nil {
		return err
	}
	r.usersEmail[u.Email] = u
	r.usersIDs[u.ID] = u
	u.ID = len(r.usersEmail)
	return nil
}

func (r *PostgresStoreRepository) FindByEmail(email string) (*model.User, error) {

	u, ok := r.usersEmail[email]
	if !ok {
		return nil, errors.New("record not found")
	}
	return u, nil
}

func (r *PostgresStoreRepository) FindByID(id string) (*model.User, error) {
	ID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	u, ok := r.usersIDs[ID]
	if !ok {
		return nil, errors.New("record not found")
	}

	return u, nil
}
