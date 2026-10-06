package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rdwr-kumaresans/EdgeGo/internal/app"
	"github.com/rdwr-kumaresans/EdgeGo/internal/config"
	"github.com/rdwr-kumaresans/EdgeGo/internal/logger"
)

func main() {
	cfg, err := config.Load("configs/edgego.yaml")

	if err != nil {
		log.Fatal(err)
	}

	logr := logger.New()

	application := app.New(cfg, logr)

	serverErr := make(chan error, 1)

	go func() {
		serverErr <- application.Run()
	}()

	signalCtx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	select {

	case err := <-serverErr:
		if err != nil {
			logr.Error(
				"Server stopped unexceptedly",
				"error", err,
			)
			os.Exit(1)
		}

	case <-signalCtx.Done():
		logr.Info("Shutdown signal received")

		shutdownCtx, cancel := context.WithTimeout(
			context.Background(),
			10*time.Second,
		)
		defer cancel()

		if err := application.Shutdown(shutdownCtx); err != nil {
			logr.Error("Graceful shutdown failed",
				"error", err,
			)
			os.Exit(1)
		}

	}

}
