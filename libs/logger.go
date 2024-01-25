package libs

import (
	"os"

	"go.uber.org/zap"
)

func NewLogger() *zap.Logger {
	logger, _ := zap.NewProduction(
		zap.AddCaller(),
		zap.AddStacktrace(zap.ErrorLevel),
	)

	logger.With(zap.String("APP_NAME", os.Getenv("APP_NAME")))
	return logger
}
