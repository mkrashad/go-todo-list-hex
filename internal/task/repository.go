package task

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetAllTasks() []Task
	GetTaskById(id uint64) (Task, error)
	CreateTask(task Task) (Task, error)
	UpdateTaskById(id uint64, task Task) (Task, error)
	DeleteTaskById(id uint64) error
}

type taskRepository struct {
	DB *gorm.DB
}

// Dependency injection
func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{DB: db}
}

func (tr taskRepository) GetAllTasks() []Task {
	var tasks []Task
	tr.DB.Find(&tasks)
	return tasks
}

func (tr taskRepository) GetTaskById(id uint64) (Task, error) {
	var task Task
	result := tr.DB.Find(&task, id)
	return task, result.Error
}

func (tr taskRepository) CreateTask(task Task) (Task, error) {
	result := tr.DB.Create(&task)
	tr.DB.Save(&task)
	return task, result.Error
}

func (tr taskRepository) UpdateTaskById(id uint64, data Task) (Task, error) {
	var task Task
	tr.DB.Find(&task, id)
	result := tr.DB.Model(&task).Updates(data)
	return task, result.Error
}

func (tr taskRepository) DeleteTaskById(id uint64) error {
	result := tr.DB.Delete(&Task{}, id)
	return result.Error
}
