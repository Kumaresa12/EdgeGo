package app

import (
	"github.com/rdwr-kumaresans/EdgeGo/internal/config"
	"github.com/rdwr-kumaresans/EdgeGo/internal/logger"
)

type App struct {
	Config *config.Config
	Logger logger.Logger
}

func New(cfg *config.Config, log logger.Logger) *App {
	return &App{
		Config: cfg,
		Logger: log,
	}
}

func (a *App) Run() error {
	a.Logger.Info("Starting Edgego")

	a.Logger.Info(
	"Configuration Loaded",
	"host", a.Config.Server.Host,
	"port", a.Config.Server.Port,
	)
	return nil
}