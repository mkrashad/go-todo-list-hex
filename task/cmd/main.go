package main

import (
	"log"
	"net"
	database "github.com/mkrashad/go-todo/task/database"
	"github.com/mkrashad/go-todo/task/cmd/server"
	task "github.com/mkrashad/go-todo/task/internal"
	"github.com/mkrashad/go-todo/task/pb"
	"google.golang.org/grpc"
)


func init() {
	database.ConnectToDB()
}
func main() {
	taskRepository := task.NewTaskRepository(database.DB)
	taskService := task.NewTaskService(taskRepository)
	srv := server.NewServer(taskService)

	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatalf("Can't create listen: %s", err)
	}
	s := grpc.NewServer()

	pb.RegisterTaskServiceServer(s, srv)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
