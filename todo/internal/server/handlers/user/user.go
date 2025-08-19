package user

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/LexusEgorov/todo/internal/models"
	"github.com/LexusEgorov/todo/internal/models/dto"
	"github.com/labstack/echo/v4"
)

const (
	prefix     = "User."
	opGetBody  = prefix + "getBody"
	opRegister = prefix + "Register"
	opAuth     = prefix + "Auth"
	opGet      = prefix + "Get"
	opUpdate   = prefix + "Update"
	opDelete   = prefix + "Delete"
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
	body, err := h.getBody(c)
	if err != nil {
		h.logger.Error(fmt.Errorf("%s: %w", opRegister, err).Error())
		return h.sendBadResponse(c, http.StatusBadRequest, models.ErrGetBody)
	}

	var userData dto.Register
	err = json.Unmarshal(body, &userData)
	if err != nil {
		h.logger.Error(fmt.Errorf("%s: %w", opRegister, err).Error())
		return h.sendBadResponse(c, http.StatusBadRequest, models.ErrReadJSON)
	}

	tokens, err := h.service.Register(userData)
	if err != nil {
		h.logger.Error(fmt.Errorf("%s: %w", opRegister, err).Error())
		return h.sendBadResponse(c, http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, tokens)
}

func (h Handler) Auth(c echo.Context) error {
	body, err := h.getBody(c)
	if err != nil {
		h.logger.Error(fmt.Errorf("%s: %w", opAuth, err).Error())
		return h.sendBadResponse(c, http.StatusBadRequest, models.ErrGetBody)
	}

	var userData dto.Auth
	err = json.Unmarshal(body, &userData)
	if err != nil {
		h.logger.Error(fmt.Errorf("%s: %w", opAuth, err).Error())
		return h.sendBadResponse(c, http.StatusBadRequest, models.ErrReadJSON)
	}

	tokens, err := h.service.Auth(userData)
	if err != nil {
		h.logger.Error(fmt.Errorf("%s: %w", opAuth, err).Error())
		return h.sendBadResponse(c, http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, tokens)
}

func (h Handler) Update(c echo.Context) error {
	body, err := h.getBody(c)
	if err != nil {
		h.logger.Error(fmt.Errorf("%s: %w", opUpdate, err).Error())
		return h.sendBadResponse(c, http.StatusBadRequest, models.ErrGetBody)
	}

	var userData dto.UserUpdate
	err = json.Unmarshal(body, &userData)
	if err != nil {
		h.logger.Error(fmt.Errorf("%s: %w", opUpdate, err).Error())
		return h.sendBadResponse(c, http.StatusBadRequest, models.ErrReadJSON)
	}

	user, err := h.service.Update(userData)
	if err != nil {
		h.logger.Error(fmt.Errorf("%s: %w", opUpdate, err).Error())
		return h.sendBadResponse(c, http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (h Handler) Get(c echo.Context) error {
	uId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error(fmt.Errorf("%s: %w", opGet, err).Error())
		return h.sendBadResponse(c, http.StatusNotFound, err.Error())
	}

	user, err := h.service.Get(uId)
	if err != nil {
		h.logger.Error(fmt.Errorf("%s: %w", opGet, err).Error())
		return h.sendBadResponse(c, http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (h Handler) Delete(c echo.Context) error {
	uId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error(fmt.Errorf("%s: %w", opDelete, err).Error())
		return h.sendBadResponse(c, http.StatusNotFound, err.Error())
	}

	err = h.service.Delete(uId)
	if err != nil {
		h.logger.Error(fmt.Errorf("%s: %w", opDelete, err).Error())
		return h.sendBadResponse(c, http.StatusBadRequest, err.Error())
	}

	return nil
}

func (h Handler) getBody(c echo.Context) ([]byte, error) {
	bodyReader := c.Request().Body
	defer bodyReader.Close()

	body, err := io.ReadAll(bodyReader)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", opGetBody, err)
	}

	return body, nil
}

func (h Handler) sendBadResponse(c echo.Context, code int, message string) error {
	return c.JSON(code, dto.BadResponse{
		Message: message,
	})
}
