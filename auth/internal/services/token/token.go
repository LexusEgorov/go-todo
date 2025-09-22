package token

import (
	"time"

	"github.com/LexusEgorov/auth/internal/models"
)

type TokenRepository interface {
	Add(userId int, token string) error
	Get(token string) error
}

type Service struct {
	storage TokenRepository
}

func New(storage TokenRepository) *Service {
	return &Service{
		storage: storage,
	}
}

func (s Service) Block(blockData models.BlockDTO) error {
	//TODO: validate uid, token
	return nil
}

func (s Service) Access(token string) error {
	//TODO: just check for exp and sign
	//TODO: check blackLIst
	return nil
}

func (s Service) Refresh(token string) (models.TokensDTO, error) {
	//TODO: check token. If ok, create new pair of tokens
	return models.TokensDTO{}, nil
}

func CreateJWT(uID int, lifetime time.Duration) (string, error) {
	//TODO: create token
	return "", nil
}
