package utils

import (
	"os"
	"sync"

	"go.uber.org/zap"
)

var (
	logger     *zap.Logger
	loggerOnce sync.Once
)

func initializeLogger() {
	if os.Getenv("APP_ENV") == "develop" {
		logger = zap.Must(zap.NewDevelopment())
	}
	logger = zap.Must(zap.NewProduction())
}

func GetLogger() *zap.Logger {
	loggerOnce.Do(initializeLogger)
	return logger
}
