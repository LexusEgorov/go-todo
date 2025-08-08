package user

import (
	"log/slog"

	"github.com/LexusEgorov/todo/internal/models/dto"
	"github.com/labstack/echo/v4"
)

type UserService interface {
	Register(data dto.Register) (dto.Tokens, error)
	Auth(data dto.Auth) (dto.Tokens, error)
	Get(uID int) (dto.User, error)
	Update(user dto.UserUpdate) (dto.User, error)
	Delete(uID int) error
}

type Handler struct {
	logger  *slog.Logger
	service UserService
}

func New(logger *slog.Logger, userService UserService) *Handler {
	return &Handler{
		logger:  logger,
		service: userService,
	}
}

func (h Handler) Register(c echo.Context) error {
	return nil
}

func (h Handler) Auth(c echo.Context) error {
	return nil
}

func (h Handler) Get(c echo.Context) error {
	return nil
}

func (h Handler) Update(c echo.Context) error {
	return nil
}

func (h Handler) Delete(c echo.Context) error {
	return nil
}
