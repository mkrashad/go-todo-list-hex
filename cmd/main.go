package main

import (
	"log"

	"github.com/mkrashad/go-todo/api"
	initilalizers "github.com/mkrashad/go-todo/database"
	"github.com/mkrashad/go-todo/internal/task"
)

func init() {
	initilalizers.LoadEnvVariables()
	initilalizers.ConnectToDB()
	initilalizers.SyncDB()
}

func main() {
	taskRepository := task.NewTaskRepository(initilalizers.DB)
	taskService := task.NewTaskService(taskRepository)
	taskHandler := api.NewTaskHandler(taskService)

	r := api.Router(taskHandler)
	err := r.Run()
	if err != nil {
		log.Fatal("Could not run application")
	}

}
