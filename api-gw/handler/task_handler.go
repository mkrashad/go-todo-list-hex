package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mkrashad/go-todo/api-gw/pb"
)

type TaskHandler struct {
	taskClient pb.TaskServiceClient
	userClient pb.UserServiceClient
}

func NewTaskHandler(taskClient pb.TaskServiceClient, userClient pb.UserServiceClient) *TaskHandler {
	return &TaskHandler{
		taskClient: taskClient,
		userClient: userClient,
	}
}

func (th *TaskHandler) GetAllTasks(c *gin.Context) {
	tasks, _ := th.taskClient.GetAllTasks(c.Request.Context(), &pb.GetAllTasksRequest{})
	c.JSON(200, tasks)
}

func (th *TaskHandler) GetTaskById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.Status(400)
		return
	}

	task, err := th.taskClient.GetTaskById(c.Request.Context(), &pb.GetTaskByIdRequest{Id: int64(id)})
	if err == nil {
		c.JSON(200, task)
	}
	log.Print("Task not found")
	c.Status(404)
}

func (th *TaskHandler) CreateTask(c *gin.Context) {
	var input pb.CreateTaskRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := th.taskClient.CreateTask(c.Request.Context(), &input)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Something went wrong",
		})
	}

	c.JSON(201, task)
}

func (th *TaskHandler) UpdateTaskById(c *gin.Context) {
	var input pb.UpdateTaskRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	task, err := th.taskClient.UpdateTask(c.Request.Context(), &input)
	if err != nil {
		c.Status(404)
		return
	}
	c.JSON(200, task)
}

func (th *TaskHandler) DeleteTaskById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.Status(400)
		return
	}

	th.taskClient.DeleteTask(c.Request.Context(), &pb.DeleteTaskRequest{Id: int64(id)})

	c.Status(204)
}
