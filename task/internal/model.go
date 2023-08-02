package task

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	TaskName      string `json:"task_name"`
	Completed bool   `json:"completed"`
	UserID    uint64 `json:"user_id"`
}
