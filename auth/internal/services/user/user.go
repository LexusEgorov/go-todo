package user

import (
	"errors"
	"fmt"

	"github.com/LexusEgorov/auth/internal/config"
	"github.com/LexusEgorov/auth/internal/models"
	"github.com/LexusEgorov/auth/internal/services/token"
	"golang.org/x/crypto/bcrypt"
)

const prefix = "Services.User."

var (
	ErrEmptyLogin    = errors.New("login is required")
	ErrEmptyPassword = errors.New("password is required")
)

type UserRepository interface {
	Add(data models.Register) (int, error)
	Get(data models.Auth) (models.UserPassword, error)
}

type Service struct {
	storage UserRepository
	config  *config.AuthConfig
}

func New(storage UserRepository, config *config.AuthConfig) *Service {
	return &Service{
		storage: storage,
		config:  config,
	}
}

func (s Service) Auth(data models.AuthDTO) (models.TokensDTO, error) {
	const op = prefix + "Auth"

	if data.Login == "" {
		return models.TokensDTO{}, ErrEmptyLogin
	}

	if data.Password == "" {
		return models.TokensDTO{}, ErrEmptyPassword
	}

	user, err := s.storage.Get(models.Auth(data))
	if err != nil {
		return models.TokensDTO{}, fmt.Errorf("%s: %w", op, err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(data.Password))
	if err != nil {
		return models.TokensDTO{}, fmt.Errorf("%s: %w", op, err)
	}

	tokens, err := token.CreateTokens(user.UID, s.config.AccessLifetime, s.config.RefreshLifetime)
	if err != nil {
		return models.TokensDTO{}, fmt.Errorf("%s: %w", op, err)
	}

	return tokens.ToDTO(), nil
}

func (s Service) Register(data models.RegisterDTO) (models.TokensDTO, error) {
	const op = prefix + "Auth"

	if data.Login == "" {
		return models.TokensDTO{}, ErrEmptyLogin
	}

	if data.Password == "" {
		return models.TokensDTO{}, ErrEmptyPassword
	}

	cryptedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.TokensDTO{}, fmt.Errorf("%s: %w", op, err)
	}

	uId, err := s.storage.Add(models.Register{
		TgID:     data.TgID,
		Login:    data.Login,
		Password: string(cryptedPassword),
	})
	if err != nil {
		return models.TokensDTO{}, fmt.Errorf("%s: %w", op, err)
	}

	tokens, err := token.CreateTokens(uId, s.config.AccessLifetime, s.config.RefreshLifetime)
	if err != nil {
		return models.TokensDTO{}, fmt.Errorf("%s: %w", op, err)
	}

	return tokens.ToDTO(), nil
}
