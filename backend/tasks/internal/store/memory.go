package store

import (
	"context"
	"errors"

	"github.com/bespinian/ict-todo/backend/tasks/internal/models"
	"github.com/google/uuid"
)

type MemoryStore struct {
	tasks map[string]*models.Task
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{map[string]*models.Task{}}
}

func (m *MemoryStore) List(ctx context.Context) []*models.Task {
	values := make([]*models.Task, 0, len(m.tasks))
	for _, value := range m.tasks {
		values = append(values, value)
	}
	return values
}

func (m *MemoryStore) Get(ctx context.Context, id string) *models.Task {
	return m.tasks[id]
}

func (m *MemoryStore) Add(ctx context.Context, task *models.Task) (*models.Task, error) {
	if task.Id == "" {
		id := uuid.New()
		task.Id = id.String()
	}
	_, exists := m.tasks[task.Id]
	if exists {
		return nil, errors.New("The task ID already exists!")
	}

	m.tasks[task.Id] = task

	return task, nil
}

func (m *MemoryStore) Update(ctx context.Context, task *models.Task) error {
	m.tasks[task.Id] = task

	return nil
}

func (m *MemoryStore) Delete(ctx context.Context, id string) error {
	delete(m.tasks, id)
	return nil
}
