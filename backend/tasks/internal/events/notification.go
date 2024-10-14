package events

import (
	"context"

	"github.com/bespinian/ict-todo/backend/tasks/internal/models"
)

func NotifyTaskCreated(ctx context.Context, task models.Task) {
	msg := NewEvent(EventTypeTaskCreate, task)
	SendEventToQueues(ctx, msg)
}

func NotifyTaskUpdated(ctx context.Context, task models.Task) {
	msg := NewEvent(EventTypeTaskUpdate, task)
	SendEventToQueues(ctx, msg)
}

func NotifyTaskDeleted(ctx context.Context, task models.Task) {
	msg := NewEvent(EventTypeTaskDelete, task)
	SendEventToQueues(ctx, msg)
}
