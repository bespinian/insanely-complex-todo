package store_test

import (
	"context"
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
	task := models.Task{Name: "Testtask"}

	_, err := store.Add(context.Background(), task)
	if err != nil {
		return nil, ""
	}
	return store, task.Id
}

func TestMemoryList(t *testing.T) {
	store, id := getStoreWithATask()

	listOfTasks, _ := store.List(context.Background())
	assert.Len(t, listOfTasks, 1)
	assert.Equal(t, listOfTasks[0].Id, id)
}

func TestMemoryListEmpty(t *testing.T) {
	store := store.MemoryStore{}

	listOfTasks, _ := store.List(context.Background())
	assert.NotNil(t, listOfTasks)
	assert.Len(t, listOfTasks, 0)
}

func TestMemoryGet(t *testing.T) {
	store, id := getStoreWithATask()

	task, err := store.Get(context.Background(), id)
	assert.Nil(t, err)
	assert.Equal(t, task.Id, id)
}

func TestMemoryAdd(t *testing.T) {
	store := store.NewMemoryStore()
	task := models.Task{Name: "Testtask"}

	task, err := store.Add(context.Background(), task)

	assert.Nil(t, err)
	assert.NotNil(t, task)
	assert.NotEmpty(t, task.Id, "adds an id")
}

func TestMemoryUpdate(t *testing.T) {
	store, id := getStoreWithATask()
	task, _ := store.Get(context.Background(), id)

	task.Name = "new name"
	_, err := store.Update(context.Background(), task)

	assert.Nil(t, err)
}

func TestMemoryDelete(t *testing.T) {
	store, id := getStoreWithATask()

	err := store.Delete(context.Background(), id)
	assert.Nil(t, err)

	tasks, _ := store.List(context.Background())
	assert.Len(t, tasks, 0)
}
