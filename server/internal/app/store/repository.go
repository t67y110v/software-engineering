package store

import "github.com/t67y110v/software-engineering/internal/app/model"

type PostgresStoreRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
	FindByID(string) (*model.User, error)
}
