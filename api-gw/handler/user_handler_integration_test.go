package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mkrashad/go-todo/user/internal"
	userMocks "github.com/mkrashad/go-todo/user/internal/mocks"
	"github.com/stretchr/testify/suite"
)

type UserHandlerIntegrationTestSuite struct {
	suite.Suite
	underTest       *UserHandler
	mockUserService *userMocks.Service
	router          *gin.Engine
}

func (ts *UserHandlerIntegrationTestSuite) SetupSuite() {
	ts.mockUserService = new(userMocks.Service)
	ts.underTest = NewUserHandler(ts.mockUserService)
	ts.router = gin.Default()
	ts.router.GET("/users", ts.underTest.GetAllUsers)
	ts.router.GET("/users/:id", ts.underTest.GetUserById)
	ts.router.PUT("/users", ts.underTest.UpdateUserById)
	ts.router.POST("/users", ts.underTest.CreateUser)
	ts.router.DELETE("/users/:id", ts.underTest.DeleteUserById)
}

func TestUserHandlerIntegrationTestSuite(t *testing.T) {
	suite.Run(t, &UserHandlerIntegrationTestSuite{})
}

var users = []internal.User{
	{

		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@mail.com",
	},
	{
		FirstName: "Adam",
		LastName:  "Smith",
		Email:     "adam.smith@mail.com",
	},
}

func (ts *UserHandlerIntegrationTestSuite) TestGetAllUsers() {
	// given
	ts.mockUserService.On("GetAllUsers").Once().Return(users)
	// when
	req, _ := http.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()
	ts.router.ServeHTTP(w, req)

	response, _ := io.ReadAll(w.Body)
	// then
	var responseUsers []internal.User
	err := json.Unmarshal(response, &responseUsers)
	if err != nil {
		ts.Fail("Failed to convert")
	}
	ts.Equal(users, responseUsers)
	ts.Equal(http.StatusOK, w.Code)
	ts.mockUserService.AssertExpectations(ts.T())
}

func (ts *UserHandlerIntegrationTestSuite) TestGetUserById_ValidId() {
	// given
	var id uint64 = 1
	ts.mockUserService.On("GetUserById", id).Once().Return(users[0], nil)
	// when
	req, _ := http.NewRequest("GET", "/users/1", nil)
	w := httptest.NewRecorder()
	ts.router.ServeHTTP(w, req)

	response, _ := io.ReadAll(w.Body)
	// then
	var responseUser internal.User
	err := json.Unmarshal(response, &responseUser)
	if err != nil {
		ts.Fail("Failed to convert")
	}
	ts.Equal(users[0], responseUser)
	ts.Equal(http.StatusOK, w.Code)
	ts.mockUserService.AssertExpectations(ts.T())
}

func (ts *UserHandlerIntegrationTestSuite) TestGetUserById_InvalidId() {
	// when
	req, _ := http.NewRequest("GET", "/users/-1", nil)
	w := httptest.NewRecorder()
	ts.router.ServeHTTP(w, req)
	// then
	ts.Equal(http.StatusBadRequest, w.Code)
	ts.mockUserService.AssertExpectations(ts.T())
}

func (ts *UserHandlerIntegrationTestSuite) TestDeleteUserById_ValidId() {
	// given
	var id uint64 = 1
	ts.mockUserService.On("DeleteUserById", id).Once().Return(nil)
	// when
	req, _ := http.NewRequest("DELETE", "/users/1", nil)
	w := httptest.NewRecorder()
	ts.router.ServeHTTP(w, req)
	// then
	ts.Equal(http.StatusNoContent, w.Code)
	ts.mockUserService.AssertExpectations(ts.T())
}

func (ts *UserHandlerIntegrationTestSuite) TestDeleteUserById_InvalidId() {
	// when
	req, _ := http.NewRequest("DELETE", "/users/-1", nil)
	w := httptest.NewRecorder()
	ts.router.ServeHTTP(w, req)
	// then
	ts.Equal(http.StatusBadRequest, w.Code)
	ts.mockUserService.AssertExpectations(ts.T())
}

func (ts *UserHandlerIntegrationTestSuite) TestCreateUser_Valid() {
	// given
	newUser := internal.User{
		FirstName: "Bob",
		LastName:  "White",
		Email:     "bob@mail.com",
	}
	jsonValue, _ := json.Marshal(newUser)
	ts.mockUserService.On("CreateUser").Once().Return(newUser, nil)
	// when
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	ts.router.ServeHTTP(w, req)

	response, _ := io.ReadAll(w.Body)
	// then
	var responseUser internal.User
	err := json.Unmarshal(response, &responseUser)
	if err != nil {
		ts.Fail("Failed to convert")
	}
	ts.Equal(newUser, responseUser)
	ts.Equal(http.StatusCreated, w.Code)
	ts.mockUserService.AssertExpectations(ts.T())
}
