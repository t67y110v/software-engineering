package server

import (
	//"html/template"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/t67y110v/software-engineering/docs"
	"github.com/t67y110v/software-engineering/internal/app/handlers"
	"github.com/t67y110v/software-engineering/internal/app/logging"
	"github.com/t67y110v/software-engineering/internal/app/store"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger" // swagger handler
)

type server struct {
	router   *fiber.App
	logger   logging.Logger
	pgStore  store.PostgresStore
	mgStore  store.MongoStore
	config   *Config
	handlers *handlers.Handlers
}

func newServer(pgstore store.PostgresStore, mgstore store.MongoStore, config *Config, log logging.Logger) *server {
	s := &server{
		router:   fiber.New(fiber.Config{ServerHeader: "software engineering course api", AppName: "Api v1.0.1"}),
		logger:   log,
		pgStore:  pgstore,
		mgStore:  mgstore,
		config:   config,
		handlers: handlers.NewHandlers(pgstore, mgstore, log),
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
	// localhost:4000/user/register

	s.router.Get("/swagger/*", swagger.HandlerDefault)

	///////// USER GROUP ///////////////
	////////////////////////////////////
	user := s.router.Group("/user")
	user.Use(logger.New())
	user.Post("/register", s.handlers.Register())
	user.Post("/login", s.handlers.Login())
	user.Post("/check", s.handlers.CheckJWT())
	//////////////////////////////////////

	//////// PRODUCT GRUOP ////////////
	///////////////////////////////////
	product := s.router.Group("/product")
	product.Use(logger.New())
	product.Post("/add", s.handlers.AddProduct())
	product.Get("/all", s.handlers.GetAllProducts())
	product.Get("/filter/:category", s.handlers.FilterByCategory())
	product.Delete("/delete", s.handlers.DeleteProduct())

	////////////////////////////////////

	//////////// CART GROUP /////////////
	/////////////////////////////////////
	cart := s.router.Group("/cart")
	cart.Use(logger.New())
	cart.Post("/add", s.handlers.AddToCart())
	cart.Get("/get/:user_id", s.handlers.GetCart())
	cart.Delete("/delete", s.handlers.DeleteFromCart())
	cart.Delete("/clear", s.handlers.ClearCart())
}
