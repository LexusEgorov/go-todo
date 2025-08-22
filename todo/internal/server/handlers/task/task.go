package task

import (
	"context"
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
	prefix    = "Handlers.Task."
	opGetBody = prefix + "getBody"
	opGet     = prefix + "Get"
	opGetAll  = prefix + "GetAll"
	opCreate  = prefix + "Create"
	opUpdate  = prefix + "Update"
	opDelete  = prefix + "Delete"
)

type TaskService interface {
	Get(ctx context.Context, taskID int) (dto.Task, error)
	GetAll(ctx context.Context, userID int) ([]dto.ShortTask, error)
	Create(ctx context.Context, task dto.TaskUpdate, uId int) (dto.Task, error)
	Update(ctx context.Context, task dto.TaskUpdate) (dto.TaskUpdate, error)
	Delete(ctx context.Context, taskID int) error
}

type Handler struct {
	logger  *slog.Logger
	service TaskService
}

func New(logger *slog.Logger, taskService TaskService) *Handler {
	return &Handler{
		logger:  logger,
		service: taskService,
	}
}

func (h Handler) Get(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error(fmt.Errorf("%s: %w", opGet, err).Error())
		return h.sendBadResponse(c, http.StatusNotFound, err.Error())
	}

	task, err := h.service.Get(c.Request().Context(), id)
	if err != nil {
		h.logger.Error(fmt.Errorf("%s: %w", opGet, err).Error())
		return h.sendBadResponse(c, http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, task)
}

func (h Handler) GetAll(c echo.Context) error {
	uId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error(fmt.Errorf("%s: %w", opGetAll, err).Error())
		return h.sendBadResponse(c, http.StatusNotFound, err.Error())
	}

	tasks, err := h.service.GetAll(c.Request().Context(), uId)
	if err != nil {
		h.logger.Error(fmt.Errorf("%s: %w", opGetAll, err).Error())
		return h.sendBadResponse(c, http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, tasks)
}

func (h Handler) Create(c echo.Context) error {
	body, err := h.getBody(c)
	if err != nil {
		h.logger.Error(fmt.Errorf("%s: %w", opCreate, err).Error())
		return h.sendBadResponse(c, http.StatusBadRequest, models.ErrGetBody)
	}

	var task dto.TaskUpdate
	err = json.Unmarshal(body, &task)
	if err != nil {
		h.logger.Error(fmt.Errorf("%s: %w", opCreate, err).Error())
		return h.sendBadResponse(c, http.StatusBadRequest, models.ErrReadJSON)
	}

	uId := 1 //TODO: get id from JWT
	created, err := h.service.Create(c.Request().Context(), task, uId)
	if err != nil {
		h.logger.Error(fmt.Errorf("%s: %w", opCreate, err).Error())
		return h.sendBadResponse(c, http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, created)
}

func (h Handler) Update(c echo.Context) error {
	body, err := h.getBody(c)
	if err != nil {
		h.logger.Error(fmt.Errorf("%s: %w", opUpdate, err).Error())
		return h.sendBadResponse(c, http.StatusBadRequest, models.ErrGetBody)
	}

	var task dto.TaskUpdate
	err = json.Unmarshal(body, &task)
	if err != nil {
		h.logger.Error(fmt.Errorf("%s: %w", opUpdate, err).Error())
		return h.sendBadResponse(c, http.StatusBadRequest, models.ErrReadJSON)
	}

	updated, err := h.service.Update(c.Request().Context(), task)
	if err != nil {
		h.logger.Error(fmt.Errorf("%s: %w", opUpdate, err).Error())
		return h.sendBadResponse(c, http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, updated)
}

func (h Handler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error(fmt.Errorf("%s: %w", opDelete, err).Error())
		return h.sendBadResponse(c, http.StatusNotFound, err.Error())
	}

	err = h.service.Delete(c.Request().Context(), id)
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
