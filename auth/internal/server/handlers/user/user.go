package user

import (
	"log/slog"

	"github.com/labstack/echo/v4"
)

type UserService interface {
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
	return nil
}

func (s Handler) Register(c echo.Context) error {
	return nil
}

func (s Handler) Access(c echo.Context) error {
	return nil
}

func (s Handler) Refresh(c echo.Context) error {
	return nil
}
