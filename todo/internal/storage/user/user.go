package user

import (
	"context"
	"fmt"
	"time"

	"github.com/LexusEgorov/todo/internal/models"
	"github.com/LexusEgorov/todo/internal/storage/db"
)

const (
	prefix   = "Storage.User."
	opCreate = prefix + "Create"
	opDelete = prefix + "Delete"
	opGet    = prefix + "Get"
	opSet    = prefix + "Set"
)

type Storage struct {
	db *db.DB
}

func New(db *db.DB) *Storage {
	return &Storage{db: db}
}

// Create implements user.UserRepository.
func (s *Storage) Create(ctx context.Context, user models.User) error {
	_, err := s.db.DB.Exec(ctx, queryCreate, user.TgID, user.Name)
	if err != nil {
		return fmt.Errorf("%s: %w", opCreate, err)
	}

	return nil
}

// Delete implements user.UserRepository.
func (s *Storage) Delete(ctx context.Context, uId int) error {
	_, err := s.db.DB.Exec(ctx, queryDelete, uId)
	if err != nil {
		return fmt.Errorf("%s: %w", opDelete, err)
	}

	return nil
}

// Get implements user.UserRepository.
func (s *Storage) Get(ctx context.Context, uId int) (models.User, error) {
	user := models.User{}
	err := s.db.DB.QueryRow(ctx, queryGet, uId).Scan(&user.ID, &user.TgID, &user.Name, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return models.User{}, fmt.Errorf("%s: %w", opGet, err)
	}

	return user, nil
}

// Set implements user.UserRepository.
func (s *Storage) Set(ctx context.Context, user models.User) error {
	updateDate := time.Now()
	_, err := s.db.DB.Exec(ctx, querySet, user.ID, user.TgID, user.Name, updateDate)
	if err != nil {
		return fmt.Errorf("%s: %w", opSet, err)
	}

	return nil
}
