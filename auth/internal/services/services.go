package services

import (
	"fmt"

	"github.com/LexusEgorov/auth/internal/config"
	"github.com/LexusEgorov/auth/internal/services/token"
	"github.com/LexusEgorov/auth/internal/services/user"
	"github.com/LexusEgorov/auth/internal/storage/db"
	tokenRepo "github.com/LexusEgorov/auth/internal/storage/token"
	userRepo "github.com/LexusEgorov/auth/internal/storage/user"
)

const (
	opNew = "Services.New"
)

type Services struct {
	User  user.Service
	Token token.Service
}

func New(cfg config.Config) (*Services, error) {
	db, err := db.NewDB(cfg.DB)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", opNew, err)
	}

	return &Services{
		Token: *token.New(tokenRepo.New(db)),
		User:  *user.New(userRepo.New(db), &cfg.Auth),
	}, nil
}
