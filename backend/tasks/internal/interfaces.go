package internal

import (
	"context"

	"github.com/bespinian/ict-todo/backend/tasks/internal/models"
)

type TaskStore interface {
	List(context.Context) []*models.Task
	Get(context.Context, string) *models.Task
	Add(context.Context, *models.Task) (*models.Task, error)
	Update(context.Context, *models.Task) error
	Delete(context.Context, string) error
}
