package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/stpotter16/heart-failure/handlers"
)

func main() {
    port := loadPortConfiguration()
    s := handlers.New()
    if err := run(port, s); err != nil {
        log.Fatalf("Application run failed: %v", err)
    }
}

func run(port string, s handlers.Server) error {
    ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM)
    defer stop()

    http.ListenAndServe(port, s.Handler())

    <-ctx.Done()
    log.Print("Received termination signal. Shutting down")

    return nil
}

func loadPortConfiguration() string {
    port, present := os.LookupEnv("PORT")
    if !present {
        log.Fatalf("Unable to find $PORT environment variable")
    }
    return port
}
