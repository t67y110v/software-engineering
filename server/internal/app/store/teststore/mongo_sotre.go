package teststore

import (
	model "github.com/t67y110v/software-engineering/internal/app/model/product"
	"github.com/t67y110v/software-engineering/internal/app/store"
)

type mgStore struct {
	mongoRepository *MongoStoreRepository
}

func NewMongo() *mgStore {
	return &mgStore{}
}

func (s *mgStore) ProductRepository() store.MongoStoreRepository {
	if s.mongoRepository != nil {
		return s.mongoRepository
	}
	s.mongoRepository = &MongoStoreRepository{
		mgStore:  s,
		products: make(map[int]*model.Product),
		carts:    make(map[int]*model.Cart),
	}

	return s.mongoRepository
}
