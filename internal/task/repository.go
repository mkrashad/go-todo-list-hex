package task

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetAllTasks() []Task
	GetTaskById(id uint64) (Task, error)
	CreateTask(task Task) (Task, error)
	UpdateTaskById(id uint64, data Task) (Task, error)
	DeleteTaskById(id uint64) error
}

type taskRepository struct {
	DB *gorm.DB
}

type taskGorm struct {
	gorm.Model
	Task
}

// Dependency injection
func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{DB: db}
}

func (tr taskRepository) GetAllTasks() []taskGorm {
	var tasks []taskGorm
	tr.DB.Find(&tasks)
	return tasks
}

func (tr taskRepository) GetTaskById(id uint64) (taskGorm, error) {
	var task taskGorm
	result := tr.DB.Find(&task, id)
	return task, result.Error
}

func (tr taskRepository) CreateTask() (taskGorm, error) {
	var task taskGorm
	result := tr.DB.Create(&task)
	return task, result.Error
}

func (tr taskRepository) UpdateTaskById(id uint64, data Task) (taskGorm, error) {
	var task taskGorm
	tr.DB.Find(&task, id)
	result := tr.DB.Model(&task).Updates(data)
	return task, result.Error
}

func (tr taskRepository) DeleteTaskById(id uint64) error {
	result := tr.DB.Delete(&Task{}, id)
	return result.Error
}
