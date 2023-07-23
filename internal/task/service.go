package task

import "log"

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

// Dependency injection
func NewTaskService(repository Repository) *taskService {
	return &taskService{repository}
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
		log.Println("An error occur while creating the task", err)
	}
	return task, err
}

func (ts taskService) UpdateTaskById(id uint64, task Task) (Task, error) {
	task, err := ts.repository.UpdateTaskById(id, task)
	if err != nil {
		log.Println("Could not update the task", err)
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
