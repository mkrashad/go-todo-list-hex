package main

import (
	"log"
	"net"

	database "github.com/mkrashad/go-todo/user/database"
	"github.com/mkrashad/go-todo/user/cmd/server"
	user "github.com/mkrashad/go-todo/user/internal"
	"github.com/mkrashad/go-todo/user/pb"
	"google.golang.org/grpc"
)

func init() {
	database.LoadEnvVariables()
	database.ConnectToDB()
	database.SyncDB()
}

func main() {
	// Users
	userRepository := user.NewUserRepository(database.DB)
	userService := user.NewUserService(userRepository)
	srv := server.NewServer(userService)


	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("Can't create listen: %s", err)
	}

	s := grpc.NewServer()

	pb.RegisterUserServiceServer(s, srv)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
