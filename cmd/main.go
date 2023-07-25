package main

import (
	"log"

	"github.com/mkrashad/go-todo/api"
	initilalizers "github.com/mkrashad/go-todo/database"
	"github.com/mkrashad/go-todo/internal/task"
	"github.com/mkrashad/go-todo/internal/user"
)

func init() {
	initilalizers.LoadEnvVariables()
	initilalizers.ConnectToDB()
	initilalizers.SyncDB()
}

func main() {
	// Users
	userRepository := user.NewUserRepository(initilalizers.DB)
	userService := user.NewUserService(userRepository)
	userHandler := api.NewUserHandler(userService)

	// Tasks
	taskRepository := task.NewTaskRepository(initilalizers.DB)
	taskService := task.NewTaskService(taskRepository)
	taskHandler := api.NewTaskHandler(taskService, userService)

	r := api.Router(taskHandler, userHandler)
	err := r.Run()
	if err != nil {
		log.Fatal("Could not run application")
	}

}
