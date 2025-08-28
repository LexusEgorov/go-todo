package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"

	"github.com/LexusEgorov/todo/internal/config"
	"github.com/LexusEgorov/todo/internal/models"
	"github.com/LexusEgorov/todo/internal/models/dto"
)

const (
	prefix     = "Services.Auth."
	opRegister = prefix + "RegisterRequest"
	opAuth     = prefix + "AuthRequest"
	opAccess   = prefix + "AccessRequest"
	opRefresh  = prefix + "RefreshRequest"

	routeAuth     = "/auth"
	routeRegister = "/register"
	routeAccess   = "/access"
	routeRefresh  = "/refresh"
)

type client struct {
	config *config.AuthConfig
	client *resty.Client
}

func newClient(cfg *config.AuthConfig) *client {
	restyClient := resty.New()
	restyClient.RetryCount = cfg.RetryCount

	return &client{
		config: cfg,
		client: restyClient,
	}
}

func (c client) AuthRequest(data dto.Auth) (tokens dto.Tokens, err error) {
	req := c.client.R()

	body, err := json.Marshal(data)
	if err != nil {
		return dto.Tokens{}, fmt.Errorf("%s: %w", opAuth, err)
	}

	req.Body = body
	response, err := req.Execute(http.MethodPost, c.config.Addr+routeAuth)
	if err != nil {
		return dto.Tokens{}, fmt.Errorf("%s: %w", opAuth, err)
	}

	if response.StatusCode() == http.StatusUnauthorized {
		return dto.Tokens{}, models.ErrUnauthorized
	}

	err = json.Unmarshal(response.Body(), &tokens)
	if err != nil {
		return dto.Tokens{}, fmt.Errorf("%s: %w", opAuth, err)
	}

	return
}

func (c client) RegisterRequest(data dto.Register) (tokens dto.Tokens, err error) {
	req := c.client.R()

	body, err := json.Marshal(data)
	if err != nil {
		return dto.Tokens{}, fmt.Errorf("%s: %w", opRegister, err)
	}

	req.Body = body
	response, err := req.Execute(http.MethodPost, c.config.Addr+routeRegister)
	if err != nil {
		return dto.Tokens{}, fmt.Errorf("%s: %w", opRegister, err)
	}

	if response.StatusCode() == http.StatusBadRequest {
		//TODO: another error
		return dto.Tokens{}, models.ErrBadBody
	}

	err = json.Unmarshal(response.Body(), &tokens)
	if err != nil {
		return dto.Tokens{}, fmt.Errorf("%s: %w", opRegister, err)
	}

	return
}

func (c client) AccessRequest(access string) error {
	req := c.client.R()
	req.Header.Add("Authorization", access)

	response, err := req.Execute(http.MethodGet, c.config.Addr+routeAccess)
	if err != nil {
		return fmt.Errorf("%s: %w", opAccess, err)
	}

	if response.StatusCode() == http.StatusUnauthorized {
		return models.ErrUnauthorized
	}

	return nil
}

func (c client) RefreshRequest(refresh string) (tokens dto.Tokens, err error) {
	req := c.client.R()
	req.Header.Add("Authorization", refresh)

	response, err := req.Execute(http.MethodGet, c.config.Addr+routeRefresh)
	if err != nil {
		return dto.Tokens{}, fmt.Errorf("%s: %w", opRefresh, err)
	}

	if response.StatusCode() == http.StatusUnauthorized {
		return dto.Tokens{}, models.ErrUnauthorized
	}

	err = json.Unmarshal(response.Body(), &tokens)
	if err != nil {
		return dto.Tokens{}, fmt.Errorf("%s: %w", opRegister, err)
	}

	return
}

//TODO: updateRequest
