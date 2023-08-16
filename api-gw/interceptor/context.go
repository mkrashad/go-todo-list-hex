package interceptor

import (
	"context"
	"github.com/mkrashad/go-todo/api-gw/ctxutils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func ContextPropagation() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req,
		reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		md := metadata.Pairs(
			ctxutils.RequestIdKey, ctxutils.GetRequestId(ctx),
		)
		ctx = metadata.NewOutgoingContext(ctx, md)
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
