package main

import (
	"log"
	"github.com/mkrashad/go-todo/api-gw/pb"
	"github.com/mkrashad/go-todo/api-gw/router"
	"github.com/mkrashad/go-todo/api-gw/handler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	tc, err := grpc.Dial("localhost:8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Could not connect to task microservice: ", err)
	}
	taskClient := pb.NewTaskServiceClient(tc)

	uc, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Could not connect to user microservice: ", err)
	}
	userClient := pb.NewUserServiceClient(uc)

	taskHandler := handler.NewTaskHandler(taskClient, userClient)
	userHandler := handler.NewUserHandler(userClient)
	authHandler := handler.NewAuthHandler(userClient)

	r := router.NewRouter(*taskHandler, *userHandler, *authHandler)
	// r := router.NewRouter(*taskHandler, *userHandler)

	// run api gw
	if err := r.Run(); err != nil {
		log.Fatal("Could not start router: ", err)
	}
}
