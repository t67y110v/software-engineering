package store

import (
	"database/sql"
	"log"

	"github.com/t67y110v/software-engineering/internal/app/store"

	_ "github.com/lib/pq"
)

type Store struct {
	db                      *sql.DB
	postgresStoreRepository *PostgresStoreRepository
}

func NewPostgresDB(db *sql.DB) *Store {
	log.Println("PostgreSQL initialization")

	return &Store{
		db: db,
	}
}

func (s *Store) UserRepository() store.PostgresStoreRepository {
	if s.postgresStoreRepository != nil {
		return s.postgresStoreRepository
	}

	s.postgresStoreRepository = &PostgresStoreRepository{
		store: s,
	}
	return s.postgresStoreRepository
}
