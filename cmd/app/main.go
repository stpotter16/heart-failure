package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/stpotter16/heart-failure/handlers"
)

func main() {
    s := handlers.New()
    if err := run(s); err != nil {
        log.Fatalf("Application run failed: %v", err)
    }
}

func run(s handlers.Server) error {
    ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM)
    defer stop()

    http.ListenAndServe(":8080", s.Handler())

    <-ctx.Done()
    log.Print("Received termination signal. Shutting down")

    return nil
}
