package main

import (
	"log"

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

	if err := application.Run(); err != nil {
		log.Fatal((err))
	}
}