package handlers

import (
	"github.com/t67y110v/software-engineering/internal/app/logging"
	"github.com/t67y110v/software-engineering/internal/app/store"
)

type Handlres struct {
	logger logging.Logger
	store  store.PostgresStore
}

func NewHandlres(store store.PostgresStore, logger logging.Logger) *Handlres {
	return &Handlres{
		store:  store,
		logger: logger,
	}
}
