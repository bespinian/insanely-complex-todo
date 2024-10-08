package internal

import "github.com/bespinian/ict-todo/backend/tasks/internal/models"

type TaskStore interface {
	List() []*models.Task
	Get(string) *models.Task
	Add(*models.Task) (*models.Task, error)
	Update(*models.Task) error
	Delete(string) error
}
