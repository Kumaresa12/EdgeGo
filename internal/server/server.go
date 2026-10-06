package server

import (
	"context"
	"fmt"
	"github.com/rdwr-kumaresans/EdgeGo/internal/config"
	"github.com/rdwr-kumaresans/EdgeGo/internal/handler"
	"github.com/rdwr-kumaresans/EdgeGo/internal/logger"
	"net/http"
)

type Server struct {
	httpServer *http.Server
	logger     logger.Logger
}

func New(
	cfg *config.Config,
	log logger.Logger,
) *Server {
	mux := http.NewServeMux()

	rootHandler := handler.NewRootHandler()
	healthHandler := handler.NewHealthHandler()
	versionHandler := handler.NewVersionHandler()

	mux.Handle("/", rootHandler)
	mux.Handle("/health", healthHandler)
	mux.Handle("/version", versionHandler)

	address := fmt.Sprintf(
		"%s:%d",
		cfg.Server.Host,
		cfg.Server.Port,
	)

	httpServer := &http.Server{
		Addr:         address,
		Handler:      mux,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}

	return &Server{
		httpServer: httpServer,
		logger:     log,
	}
}

func (s *Server) Start() error {
	s.logger.Info(
		"Starting HTTP server",
		"address", s.httpServer.Addr,
	)

	err := s.httpServer.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("HTTP server failed: %w", err)
	}
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.logger.Info("Shutting down HTTP server")

	if err := s.httpServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("HTTP server shutdown failed: %w", err)
	}

	s.logger.Info("HTTP server stopped")

	return nil
}
