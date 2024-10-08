package store_test

import (
	"testing"

	"github.com/bespinian/ict-todo/backend/tasks/internal"
	"github.com/bespinian/ict-todo/backend/tasks/internal/models"
	"github.com/bespinian/ict-todo/backend/tasks/internal/store"
	"github.com/stretchr/testify/assert"
)

// make sure MemoryStore implements internal.TaskStore
var _ internal.TaskStore = &store.MemoryStore{}

func getStoreWithATask() (*store.MemoryStore, string) {
	store := store.NewMemoryStore()
	task := &models.Task{Name: "Testtask"}

	_, err := store.Add(task)
	if err != nil {
		return nil, ""
	}
	return store, task.Id
}

func TestMemoryList(t *testing.T) {
	store, id := getStoreWithATask()

	listOfTasks := store.List()
	assert.Len(t, listOfTasks, 1)
	assert.Equal(t, listOfTasks[0].Id, id)
}

func TestMemoryGet(t *testing.T) {
	store, id := getStoreWithATask()

	task := store.Get(id)
	assert.Equal(t, task.Id, id)
}

func TestMemoryAdd(t *testing.T) {
	store := store.NewMemoryStore()
	task := &models.Task{Name: "Testtask"}

	task, err := store.Add(task)

	assert.Nil(t, err)
	assert.NotNil(t, task)
	assert.NotEmpty(t, task.Id, "adds an id")
}

func TestMemoryUpdate(t *testing.T) {
	store, id := getStoreWithATask()
	task := store.Get(id)

	task.Name = "new name"
	err := store.Update(task)

	assert.Nil(t, err)
}

func TestMemoryDelete(t *testing.T) {
	store, id := getStoreWithATask()

	err := store.Delete(id)
	assert.Nil(t, err)

	tasks := store.List()
	assert.Len(t, tasks, 0)
}
