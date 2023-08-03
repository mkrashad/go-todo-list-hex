package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mkrashad/go-todo/task/internal"
	taskMocks "github.com/mkrashad/go-todo/task/internal/mocks"
	"github.com/stretchr/testify/suite"
)

type TaskHandlerIntegrationTestSuite struct {
	suite.Suite
	underTest       *TaskHandler
	mockTaskService *taskMocks.Service
	router          *gin.Engine
}

func TestTaskHandlerIntegrationTestSuite(t *testing.T) {
	suite.Run(t, &TaskHandlerIntegrationTestSuite{})
}

var tasks = []internal.Task{
	{
		TaskName:      "Study for exam",
		Completed: false,
	},
	{
		TaskName:      "Play soccer",
		Completed: true,
	},
}

func (ts *TaskHandlerIntegrationTestSuite) SetupSuite() {
	ts.mockTaskService = new(taskMocks.Service)
	ts.mockUserService = new(userMocks.Service)
	ts.underTest = NewTaskHandler(ts.mockTaskService, ts.mockUserService)
	ts.router = gin.Default()
	ts.router.GET("/tasks", ts.underTest.GetAllTasks)
	ts.router.GET("/tasks/:id", ts.underTest.GetTaskById)
	ts.router.PUT("/tasks/:id", ts.underTest.UpdateTaskById)
	ts.router.POST("/tasks", ts.underTest.CreateTask)
	ts.router.DELETE("/tasks/:id", ts.underTest.DeleteTaskById)
}

func (ts *TaskHandlerIntegrationTestSuite) TestGetAllTasks() {
	// given
	ts.mockTaskService.On("GetAllTasks").Once().Return(tasks)
	// when
	req, _ := http.NewRequest("GET", "/tasks", nil)
	w := httptest.NewRecorder()
	ts.router.ServeHTTP(w, req)

	response, _ := io.ReadAll(w.Body)
	// then
	var responseTasks []internal.Task
	err := json.Unmarshal(response, &responseTasks)
	if err != nil {
		ts.Fail("Failed to convert", err)
	}
	ts.Equal(tasks, responseTasks)
	ts.Equal(http.StatusOK, w.Code)
	ts.mockTaskService.AssertExpectations(ts.T())
}

func (ts *TaskHandlerIntegrationTestSuite) TestGetTaskBy_ValidId() {
	// given
	var id uint64 = 1
	ts.mockTaskService.On("GetTaskById", id).Once().Return(tasks[0], nil)
	// when
	req, _ := http.NewRequest("GET", "/tasks/1", nil)
	w := httptest.NewRecorder()
	ts.router.ServeHTTP(w, req)
	response, _ := io.ReadAll(w.Body)
	// then
	var responseTask internal.Task
	err := json.Unmarshal(response, &responseTask)
	if err != nil {
		ts.Fail("Failed to convert")
	}
	ts.Equal(tasks[0], responseTask)
	ts.Equal(http.StatusOK, w.Code)
	ts.mockTaskService.AssertExpectations(ts.T())
}

func (ts *TaskHandlerIntegrationTestSuite) TestGetTaskById_InvalidId() {
	// given
	var id uint64 = 9999
	ts.mockTaskService.On("GetTaskById", id).Once().Return(internal.Task{}, errors.New("not found"))
	// when
	req, _ := http.NewRequest("GET", "/tasks/9999", nil)
	w := httptest.NewRecorder()
	ts.router.ServeHTTP(w, req)
	// then
	ts.Equal(http.StatusNotFound, w.Code)
	ts.mockTaskService.AssertExpectations(ts.T())
}

func (ts *TaskHandlerIntegrationTestSuite) TestUpdateTaskById_ValidBody() {
	// given
	var id uint64 = 1
	updatedTask := tasks[0]
	updatedTask.TaskName = "updated"
	updatedTask.Completed = true
	jsonValue, _ := json.Marshal(updatedTask)
	ts.mockTaskService.On("UpdateTaskById", id, updatedTask).Once().Return(updatedTask, nil)
	// when
	req, _ := http.NewRequest("PUT", "/tasks/1", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	ts.router.ServeHTTP(w, req)
	response, _ := io.ReadAll(w.Body)
	// then
	var responseTask internal.Task
	err := json.Unmarshal(response, &responseTask)
	if err != nil {
		ts.Fail("Failed to convert: ", string(response))
	}
	ts.Equal(updatedTask, responseTask)
	ts.Equal(http.StatusOK, w.Code)
	ts.mockTaskService.AssertExpectations(ts.T())
}

func (ts *TaskHandlerIntegrationTestSuite) TestUpdateTaskById_InvalidBody() {
	// given
	jsonValue := "invalid body"
	// when
	req, _ := http.NewRequest("PUT", "/tasks/1", bytes.NewBuffer([]byte(jsonValue)))
	w := httptest.NewRecorder()
	ts.router.ServeHTTP(w, req)
	response, _ := io.ReadAll(w.Body)
	// then
	var responseTask internal.Task
	err := json.Unmarshal(response, &responseTask)
	if err != nil {
		ts.Fail("Failed to convert: ", string(response))
	}
	ts.Equal(http.StatusBadRequest, w.Code)
	ts.mockTaskService.AssertExpectations(ts.T())
}

func (ts *TaskHandlerIntegrationTestSuite) TestDeleteTaskById_ValidId() {
	// given
	var id uint64 = 1
	ts.mockTaskService.On("DeleteTaskById", id).Once().Return(nil)
	// when
	req, _ := http.NewRequest("DELETE", "/tasks/1", nil)
	w := httptest.NewRecorder()
	ts.router.ServeHTTP(w, req)
	// then
	ts.Equal(http.StatusNoContent, w.Code)
	ts.mockTaskService.AssertExpectations(ts.T())
}

func (ts *TaskHandlerIntegrationTestSuite) TestCreateTask_ValidBody() {
	// given
	newTask := internal.Task{
		TaskName:      "Study for exam",
		Completed: true,
	}
	jsonValue, _ := json.Marshal(newTask)
	ts.mockTaskService.On("CreateTask", newTask).Once().Return(newTask, nil)
	// when
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	ts.router.ServeHTTP(w, req)
	response, _ := io.ReadAll(w.Body)
	// then
	var responseTask internal.Task
	err := json.Unmarshal(response, &responseTask)
	if err != nil {
		ts.Fail("Failed to convert: ", string(response))
	}
	ts.Equal(newTask, responseTask)
	ts.Equal(http.StatusCreated, w.Code)
	ts.mockTaskService.AssertExpectations(ts.T())
}
