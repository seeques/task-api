package main

import (
	"os"
	"os/signal"
	"syscall"
	"context"
	"time"

	"github.com/seeques/task-api/internal/api"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	server := api.NewServer()

	go server.Start()

	<-sigs

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	server.Shutdown(ctx)
}