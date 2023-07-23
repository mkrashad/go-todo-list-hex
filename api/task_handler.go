package api

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mkrashad/go-todo/internal/task"
)

type TaskHandler struct {
	taskService task.Service
}

func NewTaskHandler(taskService task.Service) *TaskHandler {
	return &TaskHandler{taskService}
}

func (th *TaskHandler) GetAllTasks(c *gin.Context) {
	task := th.taskService.GetAllTasks()
	c.JSON(200, gin.H{
		"task": task,
	})
}

func (th *TaskHandler) GetTaskById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.Status(400)
		return
	}

	task, err := th.taskService.GetTaskById(id)
	if err != nil {
		c.Status(404)
		log.Print("Something went wrong")
	}

	c.JSON(200, task)
}

func (th *TaskHandler) CreateTask(c *gin.Context) {
	var input task.Task

	c.Bind(&input)

	fmt.Print(input)
	task, err := th.taskService.CreateTask(input)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Something went wrong",
		})
	}

	c.JSON(200, gin.H{
		"tasks": task,
	})
}

func (th *TaskHandler) UpdateTaskById(c *gin.Context) {
	var input task.Task
	c.Bind(&input)

	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	tasks, err := th.taskService.UpdateTaskById(id, input)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Something went wrong",
		})
	}

	c.JSON(200, tasks)
}

func (th *TaskHandler) DeleteTaskById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Something went wrong",
		})
	}

	th.taskService.DeleteTaskById(id)

	c.Status(204)
}
