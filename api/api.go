package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mkrashad/go-todo/api/metrics"
)

func Router(taskHandler *TaskHandler, userHandler *UserHandler) *gin.Engine {
	r := gin.Default()

	metrics.PrometheusMetrics(r)

	// Tasks
	r.GET("/tasks", taskHandler.GetAllTasks)
	r.GET("/tasks/:id", taskHandler.GetTaskById)
	r.POST("/tasks", taskHandler.CreateTask)
	r.PUT("/tasks/:id", taskHandler.UpdateTaskById)
	r.DELETE("/tasks/:id", taskHandler.DeleteTaskById)

	// Users
	r.GET("/users", userHandler.GetAllUsers)
	r.GET("/users/:id", userHandler.GetUserById)
	r.POST("/users", userHandler.CreateUser)
	r.PUT("/users/:id", userHandler.UpdateUserById)
	r.DELETE("/users/:id", userHandler.DeleteUserById)

	return r
}
