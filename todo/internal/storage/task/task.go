package task

import (
	"context"
	"fmt"

	"github.com/LexusEgorov/todo/internal/models"
	"github.com/LexusEgorov/todo/internal/models/dto"
	"github.com/LexusEgorov/todo/internal/storage/db"
)

const (
	prefix   = "Storage.User."
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
func (s *Storage) Create(task models.Task) error {
	//TODO: get id
	_, err := s.db.DB.Exec(context.TODO(), queryCreate, task.UID, task.Title, task.Text, dto.TaskStatusNew)
	if err != nil {
		return fmt.Errorf("%s: %w", opCreate, err)
	}

	return nil
}

// Delete implements task.TaskRepository.
func (s *Storage) Delete(id int) error {
	_, err := s.db.DB.Exec(context.TODO(), queryDelete, id)
	if err != nil {
		return fmt.Errorf("%s: %w", opDelete, err)
	}

	return nil
}

// Get implements task.TaskRepository.
func (s *Storage) Get(id int) (models.Task, error) {
	task := models.Task{}
	err := s.db.DB.QueryRow(context.TODO(), queryGet, id).Scan(&task.ID, &task.UID, &task.Title, &task.Text, &task.Deadline, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return models.Task{}, fmt.Errorf("%s: %w", opGet, err)
	}

	return task, nil
}

// GetAll implements task.TaskRepository.
func (s *Storage) GetAll(uId int) ([]models.ShortTask, error) {
	tasks := make([]models.ShortTask, 0)
	rows, err := s.db.DB.Query(context.TODO(), queryGetAll, uId)
	if err != nil {
		return []models.ShortTask{}, fmt.Errorf("%s: %w", opGet, err)
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
func (s *Storage) Set(task models.TaskUpdate) error {
	_, err := s.db.DB.Exec(context.TODO(), querySet, task.ID, task.Title, task.Text, task.Status, task.Deadline)
	if err != nil {
		return fmt.Errorf("%s: %w", opSet, err)
	}

	return nil
}
