package user

import (
	"log/slog"

	"github.com/LexusEgorov/auth/internal/models"
	"github.com/labstack/echo/v4"
)

type UserService interface {
	Auth(data models.AuthDTO) (models.TokensDTO, error)
	Register(data models.RegisterDTO) (models.TokensDTO, error)
}

type Handler struct {
	logger  *slog.Logger
	service UserService
}

func New(logger *slog.Logger, service UserService) *Handler {
	return &Handler{
		logger:  logger,
		service: service,
	}
}

func (s Handler) Auth(c echo.Context) error {
	//TODO getBody, checkErr
	s.service.Auth(models.AuthDTO{})

	return nil
}

func (s Handler) Register(c echo.Context) error {
	//TODO getBody, checkErr
	s.service.Register(models.RegisterDTO{})
	return nil
}
