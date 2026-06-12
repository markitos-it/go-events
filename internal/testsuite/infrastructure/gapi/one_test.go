package gapi_test

import (
	"testing"

	"govent/internal/domain/shared"
	"govent/internal/infrastructure/gapi"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGoldenCanGetOne(t *testing.T) {
	golden := createPersistedRandomGolden()

	resp, err := grpcClient.GetGolden(ctx, &gapi.GetGoldenRequest{
		Id: golden.Id,
	})

	require.NoError(t, err)
	require.Equal(t, golden.Name, resp.Name)
	require.Equal(t, golden.Id, resp.Id)

	deletePersistedRandomGolden(resp.Id)
}

func TestGoldenCantGetInvalidId(t *testing.T) {
	_, err := grpcClient.GetGolden(ctx, &gapi.GetGoldenRequest{
		Id: "an-invalid-id",
	})

	require.Error(t, err)
	st, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.InvalidArgument, st.Code())
}

func TestGoldenCantGetValidIdButNonExistingResource(t *testing.T) {
	_, err := grpcClient.GetGolden(ctx, &gapi.GetGoldenRequest{
		Id: shared.UUIDv4(),
	})

	require.Error(t, err)
	st, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.NotFound, st.Code())
}
