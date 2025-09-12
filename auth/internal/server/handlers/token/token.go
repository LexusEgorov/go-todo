package token

import (
	"log/slog"

	"github.com/LexusEgorov/auth/internal/models"
	"github.com/labstack/echo/v4"
)

type TokenService interface {
	Block(userId int, token string) error
	Access(token string) error
	Refresh(token string) (models.TokensDTO, error)
}

type Handler struct {
	logger  *slog.Logger
	service TokenService
}

func New(logger *slog.Logger, service TokenService) *Handler {
	return &Handler{
		logger:  logger,
		service: service,
	}
}

func (s Handler) Block(c echo.Context) error {
	//TODO: get id + token
	return s.service.Block(1, "")
}

func (s Handler) Access(c echo.Context) error {
	//TODO: get Authorization, checkErr
	s.service.Access("")
	return nil
}

func (s Handler) Refresh(c echo.Context) error {
	//TODO: get Authorization, checkErr
	s.service.Refresh("")
	return nil
}
