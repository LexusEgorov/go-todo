package user

import (
	"context"
	"fmt"

	"github.com/LexusEgorov/todo/internal/models"
	"github.com/LexusEgorov/todo/internal/models/dto"
)

const (
	prefix     = "Services.User."
	opRegister = prefix + "Register"
	opAuth     = prefix + "Auth"
	opRefresh  = prefix + "Refresh"
	opDelete   = prefix + "Delete"
	opGet      = prefix + "Get"
	opUpdate   = prefix + "Update"
)

type UserRepository interface {
	Create(ctx context.Context, user models.User) (int, error)
	Get(ctx context.Context, uId int) (models.User, error)
	Set(ctx context.Context, user models.User) error
	Delete(ctx context.Context, uId int) error
}

type AuthService interface {
	Register(data dto.Register) (dto.Tokens, error)
	Refresh(refresh string) (dto.Tokens, error)
	Auth(data dto.Auth) (dto.Tokens, error)
	Update(data dto.UserUpdate) error
	Access(access string) error
}

type Service struct {
	storage     UserRepository
	authService AuthService
}

func New(storage UserRepository, auth AuthService) *Service {
	return &Service{
		storage:     storage,
		authService: auth,
	}
}

// Register implements user.UserService.
func (s Service) Register(ctx context.Context, data dto.Register) (dto.Tokens, error) {
	if data.Login == "" || data.Name == "" || data.Password == "" || data.TgID == 0 {
		return dto.Tokens{}, models.ErrBadBody
	}

	user := models.User{
		TgID: data.TgID,
		Name: data.Name,
	}

	id, err := s.storage.Create(ctx, user)
	if err != nil {
		return dto.Tokens{}, fmt.Errorf("%s: %w", opRegister, err)
	}

	tokens, err := s.authService.Register(data)
	if err != nil {
		return dto.Tokens{}, fmt.Errorf("%s: %w", opRegister, err)
	}

	tokens.ID = id
	return tokens, nil
}

// Auth implements user.UserService.
func (s Service) Auth(data dto.Auth) (dto.Tokens, error) {
	if data.Login == "" || data.Password == "" {
		return dto.Tokens{}, models.ErrBadBody
	}

	tokens, err := s.authService.Auth(data)
	if err != nil {
		return dto.Tokens{}, fmt.Errorf("%s: %w", opAuth, err)
	}

	return tokens, nil
}

// Refresh implements user.UserService.
func (s Service) Refresh(refresh string) (dto.Tokens, error) {
	if refresh == "" {
		return dto.Tokens{}, models.ErrBadBody
	}

	tokens, err := s.authService.Refresh(refresh)
	if err != nil {
		return dto.Tokens{}, fmt.Errorf("%s: %w", opRefresh, err)
	}

	return tokens, nil
}

// Delete implements user.UserService.
func (s Service) Delete(ctx context.Context, uID int) error {
	if uID == 0 {
		return models.ErrNotFound
	}

	err := s.storage.Delete(ctx, uID)
	if err != nil {
		return fmt.Errorf("%s: %w", opDelete, err)
	}

	return nil
}

// Get implements user.UserService.
func (s Service) Get(ctx context.Context, uID int) (dto.User, error) {
	if uID == 0 {
		return dto.User{}, models.ErrNotFound
	}

	user, err := s.storage.Get(ctx, uID)
	if err != nil {
		return dto.User{}, fmt.Errorf("%s: %w", opGet, err)
	}

	return user.ToDTO(), nil
}

// Update implements user.UserService.
func (s Service) Update(ctx context.Context, user dto.UserUpdate) (dto.User, error) {
	if user.ID == 0 {
		return dto.User{}, models.ErrNotFound
	}

	if user.Login == "" || user.Name == "" || user.Password == "" {
		return dto.User{}, models.ErrBadBody
	}

	update := models.User{
		ID:   user.ID,
		Name: user.Name,
	}

	err := s.storage.Set(ctx, update)
	if err != nil {
		return dto.User{}, fmt.Errorf("%s: %w", opUpdate, err)
	}

	err = s.authService.Update(user)
	if err != nil {
		return dto.User{}, fmt.Errorf("%s: %w", opUpdate, err)
	}

	return update.ToDTO(), nil
}
