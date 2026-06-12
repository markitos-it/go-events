package types

import (
	"govent/internal/domain/shared"
)

type GoldenContent struct {
	value string
}

const GOLDEN_CONTENT_MAX_LENGTH = 200
const GOLDEN_CONTENT_MIN_LENGTH = 1

func NewGoldenContent(value string) (*GoldenContent, error) {

	if isValidGoldenContent(value) {
		return &GoldenContent{value}, nil
	}

	return nil, shared.ErrGoldenBadRequest
}

func isValidGoldenContent(value string) bool {
	return len(value) >= GOLDEN_CONTENT_MIN_LENGTH || len(value) <= GOLDEN_CONTENT_MAX_LENGTH
}

func (b *GoldenContent) Value() string {
	return b.value
}
