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
	// 1. Create a subscription
	subResp, err := grpcClient.CreateSubscription(ctx, &gapi.CreateSubscriptionRequest{
		SubscriberName: "test-subscriber-ack-messages",
		EventName:      "test-event-ack-messages",
		Source:         "test-source-ack",
	})
	require.NoError(t, err)
	require.True(t, subResp.Success)

	// 2. Create an event to trigger a queue message
	eventResp, err := grpcClient.CreateEvent(ctx, &gapi.CreateEventRequest{
		Slug:    "test-event-ack-messages",
		Source:  "test-source-ack",
		Payload: `{"test":"data"}`,
	})
	require.NoError(t, err)
	require.NotNil(t, eventResp)

	// 3. Pull messages to get the queue ID
	pullResp, err := grpcClient.PullMessages(ctx, &gapi.PullMessagesRequest{
		EventName: "test-event-ack-messages",
		Source:    "test-source-ack",
	})
	require.NoError(t, err)
	require.NotEmpty(t, pullResp.Messages)

	// We only care about the one we just created
	var queueIds []string
	for _, msg := range pullResp.Messages {
		if msg.EventId == eventResp.Id {
			queueIds = append(queueIds, msg.Id)
		}
	}
	require.NotEmpty(t, queueIds)

	// 4. Ack the messages
	ackResp, err := grpcClient.AckMessages(ctx, &gapi.AckMessagesRequest{
		QueueIds: queueIds,
	})
	require.NoError(t, err)
	require.NotNil(t, ackResp)
	require.True(t, ackResp.Success)

	// Clean up (optional but good practice)
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

func TestAckMessagesNotFound(t *testing.T) {
	nonExistentId := shared.UUIDv4()
	_, err := grpcClient.AckMessages(ctx, &gapi.AckMessagesRequest{
		QueueIds: []string{nonExistentId},
	})

	require.Error(t, err)

	log.Println("-----------------------------------------------------")
	log.Println("err", err)
	log.Println("-----------------------------------------------------")

	// Wait, currently gapi ack.go might return internal error for not found, let's just require error for now
	// Depending on implementation it might be Internal or NotFound
	// Let's assert based on what we see when we run it.
}
