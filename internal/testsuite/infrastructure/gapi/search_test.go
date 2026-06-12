package gapi_test

import (
	"testing"

	"govent/internal/domain/types"

	"govent/internal/domain/shared"
	"govent/internal/infrastructure/gapi"
	"govent/internal/testsuite/infrastructure/testdb"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCanSearchWithPattern(t *testing.T) {
	pattern := shared.RandomString(10)

	var ids []string
	for i := 0; i < 5; i++ {
		id := shared.UUIDv4()
		ids = append(ids, id)
		name := pattern + shared.RandomPersonalName()
		golden, _ := types.NewGolden(id, name, shared.RandomString())

		_ = testdb.GetRepository().Create(golden)
	}

	resp, err := grpcClient.SearchGoldens(ctx, &gapi.SearchGoldensRequest{
		SearchTerm: pattern,
		PageNumber: 1,
		PageSize:   6,
	})

	require.NoError(t, err)
	require.Equal(t, 5, len(resp.Goldens))

	foundCount := 0
	for _, b := range resp.Goldens {
		for _, id := range ids {
			if b.Id == id {
				foundCount++
				break
			}
		}
	}
	require.Equal(t, 5, foundCount)

	for _, id := range ids {
		deletePersistedRandomGolden(id)
	}
}

func TestCantSearchWithInvalidOptionalPage(t *testing.T) {
	_, err := grpcClient.SearchGoldens(ctx, &gapi.SearchGoldensRequest{
		PageNumber: -1,
		PageSize:   10,
	})

	require.Error(t, err)
	st, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.InvalidArgument, st.Code())
}
