package server

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"

	"github.com/t67y110v/software-engineering/internal/app/logging"
	"github.com/t67y110v/software-engineering/internal/app/store/nosqlstore"
	store "github.com/t67y110v/software-engineering/internal/app/store/sqlstore"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Start(config *Config) error {
	pgdb, err := newPostgresDB(config.PostgresDatabaseURL)
	if err != nil {
		return err
	}
	defer pgdb.Close()

	mgdb, err := newMongoDB(config.MongoDatabaseURL)

	pgStore := store.NewPostgresDB(pgdb)
	mgStore := nosqlstore.NewMongoDB(mgdb)
	logger := logging.GetLogger()
	server := newServer(pgStore, mgStore, config, logger)
	//StartServerWithGracefulShutdown(server, config.BindAddr)
	return server.router.Listen(config.BindAddr)
}

func newPostgresDB(postgresDatabaseURL string) (*sql.DB, error) {
	log.Printf("Database initialization: %s\n", postgresDatabaseURL)
	db, err := sql.Open("postgres", postgresDatabaseURL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func newMongoDB(mongoDatabaseURL string) (*mongo.Client, error) {
	ctx := context.TODO()
	clientOptions := options.Client().ApplyURI(mongoDatabaseURL)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client, nil

}

func StartServerWithGracefulShutdown(s *server, addr string) {
	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := s.router.Shutdown(); err != nil {
			s.logger.Warningf("Server is not shutting down! reason: %v", err)
		}
		close(idleConnsClosed)
	}()
	if err := s.router.Listen(addr); err != nil {
		s.logger.Warningf("Server is not running! reason: %v", err)
	}
	<-idleConnsClosed

}
