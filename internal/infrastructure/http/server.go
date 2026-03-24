package http

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/akonovalovdev/go-simple-DDD/internal/adapter/controller"
	middleware "github.com/akonovalovdev/go-simple-DDD/internal/infrastructure/middleware/http"
)

type Server struct {
	server *http.Server
	logger *slog.Logger
}

func NewServer(port int, readTimeout, writeTimeout time.Duration, logger *slog.Logger) *Server {
	s := &Server{logger: logger}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", controller.HealthHandler())

	chain := middleware.Logging(logger)(
		middleware.Recovery(logger)(
			middleware.CORS()(mux),
		),
	)

	s.server = &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      chain,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	return s
}

func (s *Server) Start() error {
	s.logger.Info("starting HTTP server", slog.String("addr", s.server.Addr))
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.logger.Info("shutting down HTTP server")
	return s.server.Shutdown(ctx)
}
