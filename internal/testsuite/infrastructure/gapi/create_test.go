package gapi_test

import (
	"testing"

	"govent/internal/infrastructure/gapi"
	internal_test "govent/internal/testsuite/internal"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGoldenCanCreate(t *testing.T) {
	golden := internal_test.NewRandomOnlyNameGolden()

	resp, err := grpcClient.CreateGolden(ctx, &gapi.CreateGoldenRequest{
		Name: golden.Name,
		/* ___CUSTOM_TEST_FIELDS___*/
	})

	require.NoError(t, err)
	require.NotEmpty(t, resp.Id)
	require.Equal(t, golden.Name, resp.Name)

	deletePersistedRandomGolden(resp.Id)
}

func TestGoldenCantCreateWithoutName(t *testing.T) {
	_, err := grpcClient.CreateGolden(ctx, &gapi.CreateGoldenRequest{
		Name: "",
		/* ___CUSTOM_REQUIRED_VALUES___*/
	})

	require.Error(t, err)
	st, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.InvalidArgument, st.Code())
}

func TestGoldenCantCreateWithoutValidName(t *testing.T) {
	_, err := grpcClient.CreateGolden(ctx, &gapi.CreateGoldenRequest{
		Name: "!!!!!invalid!!!name!!!",
		/* ___CUSTOM_REQUIRED_VALUES___*/
	})

	require.Error(t, err)
	st, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.InvalidArgument, st.Code())
}
