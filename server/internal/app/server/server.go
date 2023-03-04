package server

import (
	//"html/template"
	"net/http"

	"github.com/t67y110v/software-engineering/internal/app/handlers"
	"github.com/t67y110v/software-engineering/internal/app/logging"
	"github.com/t67y110v/software-engineering/internal/app/store"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/gofiber/fiber/v2/middleware/logger"
)

type server struct {
	router   *fiber.App
	logger   logging.Logger
	pgStore  store.PostgresStore
	mgStore  store.MongoStore
	config   *Config
	handlers *handlers.Handlres
}

func newServer(pgstore store.PostgresStore, mgstore store.MongoStore, config *Config, log logging.Logger) *server {
	s := &server{
		router:   fiber.New(fiber.Config{ServerHeader: "software engineering course api", AppName: "Api v1.0.1"}),
		logger:   log,
		pgStore:  pgstore,
		config:   config,
		handlers: handlers.NewHandlres(pgstore, mgstore, log),
	}
	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

func (s *server) configureRouter() {
	s.router.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowHeaders:     "Origin, Content-Type, Accept",
	}))
	//api := s.router.Group("/api")
	//api.Use(logger.New())
	s.router.Post("/api/register", s.handlers.Register(), logger.New())
	s.router.Post("/api/login", s.handlers.Login(), logger.New())
	s.router.Post("/api/user", s.handlers.CheckJWT(), logger.New())
	s.router.Post("/api/logout", s.handlers.Logout(), logger.New())
	s.router.Get("/test", s.handlers.AddProduct(), logger.New())

}
