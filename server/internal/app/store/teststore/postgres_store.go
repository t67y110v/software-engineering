package teststore

import (
	model "github.com/t67y110v/software-engineering/internal/app/model/user"
	"github.com/t67y110v/software-engineering/internal/app/store"
)

type pgStore struct {
	postgresRepository *PostgresStoreRepository
}

func NewPostgres() *pgStore {
	return &pgStore{}
}

func (s *pgStore) UserRepository() store.PostgresStoreRepository {
	if s.postgresRepository != nil {
		return s.postgresRepository
	}
	s.postgresRepository = &PostgresStoreRepository{
		pgStore:    s,
		usersEmail: make(map[string]*model.User),
		usersIDs:   make(map[int]*model.User),
	}

	return s.postgresRepository
}
