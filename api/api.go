package api

import (
	"github.com/gin-gonic/gin"
)

func Router(taskHandler *TaskHandler) *gin.Engine {
	r := gin.Default()
	// Tasks
	r.GET("/tasks", taskHandler.GetAllTasks)
	r.GET("/tasks/:id", taskHandler.GetTaskById)
	r.POST("/tasks", taskHandler.CreateTask)
	r.PUT("/tasks/:id", taskHandler.UpdateTaskById)
	r.DELETE("/tasks/:id", taskHandler.DeleteTaskById)

	// Users
	// r.POST("/users", controllers.CreateUser)
	// r.PUT("/users/:id", controllers.UpdateUser)
	// r.GET("/users", controllers.GetAllUsers)
	// r.GET("/users/:id", controllers.GetUserById)
	// r.DELETE("/users/:id", controllers.DeleteUser)
	// r.Run()
	return r
}
