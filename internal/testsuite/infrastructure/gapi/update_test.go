package gapi_test

import (
	"testing"

	"govent/internal/domain/shared"
	"govent/internal/infrastructure/gapi"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCanUpdateAGolden(t *testing.T) {
	golden := createPersistedRandomGolden()
	updatedName := golden.Name + " UPDATED"

	resp, err := grpcClient.UpdateGolden(ctx, &gapi.UpdateGoldenRequest{
		Id:   golden.Id,
		Name: updatedName,
		/* ___CUSTOM_TEST_FIELDS___*/
	})

	require.NoError(t, err)
	require.Equal(t, golden.Id, resp.Updated)

	getResp, _ := grpcClient.GetGolden(ctx, &gapi.GetGoldenRequest{Id: golden.Id})
	require.Equal(t, updatedName, getResp.Name)

	deletePersistedRandomGolden(golden.Id)
}

func TestCantUpdateANonExistingGolden(t *testing.T) {
	randomId := shared.UUIDv4()
	_, err := grpcClient.UpdateGolden(ctx, &gapi.UpdateGoldenRequest{
		Id:   randomId,
		Name: shared.RandomPersonalName(),
		/* ___CUSTOM_TEST_FIELDS___*/
	})

	require.Error(t, err)
	st, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.NotFound, st.Code())
}

func TestCantUpdateAnInvalidGoldenId(t *testing.T) {
	_, err := grpcClient.UpdateGolden(ctx, &gapi.UpdateGoldenRequest{
		Id:   "an-invalid-id-format",
		Name: shared.RandomPersonalName(),
		/* ___CUSTOM_TEST_FIELDS___*/
	})

	require.Error(t, err)
	st, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.InvalidArgument, st.Code())
}
