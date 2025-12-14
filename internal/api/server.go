package api

import (
	"net/http"
	"context"

	"github.com/rs/zerolog/log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/seeques/task-api/internal/config"
	"github.com/seeques/task-api/internal/handler"
	"github.com/seeques/task-api/internal/storage"
)

type Server struct {
	httpServer *http.Server
	router chi.Router
	postgresStorage *storage.PostgresStorage
}

func NewServer() *Server {
	cfg := config.LoadConfig()
	pool, err := storage.CreatePool()
	if err != nil {
		log.Fatal().Err(err).Msg("Pool creation failed")
	}

	pgStorage := storage.NewPostgresStorage(pool)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/health", handler.Health)

	server := &http.Server{
		Addr: cfg.Port,
		Handler: r,
	}

	return &Server{
		httpServer: server,
		router: r,
		postgresStorage: pgStorage,
	}
}

func (s *Server) Start() {
	s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) {
	s.httpServer.Shutdown(ctx)
}

func (s *Server) ClosePool() {
	s.postgresStorage.ClosePool()
}