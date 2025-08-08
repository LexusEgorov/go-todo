package services

import (
	"github.com/LexusEgorov/todo/internal/services/task"
	"github.com/LexusEgorov/todo/internal/services/user"
)

type Services struct {
	User user.Service
	Task task.Service
}

func New() *Services {
	return &Services{}
}
