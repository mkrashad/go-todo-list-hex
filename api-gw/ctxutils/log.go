package ctxutils
 
import (
    "context"
    "go.uber.org/zap"
)
 
type loggerContextKey struct{}
 
func SetLogger(ctx context.Context) context.Context {
    logger, _ := zap.NewProduction()
    return context.WithValue(ctx, loggerContextKey{}, logger)
}
 
func GetLogger(ctx context.Context) *zap.Logger {
    logger := ctx.Value(loggerContextKey{}).(*zap.Logger)
    if logger == nil {
        panic("logger is not defined!")
    }
 
    return logger
}