package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mkrashad/go-todo/api-gw/handler/auth"
	"github.com/mkrashad/go-todo/api-gw/pb"
)

type AuthHandler struct {
	userClient pb.UserServiceClient
}

func NewAuthHandler(userClient pb.UserServiceClient) *AuthHandler {
	return &AuthHandler{
		userClient: userClient,
	}
}

func (ah *AuthHandler) Login(c *gin.Context) {
	var input pb.GetByUserNameAndPasswordRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := ah.userClient.GetByUserNameAndPassword(c.Request.Context(), &pb.GetByUserNameAndPasswordRequest{
		Username: input.Username,
		Password: input.Password,
	})


	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Login or password is invalid"})
		return
	}

	jwt, err := auth.GenerateToken(u.User)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Header("Authorization", jwt)
	c.JSON(http.StatusOK, u)
}

func (ah *AuthHandler) Register(c *gin.Context) {
	var input pb.CreateUserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := ah.userClient.CreateUser(c.Request.Context(), &input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"msg": "Successfully registered"})
}
