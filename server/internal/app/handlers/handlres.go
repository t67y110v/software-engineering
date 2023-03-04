package handlers

import (
	"github.com/t67y110v/software-engineering/internal/app/logging"
	"github.com/t67y110v/software-engineering/internal/app/store"
)

type Handlers struct {
	logger  logging.Logger
	pgStore store.PostgresStore
	mgStore store.MongoStore
}

func NewHandlers(pgstore store.PostgresStore, mgstore store.MongoStore, logger logging.Logger) *Handlers {
	return &Handlers{
		pgStore: pgstore,
		mgStore: mgstore,
		logger:  logger,
	}
}
