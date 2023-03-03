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
	store    store.AllStore
	config   *Config
	handlers *handlers.Handlres
}

func newServer(store store.AllStore, config *Config, log logging.Logger) *server {
	s := &server{
		router:   fiber.New(fiber.Config{ServerHeader: "software engineering course api", AppName: "Api v1.0.1"}),
		logger:   log,
		store:    store,
		config:   config,
		handlers: handlers.NewHandlres(store, log),
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
	s.router.Post("/api/register", s.handlers.RegisterHandler(), logger.New())
	s.router.Post("/api/login", s.handlers.FiberLogin(), logger.New())
	s.router.Post("/api/user", s.handlers.User(), logger.New())
	s.router.Post("/api/logout", s.handlers.Logout(), logger.New())

}
