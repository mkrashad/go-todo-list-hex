package task

import "log"


//go:generate mockery --name Service
type Service interface {
	GetAllTasks() []Task
	GetTaskById(id uint64) (Task, error)
	CreateTask(task Task) (Task, error)
	UpdateTaskById(id uint64, data Task) (Task, error)
	DeleteTaskById(id uint64) error
}

type taskService struct {
	repository Repository
}

func (ts taskService) GetAllTasks() []Task {
	return ts.repository.GetAllTasks()
}

func (ts taskService) GetTaskById(id uint64) (Task, error) {
	task, err := ts.repository.GetTaskById(id)

	if err != nil {
		log.Printf("Error %d id doesn't exist. %s", id, err)
	}
	return task, err
}

func (ts taskService) CreateTask(task Task) (Task, error) {
	task, err := ts.repository.CreateTask(task)
	if err != nil {
		log.Printf("An error occur while creating the task: %s\n", err)
	}
	return task, err
}

func (ts taskService) UpdateTaskById(id uint64, task Task) (Task, error) {
	task, err := ts.repository.UpdateTaskById(id, task)
	if err != nil {
		log.Printf("Could not update the task: %s\n", err)
	}
	return task, err
}

func (ts taskService) DeleteTaskById(id uint64) error {
	err := ts.repository.DeleteTaskById(id)
	if err != nil {
		log.Println("Something went wrong could not delete task", err)
	}
	return err
}

// Dependency injection
func NewTaskService(repository Repository) Service {
	return &taskService{repository}
}
