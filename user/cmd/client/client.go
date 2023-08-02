package main

import (
	"context"
	"log"

	"github.com/mkrashad/go-todo/user/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, _ := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))

	// err check

	// note: pb is the package of generated go files, it can differ based on the directory in which you generated your files
	client := pb.NewUserServiceClient(conn)

	users, _ := client.GetAllUsers(context.Background(), &pb.GetAllUsersRequest{})

	// err check

	log.Printf("users: %v", users)
}