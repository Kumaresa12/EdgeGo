package app

import (
	"context"

	"github.com/rdwr-kumaresans/EdgeGo/internal/config"
	"github.com/rdwr-kumaresans/EdgeGo/internal/logger"
	"github.com/rdwr-kumaresans/EdgeGo/internal/server"
)

type App struct {
	Config *config.Config
	Logger logger.Logger
	Server *server.Server
}

func New(cfg *config.Config, log logger.Logger) *App {
	httpServer := server.New(cfg, log)

	return &App{
		Config: cfg,
		Logger: log,
		Server: httpServer,
	}
}

func (a *App) Run() error {
	return a.Server.Start()
}

func (a *App) Shutdown(ctx context.Context) error {
	return a.Server.Shutdown(ctx)
}
