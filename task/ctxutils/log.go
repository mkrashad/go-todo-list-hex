package ctxutils

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.uber.org/zap"
)

type loggerContextKey struct{}

func SetLogger(ctx context.Context) context.Context {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	return context.WithValue(ctx, loggerContextKey{}, logger)
}

func GetSystemLogger(ctx context.Context) *zap.Logger {
	logger := ctx.Value(loggerContextKey{}).(*zap.Logger)
	if logger == nil {
		panic("logger is not defined!")
	}

	return logger
}

func GetRequestLogger(ctx context.Context) *zap.Logger {
	return ctxzap.Extract(ctx)
}
