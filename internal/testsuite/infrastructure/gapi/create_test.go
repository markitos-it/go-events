package gapi_test

import (
	"testing"

	"go-vents/internal/infrastructure/gapi"
	internal_test "go-vents/internal/testsuite/internal"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestEventCanCreate(t *testing.T) {
	event := internal_test.NewRandomOnlyNameEvent()

	resp, err := grpcClient.CreateEvent(ctx, &gapi.CreateEventRequest{
		Slug: event.Slug,
	})

	require.NoError(t, err)
	require.NotEmpty(t, resp.Id)
	require.Equal(t, event.Slug, resp.Slug)

	deletePersistedRandomEvent(resp.Id)
}

func TestEventCantCreateWithoutName(t *testing.T) {
	_, err := grpcClient.CreateEvent(ctx, &gapi.CreateEventRequest{
		Slug: "",
		/* ___CUSTOM_REQUIRED_VALUES___*/
	})

	require.Error(t, err)
	st, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.InvalidArgument, st.Code())
}

func TestEventCantCreateWithoutValidName(t *testing.T) {
	_, err := grpcClient.CreateEvent(ctx, &gapi.CreateEventRequest{
		Slug: "!!!!!invalid!!!slug!!!",
	})

	require.Error(t, err)
	st, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.InvalidArgument, st.Code())
}
