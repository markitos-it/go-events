package gapi_test

import (
	"testing"

	"govent/internal/infrastructure/gapi"

	"github.com/stretchr/testify/require"
)

func TestGoldenCanListAllResources(t *testing.T) {
	golden1 := createPersistedRandomGolden()
	golden2 := createPersistedRandomGolden()

	resp, err := grpcClient.ListGoldens(ctx, &gapi.ListGoldensRequest{})

	require.NoError(t, err)
	require.NotNil(t, resp.Goldens)

	found1, found2 := false, false
	for _, b := range resp.Goldens {
		if b.Id == golden1.Id {
			found1 = true
		}
		if b.Id == golden2.Id {
			found2 = true
		}
	}
	require.True(t, found1, "First golden not found in response")
	require.True(t, found2, "Second golden not found in response")

	deletePersistedRandomGolden(golden1.Id)
	deletePersistedRandomGolden(golden2.Id)
}
