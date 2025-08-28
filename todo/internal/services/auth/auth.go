package auth

import (
	"github.com/LexusEgorov/todo/internal/config"
	"github.com/LexusEgorov/todo/internal/models/dto"
)

type Client struct {
	config *config.AuthConfig
	client *client
}

func New(config *config.AuthConfig) *Client {
	return &Client{
		config: config,
		client: newClient(config),
	}
}

//TODO: check for errors for each method

// Auth implements user.AuthService.
func (c *Client) Auth(data dto.Auth) (dto.Tokens, error) {
	return c.client.AuthRequest(data)
}

// Register implements user.AuthService.
func (c *Client) Register(data dto.Register) (dto.Tokens, error) {
	return c.client.RegisterRequest(data)
}

// Update implements user.AuthService.
func (c *Client) Update(data dto.UserUpdate) error {
	//TODO: add logic
	return nil
}

func (c *Client) Access(access string) error {
	return c.client.AccessRequest(access)
}

func (c *Client) Refresh(refresh string) (dto.Tokens, error) {
	return c.client.RefreshRequest(refresh)
}

//TODO: add getting userData from JWT
