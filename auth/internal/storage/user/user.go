package user

import "github.com/LexusEgorov/auth/internal/storage/db"

type Storage struct {
	db *db.DB
}

func New(db *db.DB) *Storage {
	return &Storage{db: db}
}
