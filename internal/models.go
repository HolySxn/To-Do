package server

import (
	"time"
)

type List struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Position  int       `json:"position"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Task struct {
	ID        string    `json:"id"`
	ListID    string    `json:"list_id"`
	TaskName  string    `json:"task_name"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SubTask struct {
	ID          string    `json:"id"`
	TaskID      string    `json:"task_id"`
	SubTaskName string    `json:"subtask_name"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
