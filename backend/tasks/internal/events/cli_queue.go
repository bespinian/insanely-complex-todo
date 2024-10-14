package events

import (
	"context"
	"encoding/json"

	"github.com/gofiber/fiber/v2/log"
)

type CLIQueue struct{}

func (n *CLIQueue) QueueMessage(ctx context.Context, msg Event) error {
	var payload interface{}

	if jsonPayload, err := json.Marshal(msg.Payload); err != nil {
		payload = msg.Payload
	} else {
		payload = string(jsonPayload)
	}
	log.Infof("Notification for %s event. %v", msg.Type, payload)

	return nil
}
