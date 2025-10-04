package models

import (
	"time"

	"github.com/LexusEgorov/todo/internal/models/dto"
)

type Tokens struct {
	Access  string
	Refresh string
}

func (t Tokens) ToDTO() dto.Tokens {
	return dto.Tokens{
		Access:  t.Access,
		Refresh: t.Refresh,
	}
}

type User struct {
	ID        int
	TgID      int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u User) ToDTO() dto.User {
	return dto.User{
		ID:   u.ID,
		Name: u.Name,
	}
}

type Task struct {
	ID        int
	UID       int
	Title     string
	Text      string
	Status    dto.TaskStatus
	Deadline  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t Task) ToDTO() dto.Task {
	return dto.Task{
		ID:        t.ID,
		UID:       t.UID,
		Title:     t.Title,
		Text:      t.Text,
		Status:    t.Status,
		Deadline:  t.Deadline,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

type TaskUpdate struct {
	ID       int
	Title    string
	Text     string
	Status   dto.TaskStatus
	Deadline time.Time
}

func (t TaskUpdate) ToDTO() dto.TaskUpdate {
	return dto.TaskUpdate{
		ID:       t.ID,
		Title:    t.Title,
		Text:     t.Text,
		Status:   t.Status,
		Deadline: t.Deadline,
	}
}

type ShortTask struct {
	ID     int
	Title  string
	Status dto.TaskStatus
}

func (s ShortTask) ToDTO() dto.ShortTask {
	return dto.ShortTask{
		ID:     s.ID,
		Title:  s.Title,
		Status: s.Status,
	}
}
