package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gaurav25kapoor/students-api/internal/config"
	"github.com/gaurav25kapoor/students-api/internal/config/http/handlers/student"
	"github.com/gaurav25kapoor/students-api/internal/config/storage/sqlite"
)

func main() {
	// Load config
	cfg := config.MustLoad()

	// Database setup
	storage, err := sqlite.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("storage initialised", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))

	// Setup routes
	router := http.NewServeMux()
	router.HandleFunc("POST /api/students", student.New(storage))

	// Setup server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}
	slog.Info("server started", slog.String("address", cfg.Addr))

	// Graceful shutdown
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to start server: %v", err)
		}
	}()

	<-done
	slog.Info("shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	} else {
		slog.Info("server shutdown successfully")
	}
}
