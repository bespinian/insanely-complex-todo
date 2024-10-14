package store_test

import (
	"github.com/bespinian/ict-todo/backend/tasks/internal"
	"github.com/bespinian/ict-todo/backend/tasks/internal/store"
)

var _ internal.TaskStore = &store.MongoStore{}
