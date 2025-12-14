package main

import (
	"net/http"
	"encoding/json"
	"os"
	"os/signal"
	"syscall"
	"context"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/health", Health)

	server := &http.Server{
		Addr: ":8080",
		Handler: r,
	}

	go func() {
		server.ListenAndServe()
	}()

	<-sigs

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	server.Shutdown(ctx)
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}