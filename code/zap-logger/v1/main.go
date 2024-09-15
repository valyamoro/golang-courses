package main

import "go.uber.org/zap"

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Info("Starting the application")
	logger.Warn("This is a warning message")
	logger.Error("An error occurred")
}
