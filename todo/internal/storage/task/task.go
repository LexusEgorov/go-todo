package task

import (
	"context"
	"fmt"

	"github.com/LexusEgorov/todo/internal/models"
	"github.com/LexusEgorov/todo/internal/models/dto"
	"github.com/LexusEgorov/todo/internal/storage/db"
)

const (
	prefix   = "Storage.Task."
	opCreate = prefix + "Create"
	opDelete = prefix + "Delete"
	opGet    = prefix + "Get"
	opGetAll = prefix + "GetAll"
	opSet    = prefix + "Set"
)

type Storage struct {
	db *db.DB
}

func New(db *db.DB) *Storage {
	return &Storage{db: db}
}

// Create implements task.TaskRepository.
func (s *Storage) Create(ctx context.Context, task models.Task) (id int, err error) {
	err = s.db.DB.QueryRow(ctx, queryCreate, task.UID, task.Title, task.Text, dto.TaskStatusNew, task.Deadline).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", opCreate, err)
	}

	return
}

// Delete implements task.TaskRepository.
func (s *Storage) Delete(ctx context.Context, id int) error {
	_, err := s.db.DB.Exec(ctx, queryDelete, id)
	if err != nil {
		return fmt.Errorf("%s: %w", opDelete, err)
	}

	return nil
}

// Get implements task.TaskRepository.
func (s *Storage) Get(ctx context.Context, id int) (models.Task, error) {
	task := models.Task{}
	err := s.db.DB.QueryRow(ctx, queryGet, id).Scan(&task.ID, &task.UID, &task.Title, &task.Text, &task.Deadline, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return models.Task{}, fmt.Errorf("%s: %w", opGet, err)
	}

	return task, nil
}

// GetAll implements task.TaskRepository.
func (s *Storage) GetAll(ctx context.Context, uId int) ([]models.ShortTask, error) {
	tasks := make([]models.ShortTask, 0)
	rows, err := s.db.DB.Query(ctx, queryGetAll, uId)
	if err != nil {
		return []models.ShortTask{}, fmt.Errorf("%s: %w", opGetAll, err)
	}

	defer rows.Close()
	for rows.Next() {
		task := models.ShortTask{}
		rows.Scan(&task.ID, &task.Title, &task.Status)
		tasks = append(tasks, task)
	}

	return tasks, nil
}

// Set implements task.TaskRepository.
func (s *Storage) Set(ctx context.Context, task models.TaskUpdate) error {
	_, err := s.db.DB.Exec(ctx, querySet, task.ID, task.Title, task.Text, task.Status, task.Deadline)
	if err != nil {
		return fmt.Errorf("%s: %w", opSet, err)
	}

	return nil
}
