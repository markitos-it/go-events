package types

import (
	"go-vents/internal/domain/shared"
)

type Slug struct {
	value string
}

func NewSlug(value string) (*Slug, error) {

	if isValidSlug(value) {
		return &Slug{value}, nil
	}

	return nil, shared.ErrEventBadRequest
}

func isValidSlug(value string) bool {
	return shared.IsValidSlug(value)
}

func (b *Slug) Value() string {
	return b.value
}
