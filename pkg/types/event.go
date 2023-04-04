package types

import (
	"encoding/json"
	"github.com/google/uuid"
)

type (
	Event struct {
		ID            uuid.UUID
		Topic         string
		CorrelationID string
		Message       []byte
	}
)

func FromJSON(topic string, correlationID string, message any) *Event {
	b, err := json.Marshal(message)
	if err != nil {
		return nil
	}

	return NewEvent(topic, correlationID, b)
}

func NewEvent(topic string, correlationID string, message []byte) *Event {
	return &Event{
		ID:            uuid.New(),
		Topic:         topic,
		CorrelationID: correlationID,
		Message:       message,
	}
}
