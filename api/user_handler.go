package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mkrashad/go-todo/internal/user"
)

type UserHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *UserHandler {
	return &UserHandler{userService}
}

func (uh *UserHandler) GetAllUsers(c *gin.Context) {
	users := uh.userService.GetAllUsers()
	c.JSON(200, users)
}

func (uh *UserHandler) GetUserById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.Status(400)
		return
	}

	user, err := uh.userService.GetUserById(id)
	if err == nil {
		c.JSON(200, user)
	}
	log.Print("Task not found")
	c.Status(404)
}

func (uh *UserHandler) CreateUser(c *gin.Context) {
	var input user.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uh.userService.CreateUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, user)
}

func (uh *UserHandler) UpdateUserById(c *gin.Context) {
	var input user.User
	c.Bind(&input)

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	tasks, err := uh.userService.UpdateUserById(id, input)
	if err != nil {
		c.Status(404)
		return
	}

	c.JSON(200, tasks)
}

func (uh *UserHandler) DeleteUserById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.Status(400)
		return
	}

	uh.userService.DeleteUserById(id)

	c.Status(204)
}
