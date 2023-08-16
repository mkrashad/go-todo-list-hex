package main

import (
	"context"
	"log"
	"net"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/mkrashad/go-todo/user/cmd/server"
	database "github.com/mkrashad/go-todo/user/database"
	user "github.com/mkrashad/go-todo/user/internal"
	"github.com/mkrashad/go-todo/user/pb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"github.com/mkrashad/go-todo/user/ctxutils"
	"github.com/mkrashad/go-todo/user/interceptor"
)

var ctx context.Context
func init() {
	ctx = context.Background()
	ctx = ctxutils.SetLogger(ctx)

	database.ConnectToDB()
}


func main() {
	userRepository := user.NewUserRepository(database.DB)
	userService := user.NewUserService(userRepository)
	srv := server.NewServer(userService)


	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("Can't create listen: %s", err)
	}

	grpcLogger, err := zap.NewProduction()
	defer grpcLogger.Sync()

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_zap.UnaryServerInterceptor(grpcLogger),
				interceptor.ContextToZapFields(),
			)),
	}


	s := grpc.NewServer(opts...)
	pb.RegisterUserServiceServer(s, srv)
	if err := s.Serve(lis); err != nil {
		ctxutils.GetSystemLogger(ctx).Sugar().Fatal("failed to serve: ", err)
	}
}
