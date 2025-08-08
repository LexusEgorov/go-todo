package task

import "github.com/LexusEgorov/todo/internal/models/dto"

type Service struct{}

// Create implements task.TaskService.
func (s Service) Create(task dto.TaskUpdate) (dto.Task, error) {
	panic("unimplemented")
}

// Delete implements task.TaskService.
func (s Service) Delete(taskID int) error {
	panic("unimplemented")
}

// Get implements task.TaskService.
func (s Service) Get(taskID int) (dto.Task, error) {
	panic("unimplemented")
}

// GetAll implements task.TaskService.
func (s Service) GetAll(userID int) ([]dto.Task, error) {
	panic("unimplemented")
}

// Update implements task.TaskService.
func (s Service) Update(task dto.TaskUpdate) (dto.Task, error) {
	panic("unimplemented")
}

func New() *Service {
	return &Service{}
}
