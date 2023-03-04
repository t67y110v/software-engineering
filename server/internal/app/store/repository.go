package store

import (
	mongoModel "github.com/t67y110v/software-engineering/internal/app/model/product"
	postgresModel "github.com/t67y110v/software-engineering/internal/app/model/user"
)

type PostgresStoreRepository interface {
	Create(*postgresModel.User) error
	FindByEmail(string) (*postgresModel.User, error)
	FindByID(string) (*postgresModel.User, error)
}

type MongoStoreRepository interface {
	GetProduct() error
	GetAllProducts() ([]*mongoModel.Product, error)
	Filter(filter string) ([]*mongoModel.Product, error)
}
