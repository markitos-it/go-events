package internal_test

import (
	"go-vents/internal/domain/shared"
	"go-vents/internal/domain/types"
)

func NewRandomEvent() *types.Event {
	event, _ := types.NewEvent(
		shared.UUIDv4(),
		shared.RandomSlug(),
		shared.RandomSlug(),
		"",
	)

	return event
}

func NewRandomOnlyNameEvent() *types.Event {
	event, _ := types.NewEvent(shared.UUIDv4(), shared.RandomSlug(), "", "")

	return event
}
func NewRandomEventWithSlugAndSource(slug, source string) *types.Event {
	event, _ := types.NewEvent(
		shared.UUIDv4(),
		slug,
		source,
		"",
	)

	return event
}

func NewRandomEventWithCustomId(eventId *types.Id) *types.Event {
	event, _ := types.NewEvent(
		eventId.Value(),
		shared.RandomSlug(),
		shared.RandomSlug(),
		"",
	)
	return event
}
