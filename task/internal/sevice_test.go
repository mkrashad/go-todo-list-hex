package task_test

import (
	"errors"
	"testing"

	"github.com/mkrashad/go-todo/internal/task"
	"github.com/mkrashad/go-todo/internal/task/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TaskServiceUnitTestSuite struct {
	suite.Suite
	underTest          task.Service
	mockTaskRepository *mocks.Repository
}

var tasks = []task.Task{
	{
		Name:      "Study for exam",
		Completed: false,
	},
	{
		Name:      "Play soccer",
		Completed: true,
	},
}

var id uint64 = 1

func TestTaskServiceUnitTestSuite(t *testing.T) {
	suite.Run(t, &TaskServiceUnitTestSuite{})
}

func (ts *TaskServiceUnitTestSuite) SetupSuite() {
	ts.mockTaskRepository = new(mocks.Repository)
	ts.underTest = task.NewTaskService(ts.mockTaskRepository)
}

func (ts *TaskServiceUnitTestSuite) TestGetAllTasks() {
	// given
	ts.mockTaskRepository.On("GetAllTasks").Once().Return(tasks)

	// when
	result := ts.underTest.GetAllTasks()

	//then
	ts.Equal(tasks, result)
	ts.mockTaskRepository.AssertExpectations(ts.T())

}

func (ts *TaskServiceUnitTestSuite) TestGetTaskById_ValidId() {
	// given
	ts.mockTaskRepository.On("GetTaskById", id).Once().Return(tasks[0], nil)

	// when
	result, err := ts.underTest.GetTaskById(id)

	// then
	ts.NoError(err)
	ts.Equal(tasks[0], result)
	ts.mockTaskRepository.AssertExpectations(ts.T())

}

func (ts *TaskServiceUnitTestSuite) TestGetTaskById_InvalidId() {
	// given
	ts.mockTaskRepository.On("GetTaskById", id).Once().Return(task.Task{}, errors.New("not found"))

	// when
	result, err := ts.underTest.GetTaskById(id)

	// then
	ts.Error(err)
	ts.Zero(result)
	ts.mockTaskRepository.AssertExpectations(ts.T())
}

func (ts *TaskServiceUnitTestSuite) TestCreateTaskValid() {
	// given
	newTask := task.Task{
		Name:      "new task",
		Completed: false,
	}
	ts.mockTaskRepository.On("CreateTask", newTask).Once().Return(newTask, nil)

	// when
	result, err := ts.underTest.CreateTask(newTask)

	// then
	ts.NoError(err)
	ts.Equal(newTask, result)
	ts.mockTaskRepository.AssertExpectations(ts.T())
}

func (ts *TaskServiceUnitTestSuite) TestCreateTaskInValid() {
	// given
	newTask := task.Task{
		Name:      "new task",
		Completed: false,
	}
	ts.mockTaskRepository.On("CreateTask", newTask).Once().Return(newTask, errors.New("error"))

	// when
	_, err := ts.underTest.CreateTask(newTask)
	
	// then
	ts.Error(err)
	ts.mockTaskRepository.AssertExpectations(ts.T())
}

func (ts *TaskServiceUnitTestSuite) TestUpdateTaskById_ValidUpdate() {
	// given
	updatedTask := tasks[0]
	updatedTask.Completed = true
	updatedTask.Name = "Read a book"
	ts.mockTaskRepository.On("UpdateTaskById", id, updatedTask).Once().Return(updatedTask, nil)

	// when
	result, err := ts.underTest.UpdateTaskById(id, updatedTask)

	// then
	ts.NoError(err)
	ts.Equal(updatedTask, result)
	ts.mockTaskRepository.AssertExpectations(ts.T())
}

func (ts *TaskServiceUnitTestSuite) TestUpdateTaskById_InvalidUpdate() {
	// given
	updatedTask := tasks[0]
	updatedTask.Completed = true
	updatedTask.Name = "Read a book"
	ts.mockTaskRepository.On("UpdateTaskById", id, mock.Anything).Once().Return(task.Task{}, errors.New("error"))

	// when
	result, err := ts.underTest.UpdateTaskById(id, updatedTask)

	//then
	ts.Error(err)
	ts.Zero(result)
	ts.mockTaskRepository.AssertExpectations(ts.T())
}

func (ts *TaskServiceUnitTestSuite) TestDeleteTaskById_Valid() {
	// given
	ts.mockTaskRepository.On("DeleteTaskById", id).Once().Return(nil)

	// when
	ts.underTest.DeleteTaskById(id)

	// then
	ts.mockTaskRepository.AssertExpectations(ts.T())

}

func (ts *TaskServiceUnitTestSuite) TestDeleteTaskById_Invalid() {
	// given
	ts.mockTaskRepository.On("DeleteTaskById", id).Once().Return(errors.New("error"))

	// when
	err := ts.underTest.DeleteTaskById(id)

	// then
	ts.Error(err)
	ts.mockTaskRepository.AssertExpectations(ts.T())

}

// Benchmarks
func BenchmarkTaskService_GetAllTasks(b *testing.B) {
	mockTaskRepository := new(mocks.Repository)
	underTest := task.NewTaskService(mockTaskRepository)

	mockTaskRepository.On("GetAllTasks").Return(tasks)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		underTest.GetAllTasks()
	}
}

func BenchmarkTaskService_GetTaskById(b *testing.B) {
	mockTaskRepository := new(mocks.Repository)
	underTest := task.NewTaskService(mockTaskRepository)

	mockTaskRepository.On("GetTaskById", id).Return(tasks[0], nil)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		underTest.GetTaskById(id)
	}
}

func BenchmarkTaskService_CreateTask(b *testing.B) {
	mockTaskRepository := new(mocks.Repository)
	underTest := task.NewTaskService(mockTaskRepository)

	newTask := task.Task{
		Name:      "new task",
		Completed: false,
	}

	mockTaskRepository.On("CreateTask", newTask).Return(newTask, nil)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		underTest.CreateTask(newTask)
	}
}

func BenchmarkTaskService_UpdateTaskById(b *testing.B) {
	mockTaskRepository := new(mocks.Repository)
	underTest := task.NewTaskService(mockTaskRepository)

	updatedTask := tasks[0]
	updatedTask.Completed = true
	updatedTask.Name = "Read a book"

	mockTaskRepository.On("UpdateTaskById", id, mock.Anything).Return(updatedTask, nil)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		underTest.UpdateTaskById(id, updatedTask)
	}
}

func BenchmarkTaskService_DeleteTaskById(b *testing.B) {
	mockTaskRepository := new(mocks.Repository)
	underTest := task.NewTaskService(mockTaskRepository)

	mockTaskRepository.On("DeleteTaskById", id).Once().Return(nil)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		underTest.DeleteTaskById(id)
	}

}
