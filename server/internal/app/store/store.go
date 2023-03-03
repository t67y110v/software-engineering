package store

type PostgresStore interface {
	UserRepository() PostgresStoreRepository
}
