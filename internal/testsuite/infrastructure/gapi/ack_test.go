package gapi_test

import (
	"log"
	"testing"

	"go-vents/internal/domain/shared"
	"go-vents/internal/infrastructure/gapi"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestAckMessagesSuccess(t *testing.T) {
	subResp, err := grpcClient.CreateSubscription(ctx, &gapi.CreateSubscriptionRequest{
		SubscriberName: "test-subscriber-ack-messages",
		EventName:      "test-event-ack-messages",
		Source:         "test-source-ack",
	})
	require.NoError(t, err)
	require.True(t, subResp.Success)

	eventResp, err := grpcClient.CreateEvent(ctx, &gapi.CreateEventRequest{
		Slug:    "test-event-ack-messages",
		Source:  "test-source-ack",
		Payload: `{"test":"data"}`,
	})
	require.NoError(t, err)
	require.NotNil(t, eventResp)

	pullResp, err := grpcClient.PullMessages(ctx, &gapi.PullMessagesRequest{
		EventName:      "test-event-ack-messages",
		Source:         "test-source-ack",
		SubscriberName: "test-subscriber-ack-messages",
	})
	require.NoError(t, err)
	require.NotEmpty(t, pullResp.Messages)

	var queueIds []string
	for _, msg := range pullResp.Messages {
		if msg.EventId == eventResp.Id {
			queueIds = append(queueIds, msg.Id)
		}
	}
	require.NotEmpty(t, queueIds)

	ackResp, err := grpcClient.AckMessages(ctx, &gapi.AckMessagesRequest{
		QueueIds: queueIds,
	})
	require.NoError(t, err)
	require.NotNil(t, ackResp)
	require.True(t, ackResp.Success)

	deletePersistedRandomEvent(eventResp.Id)
}

func TestAckMessagesEmptyList(t *testing.T) {
	_, err := grpcClient.AckMessages(ctx, &gapi.AckMessagesRequest{
		QueueIds: []string{},
	})

	require.Error(t, err)
	st, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.InvalidArgument, st.Code())
}

func TestAckMessagesInvalidId(t *testing.T) {
	_, err := grpcClient.AckMessages(ctx, &gapi.AckMessagesRequest{
		QueueIds: []string{"invalid-uuid"},
	})

	require.Error(t, err)
	st, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.InvalidArgument, st.Code())
}

func TestAckMessagesNeverReceivedAnError(t *testing.T) {
	nonExistentId := shared.UUIDv4()
	_, err := grpcClient.AckMessages(ctx, &gapi.AckMessagesRequest{
		QueueIds: []string{nonExistentId},
	})

	require.NoError(t, err)

	log.Println("-----------------------------------------------------")
	log.Println("err", err)
	log.Println("-----------------------------------------------------")
}
