package logger

import "go.uber.org/zap"

func New() *zap.Logger {
	zapLogger, _ := zap.NewProduction()

	return zapLogger
}
