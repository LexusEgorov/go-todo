package token

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/LexusEgorov/auth/internal/models"
	"github.com/labstack/echo/v4"
)

const prefix = "Handlers.Token."

type TokenService interface {
	Block(blockData models.BlockDTO) error
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

func (h Handler) Block(c echo.Context) error {
	op := prefix + "Block"
	body, err := h.getBody(c)
	if err != nil {
		h.logger.Error(fmt.Errorf("%s: %w", op, err).Error())
		return c.JSON(echo.ErrBadRequest.Code, err)
	}

	blockData := models.BlockDTO{}
	err = json.Unmarshal(body, &blockData)
	if err != nil {
		h.logger.Error(fmt.Errorf("%s: %w", op, err).Error())
		return c.JSON(echo.ErrBadRequest.Code, err)
	}

	err = h.service.Block(blockData)
	if err != nil {
		h.logger.Error(fmt.Errorf("%s: %w", op, err).Error())
		return c.JSON(echo.ErrInternalServerError.Code, err)
	}

	return c.JSON(http.StatusOK, nil)
}

func (h Handler) Access(c echo.Context) error {
	op := prefix + "Access"
	access := c.Request().Header.Get(echo.HeaderAuthorization)
	if access == "" {
		return c.JSON(echo.ErrUnauthorized.Code, nil)
	}

	err := h.service.Access(access)
	if err != nil {
		h.logger.Error(fmt.Errorf("%s: %w", op, err).Error())
		return c.JSON(echo.ErrUnauthorized.Code, nil)
	}

	return c.JSON(http.StatusOK, nil)
}

func (h Handler) Refresh(c echo.Context) error {
	op := prefix + "Refresh"
	refresh := c.Request().Header.Get(echo.HeaderAuthorization)
	if refresh == "" {
		return c.JSON(echo.ErrUnauthorized.Code, nil)
	}

	tokens, err := h.service.Refresh(refresh)
	if err != nil {
		h.logger.Error(fmt.Errorf("%s: %w", op, err).Error())
		return c.JSON(echo.ErrUnauthorized.Code, nil)
	}

	return c.JSON(http.StatusOK, tokens)
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
