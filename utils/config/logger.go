package config

import (
	"log"

	"go.uber.org/zap"
)

var Logger *zap.Logger

func NewLogger(isProduction bool) *zap.Logger {
	var zapLogger *zap.Logger
	var err error
	if isProduction {
		zapLogger, err = zap.NewProduction()
		if err != nil {
			log.Println("Error while logging production")
			return nil
		}
	} else {
		zapLogger, err = zap.NewDevelopment()
		if err != nil {
			log.Println("Error while logging production")
			return nil
		}
	}
	Logger = zapLogger
	return Logger
}
func GetLogger() *zap.Logger {
	if Logger == nil {
		panic("logger not initialized. Call NewLogger() first")
	}
	return Logger
}
