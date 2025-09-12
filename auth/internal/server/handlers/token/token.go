package token

import (
	"log/slog"

	"github.com/labstack/echo/v4"
)

type TokenService interface {
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
	return nil
}
