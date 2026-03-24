package app

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/akonovalovdev/go-simple-DDD/internal/adapter/repository/postgres"
	"github.com/akonovalovdev/go-simple-DDD/internal/config"
	infrahttp "github.com/akonovalovdev/go-simple-DDD/internal/infrastructure/http"
	infrapostgres "github.com/akonovalovdev/go-simple-DDD/internal/infrastructure/postgres"
)

type App struct {
	cfg    *config.Config
	db     *sql.DB
	server *infrahttp.Server
	logger *slog.Logger
}

func New(cfg *config.Config) (*App, error) {
	logger := newLogger(cfg.Log.Level)

	db, err := infrapostgres.Open(cfg.Database)
	if err != nil {
		// Non-fatal for local development — log and continue without DB.
		logger.Warn("failed to connect to database, running without DB", slog.Any("error", err))
	}

	if db != nil {
		_ = postgres.NewRepository(db)
	}

	server := infrahttp.NewServer(
		cfg.Server.Port,
		cfg.Server.ReadTimeout,
		cfg.Server.WriteTimeout,
		logger,
	)

	return &App{cfg: cfg, db: db, server: server, logger: logger}, nil
}

func (a *App) Run() error {
	if err := a.server.Start(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("http server: %w", err)
	}
	return nil
}

func (a *App) Shutdown(ctx context.Context) error {
	if err := a.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("shutdown http server: %w", err)
	}
	if a.db != nil {
		if err := a.db.Close(); err != nil {
			return fmt.Errorf("close db: %w", err)
		}
	}
	return nil
}

func newLogger(level string) *slog.Logger {
	var logLevel slog.Level
	switch level {
	case "debug":
		logLevel = slog.LevelDebug
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo
	}
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel}))
}
