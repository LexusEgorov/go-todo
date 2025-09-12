package user

import (
	"github.com/LexusEgorov/auth/internal/config"
	"github.com/LexusEgorov/auth/internal/models"
	"github.com/LexusEgorov/auth/internal/services/token"
)

type UserRepository interface {
	Add(data models.Register) (int, error)
	Get(data models.Auth) (int, error)
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
	//TODO: get uid from storage; create Refresh+Access
	//TODO: crypt password
	token.CreateJWT(1, s.config.AccessLifetime)  //access
	token.CreateJWT(1, s.config.RefreshLifetime) //refresh
	return models.TokensDTO{}, nil
}

func (s Service) Register(data models.RegisterDTO) (models.TokensDTO, error) {
	//TODO: validate data, save user to storage, create tokens
	//TODO: crypt password
	token.CreateJWT(1, s.config.AccessLifetime)  //access
	token.CreateJWT(1, s.config.RefreshLifetime) //refresh
	return models.TokensDTO{}, nil
}
