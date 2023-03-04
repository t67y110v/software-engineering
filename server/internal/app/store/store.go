package store

type PostgresStore interface {
	UserRepository() PostgresStoreRepository
}

type MongoStore interface {
	ProductRepository() MongoStoreRepository
}
