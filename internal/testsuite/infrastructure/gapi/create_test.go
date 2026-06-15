package gapi_test

import (
	"testing"

	"go-vents/internal/domain/shared"
	"go-vents/internal/infrastructure/gapi"
	internal_test "go-vents/internal/testsuite/internal"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestEventCanCreate(t *testing.T) {
	event := internal_test.NewRandomEvent()

	resp, err := grpcClient.CreateEvent(ctx, &gapi.CreateEventRequest{
		Slug:    event.Slug,
		Source:  event.Source,
		Payload: event.Payload,
	})

	require.NoError(t, err)
	require.NotEmpty(t, resp.Id)
	require.Equal(t, event.Slug, resp.Slug)

	deletePersistedRandomEvent(resp.Id)
}

func TestEventCantCreateWithoutName(t *testing.T) {
	_, err := grpcClient.CreateEvent(ctx, &gapi.CreateEventRequest{
		Slug:    "",
		Source:  shared.Slug(),
		Payload: "{}",
	})

	require.Error(t, err)
	st, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.InvalidArgument, st.Code())
}

func TestEventCantCreateWithoutValidName(t *testing.T) {
	_, err := grpcClient.CreateEvent(ctx, &gapi.CreateEventRequest{
		Slug:    "!!!!!invalid!!!slug!!!",
		Source:  shared.Slug(),
		Payload: "{}",
	})

	require.Error(t, err)
	st, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.InvalidArgument, st.Code())
}
