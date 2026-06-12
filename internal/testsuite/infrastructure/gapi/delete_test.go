package gapi_test

import (
	"testing"

	"govent/internal/domain/shared"
	"govent/internal/infrastructure/gapi"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGoldenCanDelete(t *testing.T) {
	golden := createPersistedRandomGolden()

	resp, err := grpcClient.DeleteGolden(ctx, &gapi.DeleteGoldenRequest{
		Id: golden.Id,
	})

	require.NoError(t, err)
	require.Equal(t, golden.Id, resp.Deleted)

	_, err = grpcClient.GetGolden(ctx, &gapi.GetGoldenRequest{Id: golden.Id})
	require.Error(t, err)
}

func TestGoldenCantDeleteValidButNonExistingId(t *testing.T) {
	_, err := grpcClient.DeleteGolden(ctx, &gapi.DeleteGoldenRequest{
		Id: shared.UUIDv4(),
	})

	require.Error(t, err)
	st, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.NotFound, st.Code())
}

func TestGoldenCantDeleteInvalidGoldenId(t *testing.T) {
	_, err := grpcClient.DeleteGolden(ctx, &gapi.DeleteGoldenRequest{
		Id: "an-invalid-id-format",
	})

	require.Error(t, err)
	st, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.InvalidArgument, st.Code())
}
