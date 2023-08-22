package internal

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	TaskName  string `json:"task_name"`
	Completed *bool   `json:"completed" gorm:"default:false"`
	UserID    uint64 `json:"user_id"`
}
