package user

import (
	"fmt"

	"github.com/LexusEgorov/todo/internal/models"
	"github.com/LexusEgorov/todo/internal/models/dto"
)

const (
	prefix     = "Services.User."
	opRegister = prefix + "Register"
	opAuth     = prefix + "Auth"
	opDelete   = prefix + "Delete"
	opGet      = prefix + "Get"
	opUpdate   = prefix + "Update"
)

type UserRepository interface {
	Create(user models.User) error
	Get(uId int) (models.User, error)
	Set(user models.User) error
	Delete(uId int) error
}

type AuthService interface {
	Register(data dto.Register) (dto.Tokens, error)
	Auth(data dto.Auth) (dto.Tokens, error)
	Update(data dto.UserUpdate) error
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
func (s Service) Register(data dto.Register) (dto.Tokens, error) {
	if data.Login == "" || data.Name == "" || data.Password == "" || data.TgID == 0 {
		return dto.Tokens{}, models.ErrBadBody
	}

	user := models.User{
		TgID: data.TgID,
		Name: data.Name,
	}

	err := s.storage.Create(user)
	if err != nil {
		return dto.Tokens{}, fmt.Errorf("%s: %w", opRegister, err)
	}

	tokens, err := s.authService.Register(data)
	if err != nil {
		return dto.Tokens{}, fmt.Errorf("%s: %w", opRegister, err)
	}

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

// Delete implements user.UserService.
func (s Service) Delete(uID int) error {
	if uID == 0 {
		return models.ErrNotFound
	}

	err := s.storage.Delete(uID)
	if err != nil {
		return fmt.Errorf("%s: %w", opDelete, err)
	}

	return nil
}

// Get implements user.UserService.
func (s Service) Get(uID int) (dto.User, error) {
	if uID == 0 {
		return dto.User{}, models.ErrNotFound
	}

	user, err := s.storage.Get(uID)
	if err != nil {
		return dto.User{}, fmt.Errorf("%s: %w", opGet, err)
	}

	return user.ToDTO(), nil
}

// Update implements user.UserService.
func (s Service) Update(user dto.UserUpdate) (dto.User, error) {
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

	err := s.storage.Set(update)
	if err != nil {
		return dto.User{}, fmt.Errorf("%s: %w", opUpdate, err)
	}

	err = s.authService.Update(user)
	if err != nil {
		return dto.User{}, fmt.Errorf("%s: %w", opUpdate, err)
	}

	return update.ToDTO(), nil
}
