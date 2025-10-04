package user

import (
	"github.com/LexusEgorov/auth/internal/models"
	"github.com/LexusEgorov/auth/internal/storage/db"
)

type Storage struct {
	db *db.DB
}

func New(db *db.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) Add(data models.Register) (int, error) {
	//TODO: just add to db
	return 0, nil
}

func (s *Storage) Get(data models.Auth) (models.UserPassword, error) {
	//TODO: just get from db
	return models.UserPassword{}, nil
}
