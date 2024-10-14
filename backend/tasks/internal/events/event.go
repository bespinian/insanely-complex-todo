package events

import (
	"encoding/json"
	"time"
)

type EventType string

const (
	EventTypeTaskCreate EventType = "create_task"
	EventTypeTaskDelete EventType = "delete_task"
	EventTypeTaskUpdate EventType = "update_task"
)

type Event struct {
	Timestamp time.Time   `json:"timestamp"`
	Type      EventType   `json:"type"`
	Payload   interface{} `json:"payload"`
}

func NewEvent(eventType EventType, payload interface{}) Event {
	return Event{Timestamp: time.Now(), Type: eventType, Payload: payload}
}

func (m *Event) Serialize() ([]byte, error) {
	return json.Marshal(m)
}
