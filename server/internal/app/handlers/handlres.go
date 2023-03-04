package handlers

import (
	"github.com/t67y110v/software-engineering/internal/app/logging"
	"github.com/t67y110v/software-engineering/internal/app/store"
)

type Handlres struct {
	logger  logging.Logger
	pgStore store.PostgresStore
	mgStore store.MongoStore
}

func NewHandlres(pgstore store.PostgresStore, mgstore store.MongoStore, logger logging.Logger) *Handlres {
	return &Handlres{
		pgStore: pgstore,
		mgStore: mgstore,
		logger:  logger,
	}
}
