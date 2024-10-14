package events

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
)

type EventQueue interface {
	QueueMessage(context.Context, Event) error
}

var queues = []EventQueue{}

func AddQueue(q EventQueue) {
	queues = append(queues, q)
}

func SendEventToQueues(ctx context.Context, msg Event) {
	if len(queues) == 0 {
		log.Warn("No event queues configured. Not sending event.")
	}

	for _, n := range queues {
		if err := n.QueueMessage(ctx, msg); err != nil {
			log.Warnf("notification not sent: %v", err)
		}
	}
}
