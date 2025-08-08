package task

import (
	"log/slog"

	"github.com/LexusEgorov/todo/internal/models/dto"
	"github.com/labstack/echo/v4"
)

type TaskService interface {
	Get(taskID int) (dto.Task, error)
	GetAll(userID int) ([]dto.Task, error)
	Create(task dto.TaskUpdate) (dto.Task, error)
	Update(task dto.TaskUpdate) (dto.Task, error)
	Delete(taskID int) error
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
	return nil
}

func (h Handler) GetAll(c echo.Context) error {
	return nil
}

func (h Handler) Create(c echo.Context) error {
	return nil
}

func (h Handler) Update(c echo.Context) error {
	return nil
}

func (h Handler) Delete(c echo.Context) error {
	return nil
}
