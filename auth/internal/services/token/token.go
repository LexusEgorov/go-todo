package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/LexusEgorov/auth/internal/config"
	"github.com/LexusEgorov/auth/internal/models"
)

//TODO: check for err no rows
//TODO: getSub
//TODO: createJWT
//TODO: checkToken

const prefix = "Services.Token."

var (
	ErrEmptyToken = errors.New("token is required")
	ErrEmptyUid   = errors.New("user id is required")

	ErrExpiredToken = errors.New("token is expired")
	ErrBadToken     = errors.New("token isn't valid")
	ErrBlockedToken = errors.New("token is blocked")
)

type TokenRepository interface {
	Add(userId int, token string) error
	Get(token string) error
}

type Service struct {
	storage TokenRepository
	config  *config.AuthConfig
}

func New(storage TokenRepository, config *config.AuthConfig) *Service {
	return &Service{
		storage: storage,
		config:  config,
	}
}

func (s Service) Block(blockData models.BlockDTO) error {
	const op = prefix + "Block"

	err := s.checkToken(blockData.Token)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = s.storage.Add(blockData.UID, blockData.Token)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s Service) Access(token string) error {
	const op = prefix + "Access"

	err := s.checkToken(token)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	//check for ErrNoRows
	err = s.storage.Get(token)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s Service) Refresh(token string) (models.TokensDTO, error) {
	const op = prefix + "Refresh"
	err := s.Access(token)
	if err != nil {
		return models.TokensDTO{}, fmt.Errorf("%s: %w", op, err)
	}

	id, err := s.getSub(token)
	if err != nil {
		return models.TokensDTO{}, fmt.Errorf("%s: %w", op, err)
	}

	tokens, err := CreateTokens(id, s.config.AccessLifetime, s.config.RefreshLifetime)
	if err != nil {
		return models.TokensDTO{}, fmt.Errorf("%s: %w", op, err)
	}
	return tokens.ToDTO(), nil
}

func (s Service) getSub(token string) (int, error) {
	//TODO
	return 0, nil
}

func (s Service) checkToken(token string) error {
	const op = prefix + "checkToken"

	if token == "" {
		return ErrEmptyToken
	}

	// parser := jwt.NewParser()
	// parsedToken, _, err := parser.Parse()
	return nil
}

// TODO: На костыль похоже, возможно, стоит исправить
func CreateTokens(uId int, accessLifetime, refreshLifetime time.Duration) (models.Tokens, error) {
	const op = "createTokens"

	access, err := createJWT(uId, accessLifetime)
	if err != nil {
		return models.Tokens{}, fmt.Errorf("%s: %w", op, err)
	}

	refresh, err := createJWT(uId, refreshLifetime)
	if err != nil {
		return models.Tokens{}, fmt.Errorf("%s: %w", op, err)
	}

	return models.Tokens{
		Access:  access,
		Refresh: refresh,
	}, nil
}

func createJWT(uID int, lifetime time.Duration) (string, error) {
	//TODO: create token
	return "", nil
}
