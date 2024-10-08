package internal

import "github.com/bespinian/ict-todo/backend/tasks/internal/models"

type TaskStore interface {
	List() []*models.Task
	Get(string) *models.Task
}
