package main

import (
	"context"
	"flag"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/akonovalovdev/go-simple-DDD/internal/config"
	"github.com/akonovalovdev/go-simple-DDD/internal/infrastructure/app"
)

func main() {
	configPath := flag.String("config", "config/config.yaml", "path to config file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		slog.Error("failed to load config", slog.Any("error", err))
		os.Exit(1)
	}

	a, err := app.New(cfg)
	if err != nil {
		slog.Error("failed to initialize app", slog.Any("error", err))
		os.Exit(1)
	}

	go func() {
		if err := a.Run(); err != nil {
			slog.Error("app run error", slog.Any("error", err))
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	slog.Info("shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Server.ShutdownTimeout)
	defer cancel()

	if err := a.Shutdown(ctx); err != nil {
		slog.Error("shutdown error", slog.Any("error", err))
	}

	slog.Info("stopped")
}
