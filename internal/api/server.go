package api

import (
	"net/http"
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/seeques/task-api/internal/config"
	"github.com/seeques/task-api/internal/handler"
)

type Server struct {
	httpServer *http.Server
	router chi.Router
}

func NewServer() *Server {
	cfg := config.LoadConfig()

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
	}
}

func (s *Server) Start() {
	s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) {
	s.httpServer.Shutdown(ctx)
}