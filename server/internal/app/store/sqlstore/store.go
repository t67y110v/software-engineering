package store

import (
	"database/sql"
	"log"

	"github.com/t67y110v/software-engineering/internal/app/store"

	_ "github.com/lib/pq"
)

type Store struct {
	db                 *sql.DB
	allStoreRepository *AllStoreRepository
}

func New(db *sql.DB) *Store {
	log.Println("Store initialization")

	return &Store{
		db: db,
	}
}

func (s *Store) Everythink() store.AllStoreRepository {
	if s.allStoreRepository != nil {
		return s.allStoreRepository
	}

	s.allStoreRepository = &AllStoreRepository{
		store: s,
	}
	return s.allStoreRepository
}
