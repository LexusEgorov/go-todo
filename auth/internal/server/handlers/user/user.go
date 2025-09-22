package user

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/LexusEgorov/auth/internal/models"
	"github.com/labstack/echo/v4"
)

const prefix = "Handlers.User."

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

func (h Handler) Auth(c echo.Context) error {
	op := prefix + "Auth"
	body, err := h.getBody(c)
	if err != nil {
		h.logger.Error(fmt.Errorf("%s: %w", op, err).Error())
		return c.JSON(echo.ErrUnauthorized.Code, nil)
	}

	authData := models.AuthDTO{}
	err = json.Unmarshal(body, &authData)
	if err != nil {
		h.logger.Error(fmt.Errorf("%s: %w", op, err).Error())
		return c.JSON(echo.ErrUnauthorized.Code, nil)
	}

	tokens, err := h.service.Auth(authData)
	if err != nil {
		h.logger.Error(fmt.Errorf("%s: %w", op, err).Error())
		return c.JSON(echo.ErrUnauthorized.Code, nil)
	}

	return c.JSON(http.StatusOK, tokens)
}

func (h Handler) Register(c echo.Context) error {
	op := prefix + "Register"
	body, err := h.getBody(c)
	if err != nil {
		h.logger.Error(fmt.Errorf("%s: %w", op, err).Error())
		return c.JSON(echo.ErrBadRequest.Code, nil)
	}

	registerData := models.RegisterDTO{}
	err = json.Unmarshal(body, &registerData)
	if err != nil {
		h.logger.Error(fmt.Errorf("%s: %w", op, err).Error())
		return c.JSON(echo.ErrBadRequest.Code, nil)
	}

	tokens, err := h.service.Register(registerData)
	if err != nil {
		h.logger.Error(fmt.Errorf("%s: %w", op, err).Error())
		return c.JSON(echo.ErrInternalServerError.Code, nil)
	}

	return c.JSON(http.StatusCreated, tokens)
}

func (h Handler) getBody(c echo.Context) ([]byte, error) {
	op := prefix + "getBody"
	bodyReader := c.Request().Body
	defer bodyReader.Close()

	body, err := io.ReadAll(bodyReader)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return body, nil
}
