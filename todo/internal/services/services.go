package services

import (
	"fmt"

	"github.com/LexusEgorov/todo/internal/config"
	"github.com/LexusEgorov/todo/internal/services/task"
	"github.com/LexusEgorov/todo/internal/services/user"
	"github.com/LexusEgorov/todo/internal/storage/db"
	taskRepo "github.com/LexusEgorov/todo/internal/storage/task"
	userRepo "github.com/LexusEgorov/todo/internal/storage/user"
)

const (
	opNew = "Services.New"
)

type Services struct {
	User user.Service
	Task task.Service
}

func New(cfg config.DBConfig) (*Services, error) {
	db, err := db.NewDB(cfg)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", opNew, err)
	}

	return &Services{
		Task: *task.New(taskRepo.New(db)),
		User: *user.New(userRepo.New(db), nil),
	}, nil
}
