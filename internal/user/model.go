package user

import (
	"github.com/mkrashad/go-todo/internal/task"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Tasks     []task.Task
}
