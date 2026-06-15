package types

import "go-vents/internal/domain/shared"

type Id struct {
	value string
}

func NewId(value string) (*Id, error) {
	if shared.IsUUIDv4(value) {
		return &Id{value}, nil
	}

	return nil, shared.ErrEventBadRequest
}

func (b *Id) Value() string {
	return b.value
}
