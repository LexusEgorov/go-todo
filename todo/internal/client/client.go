package client

import "github.com/LexusEgorov/todo/internal/models/dto"

type Client struct{}

//TODO: Add logic

// Auth implements user.AuthService.
func (c *Client) Auth(data dto.Auth) (dto.Tokens, error) {
	return dto.Tokens{
		Access:  "access",
		Refresh: "refresh",
	}, nil
}

// Register implements user.AuthService.
func (c *Client) Register(data dto.Register) (dto.Tokens, error) {
	return dto.Tokens{
		Access:  "access",
		Refresh: "refresh",
	}, nil
}

// Update implements user.AuthService.
func (c *Client) Update(data dto.UserUpdate) error {
	return nil
}

func New() *Client {
	return &Client{}
}
