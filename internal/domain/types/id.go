package types

import "govent/internal/domain/shared"

type EventId struct {
	value string
}

func NewEventId(value string) (*EventId, error) {
	if shared.IsUUIDv4(value) {
		return &EventId{value}, nil
	}

	return nil, shared.ErrEventBadRequest
}

func (b *EventId) Value() string {
	return b.value
}
