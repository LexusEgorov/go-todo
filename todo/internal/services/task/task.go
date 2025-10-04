package task

import (
	"context"
	"fmt"

	"github.com/LexusEgorov/todo/internal/models"
	"github.com/LexusEgorov/todo/internal/models/dto"
)

const (
	prefix   = "Services.Task."
	opCreate = prefix + "Create"
	opDelete = prefix + "Delete"
	opGet    = prefix + "Get"
	opGetAll = prefix + "GetAll"
	opUpdate = prefix + "Update"
)

type TaskRepository interface {
	Get(ctx context.Context, id int) (models.Task, error)
	GetAll(ctx context.Context, uId int) ([]models.ShortTask, error)
	Create(ctx context.Context, task models.Task) (int, error)
	Set(ctx context.Context, task models.TaskUpdate) error
	Delete(ctx context.Context, id int) error
}

type Service struct {
	storage TaskRepository
}

func New(storage TaskRepository) *Service {
	return &Service{
		storage: storage,
	}
}

// Create implements task.TaskService.
func (s Service) Create(ctx context.Context, task dto.TaskUpdate, uId int) (dto.Task, error) {
	if task.Title == "" || task.Deadline.IsZero() {
		return dto.Task{}, models.ErrBadBody
	}

	coreTask := models.Task{
		UID:      uId,
		Title:    task.Title,
		Text:     task.Text,
		Status:   dto.TaskStatusNew,
		Deadline: task.Deadline,
	}

	id, err := s.storage.Create(ctx, coreTask)
	if err != nil {
		return dto.Task{}, fmt.Errorf("%s: %w", opCreate, err)
	}

	coreTask.ID = id
	return coreTask.ToDTO(), nil
}

// Delete implements task.TaskService.
func (s Service) Delete(ctx context.Context, taskID int) error {
	if taskID == 0 {
		return models.ErrNotFound
	}

	err := s.storage.Delete(ctx, taskID)
	if err != nil {
		return fmt.Errorf("%s: %w", opDelete, err)
	}

	return nil
}

// Get implements task.TaskService.
func (s Service) Get(ctx context.Context, taskID int) (dto.Task, error) {
	if taskID == 0 {
		return dto.Task{}, models.ErrNotFound
	}

	task, err := s.storage.Get(ctx, taskID)
	if err != nil {
		return dto.Task{}, fmt.Errorf("%s: %w", opGet, err)
	}

	return task.ToDTO(), nil
}

// GetAll implements task.TaskService.
func (s Service) GetAll(ctx context.Context, userID int) ([]dto.ShortTask, error) {
	if userID == 0 {
		return []dto.ShortTask{}, models.ErrNotFound
	}

	tasks, err := s.storage.GetAll(ctx, userID)
	if err != nil {
		return []dto.ShortTask{}, fmt.Errorf("%s: %w", opGetAll, err)
	}

	dtoTasks := make([]dto.ShortTask, len(tasks))
	for i, task := range tasks {
		dtoTasks[i] = task.ToDTO()
	}

	return dtoTasks, nil
}

// Update implements task.TaskService.
func (s Service) Update(ctx context.Context, task dto.TaskUpdate) (dto.TaskUpdate, error) {
	if task.ID == 0 {
		return dto.TaskUpdate{}, models.ErrNotFound
	}

	if task.Deadline.IsZero() || task.Title == "" || !validateStatus(task.Status) {
		return dto.TaskUpdate{}, models.ErrBadBody
	}

	coreTask := models.TaskUpdate{
		ID:       task.ID,
		Title:    task.Title,
		Text:     task.Text,
		Status:   task.Status,
		Deadline: task.Deadline,
	}

	err := s.storage.Set(ctx, coreTask)
	if err != nil {
		return dto.TaskUpdate{}, fmt.Errorf("%s: %w", opUpdate, err)
	}

	return coreTask.ToDTO(), nil
}

func validateStatus(status dto.TaskStatus) bool {
	switch status {
	case dto.TaskStatusCancelled:
		fallthrough
	case dto.TaskStatusFinished:
		fallthrough
	case dto.TaskStatusInProgress:
		fallthrough
	case dto.TaskStatusNew:
		return true
	}

	return false
}
