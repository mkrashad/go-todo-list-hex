package main

import (
	"log"

	"github.com/mkrashad/go-todo/api-gw/handler"
	"github.com/mkrashad/go-todo/api-gw/pb"
	"github.com/mkrashad/go-todo/api-gw/router"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	tc, err := grpc.Dial("app-task:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Could not connect to task microservice: ", err)
	}
	taskClient := pb.NewTaskServiceClient(tc)

	uc, err := grpc.Dial("app-user:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Could not connect to user microservice: ", err)
	}
	userClient := pb.NewUserServiceClient(uc)

	taskHandler := handler.NewTaskHandler(taskClient, userClient)
	userHandler := handler.NewUserHandler(userClient)
	authHandler := handler.NewAuthHandler(userClient)

	r := router.NewRouter(*taskHandler, *userHandler, *authHandler)

	if err := r.Run(":8083"); err != nil {
		log.Fatal("Could not start router: ", err)
	}
}
