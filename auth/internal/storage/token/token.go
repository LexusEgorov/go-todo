package token

import "github.com/LexusEgorov/auth/internal/storage/db"

type Storage struct {
	db *db.DB
}

func New(db *db.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) Add(userId int, token string) error {
	//TODO: add to db
	return nil
}

func (s *Storage) Get(token string) error {
	//TODO: get from db
	return nil
}
