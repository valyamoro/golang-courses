package main

import (
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	sugar := logger.Sugar()
	userID := 42
	requestID := "123abc"

	sugar.Infow("Handling request",
		"userID", userID,
		"requestID", requestID,
	)

	sugar.Errorw("Failed to process request",
		"userID", userID,
		"requestID", requestID,
		"error", "some error details",
	)
}
