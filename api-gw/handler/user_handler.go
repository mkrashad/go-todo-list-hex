package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mkrashad/go-todo/api-gw/pb"
)

type UserHandler struct {
	userClient pb.UserServiceClient
}

func NewUserHandler(userClient pb.UserServiceClient) *UserHandler {
	return &UserHandler{userClient}
}

func (uh *UserHandler) GetAllUsers(c *gin.Context) {
	users, _ := uh.userClient.GetAllUsers(c.Request.Context(), &pb.GetAllUsersRequest{})
	c.JSON(200, users)
}

func (uh *UserHandler) GetUserById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.Status(400)
		return
	}

	user, err := uh.userClient.GetUserById(c.Request.Context(), &pb.GetUserByIdRequest{Id: int64(id)})
	if err == nil {
		c.JSON(200, user)
	}
	log.Print("User not found")
	c.Status(http.StatusNotFound)
}

func (uh *UserHandler) CreateUser(c *gin.Context) {
	var input pb.CreateUserRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uh.userClient.CreateUser(c.Request.Context(), &input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, user)
}

func (uh *UserHandler) UpdateUserById(c *gin.Context) {
	var input pb.UpdateUserRequest

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.Id = int64(id)

	u, err := uh.userClient.UpdateUser(c.Request.Context(), &input)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, u)
}

func (uh *UserHandler) DeleteUserById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.Status(400)
		return
	}

	uh.userClient.DeleteUser(c.Request.Context(), &pb.DeleteUserRequest{Id: int64(id)})

	c.Status(204)
}
