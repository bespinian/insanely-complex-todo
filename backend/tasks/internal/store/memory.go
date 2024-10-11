package store

import (
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

func (m *MemoryStore) List() []*models.Task {
	values := make([]*models.Task, 0, len(m.tasks))
	for _, value := range m.tasks {
		values = append(values, value)
	}
	return values
}

func (m *MemoryStore) Get(id string) *models.Task {
	return m.tasks[id]
}

func (m *MemoryStore) Add(task *models.Task) (*models.Task, error) {
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

func (m *MemoryStore) Update(task *models.Task) error {
	m.tasks[task.Id] = task

	return nil
}

func (m *MemoryStore) Delete(id string) error {
	delete(m.tasks, id)
	return nil
}
