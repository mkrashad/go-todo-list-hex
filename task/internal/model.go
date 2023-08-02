package task

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
	UserID    uint64 `json:"user_id"`
}
