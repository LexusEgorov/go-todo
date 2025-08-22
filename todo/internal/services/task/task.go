package task

import (
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
	Get(id int) (models.Task, error)
	GetAll(uId int) ([]models.ShortTask, error)
	Create(task models.Task) error
	Set(task models.TaskUpdate) error
	Delete(id int) error
}

type Service struct {
	storage TaskRepository
}

func New() *Service {
	return &Service{}
}

// Create implements task.TaskService.
func (s Service) Create(task dto.TaskUpdate, uId int) (dto.Task, error) {
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

	err := s.storage.Create(coreTask)
	if err != nil {
		return dto.Task{}, fmt.Errorf("%s: %w", opCreate, err)
	}

	return coreTask.ToDTO(), nil
}

// Delete implements task.TaskService.
func (s Service) Delete(taskID int) error {
	if taskID == 0 {
		return models.ErrNotFound
	}

	err := s.storage.Delete(taskID)
	if err != nil {
		return fmt.Errorf("%s: %w", opDelete, err)
	}

	return nil
}

// Get implements task.TaskService.
func (s Service) Get(taskID int) (dto.Task, error) {
	if taskID == 0 {
		return dto.Task{}, models.ErrNotFound
	}

	task, err := s.storage.Get(taskID)
	if err != nil {
		return dto.Task{}, fmt.Errorf("%s: %w", opGet, err)
	}

	return task.ToDTO(), nil
}

// GetAll implements task.TaskService.
func (s Service) GetAll(userID int) ([]dto.ShortTask, error) {
	if userID == 0 {
		return []dto.ShortTask{}, models.ErrNotFound
	}

	tasks, err := s.storage.GetAll(userID)
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
func (s Service) Update(task dto.TaskUpdate) (dto.TaskUpdate, error) {
	if task.ID == 0 {
		return dto.TaskUpdate{}, models.ErrNotFound
	}

	if task.Deadline.IsZero() || task.Title == "" || !validStatus(task.Status) {
		return dto.TaskUpdate{}, models.ErrBadBody
	}

	coreTask := models.TaskUpdate{
		ID:       task.ID,
		Title:    task.Title,
		Text:     task.Text,
		Status:   task.Status,
		Deadline: task.Deadline,
	}

	err := s.storage.Set(coreTask)
	if err != nil {
		return dto.TaskUpdate{}, fmt.Errorf("%s: %w", opUpdate, err)
	}

	return coreTask.ToDTO(), nil
}

func validStatus(status dto.TaskStatus) bool {
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
