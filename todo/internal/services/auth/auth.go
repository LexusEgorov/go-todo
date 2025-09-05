package auth

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"

	"github.com/LexusEgorov/todo/internal/config"
	"github.com/LexusEgorov/todo/internal/models"
	"github.com/LexusEgorov/todo/internal/models/dto"
)

type Auth struct {
	config *config.AuthConfig
	client *client
}

func New(config *config.AuthConfig) *Auth {
	return &Auth{
		config: config,
		client: newClient(config),
	}
}

//TODO: check for errors for each method

// Auth implements user.AuthService.
func (a *Auth) Auth(data dto.Auth) (dto.Tokens, error) {
	return a.client.AuthRequest(data)
}

// Register implements user.AuthService.
func (a *Auth) Register(data dto.Register) (dto.Tokens, error) {
	return a.client.RegisterRequest(data)
}

// Update implements user.AuthService.
func (a *Auth) Update(data dto.UserUpdate) error {
	//TODO: add logic
	return nil
}

func (a *Auth) Access(access string) error {
	return a.client.AccessRequest(access)
}

func (a *Auth) Refresh(refresh string) (dto.Tokens, error) {
	return a.client.RefreshRequest(refresh)
}

func (a *Auth) GetId(jwtToken string) (string, error) {
	if jwtToken == "" {
		return "", models.ErrEmptyToken
	}

	token, err := jwt.Parse(jwtToken, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return a.config.Secret, nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", models.ErrInvalidToken
	}

	return token.Claims.GetSubject()
}
