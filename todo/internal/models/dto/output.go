package dto

import (
	"time"
)

type TaskStatus string

const (
	TaskStatusNew        TaskStatus = "NEW"
	TaskStatusInProgress TaskStatus = "IN_PROGRESS"
	TaskStatusFinished   TaskStatus = "FINISHED"
	TaskStatusCancelled  TaskStatus = "CANCELLED"
)

type Tokens struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Login string `json:"login"`
}

type Task struct {
	ID        int        `json:"id"`
	UID       int        `json:"u_id"`
	Title     string     `json:"title"`
	Text      string     `json:"text"`
	Status    TaskStatus `json:"status"`
	Deadline  time.Time  `json:"deadline"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type ShortTask struct {
	ID     int        `json:"id"`
	Title  string     `json:"title"`
	Status TaskStatus `json:"status"`
}
