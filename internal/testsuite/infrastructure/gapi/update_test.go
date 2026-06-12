package gapi_test

import (
	"testing"

	"govent/internal/domain/shared"
	"govent/internal/infrastructure/gapi"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCanUpdateAEvent(t *testing.T) {
	event := createPersistedRandomEvent()
	updatedName := event.Name + " UPDATED"

	resp, err := grpcClient.UpdateEvent(ctx, &gapi.UpdateEventRequest{
		Id:   event.Id,
		Name: updatedName,
		/* ___CUSTOM_TEST_FIELDS___*/
	})

	require.NoError(t, err)
	require.Equal(t, event.Id, resp.Updated)

	getResp, _ := grpcClient.GetEvent(ctx, &gapi.GetEventRequest{Id: event.Id})
	require.Equal(t, updatedName, getResp.Name)

	deletePersistedRandomEvent(event.Id)
}

func TestCantUpdateANonExistingEvent(t *testing.T) {
	randomId := shared.UUIDv4()
	_, err := grpcClient.UpdateEvent(ctx, &gapi.UpdateEventRequest{
		Id:   randomId,
		Name: shared.RandomPersonalName(),
		/* ___CUSTOM_TEST_FIELDS___*/
	})

	require.Error(t, err)
	st, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.NotFound, st.Code())
}

func TestCantUpdateAnInvalidEventId(t *testing.T) {
	_, err := grpcClient.UpdateEvent(ctx, &gapi.UpdateEventRequest{
		Id:   "an-invalid-id-format",
		Name: shared.RandomPersonalName(),
		/* ___CUSTOM_TEST_FIELDS___*/
	})

	require.Error(t, err)
	st, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.InvalidArgument, st.Code())
}
