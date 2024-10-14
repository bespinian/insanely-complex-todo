package events_test

import (
	"context"
	"testing"

	"github.com/bespinian/ict-todo/backend/tasks/internal/events"
	"github.com/bespinian/ict-todo/backend/tasks/internal/models"
	"github.com/stretchr/testify/assert"
)

type TestQueue struct {
	ReceivedEvents []events.Event
}

func NewTestQueue() TestQueue {
	return TestQueue{[]events.Event{}}
}

func (q *TestQueue) QueueMessage(ctx context.Context, evt events.Event) error {
	q.ReceivedEvents = append(q.ReceivedEvents, evt)
	return nil
}

func TestAddEventToQueue(t *testing.T) {
	q := NewTestQueue()
	events.AddQueue(&q)

	event := events.NewEvent(events.EventTypeTaskCreate, models.Task{})
	events.SendEventToQueues(context.Background(), event)

	assert.Contains(t, q.ReceivedEvents, event, "queues event in test queue")
}

func TestNoQueues(t *testing.T) {
	event := events.NewEvent(events.EventTypeTaskCreate, models.Task{})
	events.SendEventToQueues(context.Background(), event)
}
