package nosqlstore

import (
	"log"

	"github.com/t67y110v/software-engineering/internal/app/store"

	"go.mongodb.org/mongo-driver/mongo"
)

type Store struct {
	client               *mongo.Client
	mongoStoreRepository *MongoStoreRepository
}

func NewMongoDB(client *mongo.Client) *Store {
	log.Println("MongoDB initialization")

	return &Store{
		client: client,
	}
}

func (s *Store) ProductRepository() store.MongoStoreRepository {
	if s.mongoStoreRepository != nil {
		return s.mongoStoreRepository
	}
	s.mongoStoreRepository = &MongoStoreRepository{
		store: s,
	}
	return s.mongoStoreRepository
}
