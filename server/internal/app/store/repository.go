package store

import model "github.com/t67y110v/software-engineering/internal/app/model/user"

type PostgresStoreRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
	FindByID(string) (*model.User, error)
}

type MongoStoreRepository interface {
	GetProduct() error
}
