package main

import (
	"os"
	"os/signal"
	"syscall"

	"gotu-bookstore/internal/provider"

	"go.uber.org/zap"
)

func main() {
	logger := provider.LoggerProvider()
	defer logger.Sync()

	app := provider.ProvideApp()

	go func() {
		err := app.Run()
		if err != nil {
			logger.Fatal("app failed to run", zap.Error(err))
		}
	}()

	// listen for interrupt/terminate signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	logger.Info("shutting down...")

	if err := app.Stop(); err != nil {
		logger.Fatal("error when shutting down", zap.Error(err))
	}

	logger.Info("good bye.")
}
