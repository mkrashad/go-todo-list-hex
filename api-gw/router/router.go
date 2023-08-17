package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mkrashad/go-todo/api-gw/handler"
	"github.com/mkrashad/go-todo/api-gw/metrics"
	"github.com/mkrashad/go-todo/api-gw/middleware"
)

func NewRouter(taskHandler handler.TaskHandler, userHandler handler.UserHandler, authHandler handler.AuthHandler) *gin.Engine {
	r := gin.Default()
	metrics.PrometheusMetrics(r)

	r.Use(middleware.RequestId())

	protected := r.Group("/api")
	protected.Use(middleware.JwtAuthMiddleware())

	// Tasks
	protected.GET("/tasks", taskHandler.GetAllTasks)
	protected.GET("/tasks/:id", taskHandler.GetTaskById)
	protected.POST("/tasks", taskHandler.CreateTask)
	protected.PUT("/tasks/:id", taskHandler.UpdateTaskById)
	protected.DELETE("/tasks/:id", taskHandler.DeleteTaskById)

	// Users
	protected.GET("/users", userHandler.GetAllUsers)
	protected.GET("/users/:id", userHandler.GetUserById)
	protected.POST("/users", userHandler.CreateUser)
	protected.PUT("/users/:id", userHandler.UpdateUserById)
	protected.DELETE("/users/:id", userHandler.DeleteUserById)

	public := r.Group("")
	public.POST("/login", authHandler.Login)
	public.POST("/register", authHandler.Register)

	return r
}
