package internal_test

import (
	"govent/internal/domain/shared"
	"govent/internal/domain/types"
)

func NewRandomGolden() *types.Golden {
	golden, _ := types.NewGolden(
		shared.UUIDv4(),
		shared.RandomString(),
		shared.RandomString(),
	)

	return golden
}

func NewRandomOnlyNameGolden() *types.Golden {
	golden, _ := types.NewGolden(shared.UUIDv4(), shared.RandomString(), "")

	return golden
}

func NewRandomGoldenWithCustomId(goldenId *types.GoldenId) *types.Golden {
	golden, _ := types.NewGolden(
		goldenId.Value(),
		shared.RandomString(),
		shared.RandomString(),
	)
	return golden
}
