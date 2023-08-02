package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mkrashad/go-todo/internal/task"
	"github.com/mkrashad/go-todo/internal/user"
)

type TaskHandler struct {
	taskService task.Service
	userService user.Service
}

func NewTaskHandler(taskService task.Service, userService user.Service) *TaskHandler {
	return &TaskHandler{taskService, userService}
}

func (th *TaskHandler) GetAllTasks(c *gin.Context) {
	tasks := th.taskService.GetAllTasks()
	c.JSON(200, tasks)
}

func (th *TaskHandler) GetTaskById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.Status(400)
		return
	}

	task, err := th.taskService.GetTaskById(id)
	if err == nil {
		c.JSON(200, task)
	}
	log.Print("Task not found")
	c.Status(404)
}

func (th *TaskHandler) CreateTask(c *gin.Context) {
	var input task.Task

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := th.taskService.CreateTask(input)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Something went wrong",
		})
	}

	c.JSON(201, task)
}

func (th *TaskHandler) UpdateTaskById(c *gin.Context) {
	var input task.Task
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	task, err := th.taskService.UpdateTaskById(id, input)
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

	th.taskService.DeleteTaskById(id)

	c.Status(204)
}
