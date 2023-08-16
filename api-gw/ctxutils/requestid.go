package ctxutils

import "context"

const RequestIdKey = "request_id"

type requestIdContextKey struct{}

func SetRequestId(ctx context.Context, requestId string) context.Context {
	return context.WithValue(ctx, requestIdContextKey{}, requestId)
}

func GetRequestId(ctx context.Context) string {
	return ctx.Value(requestIdContextKey{}).(string)
}
