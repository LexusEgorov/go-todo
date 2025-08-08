package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"time"

	"github.com/LexusEgorov/auth/internal/app"
	"github.com/LexusEgorov/auth/internal/config"
)

func main() {
	config, err := config.New()
	if err != nil {
		log.Fatalf("main: %v", err)
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: config.Logger.AddSource,
	}))

	app := app.New(logger, config)

	app.Run()

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt)

	<-stopChan
	logger.Info("Recieved interrupt signal")
	timeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	app.Stop(ctx)
}
