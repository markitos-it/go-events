package main

import (
	"context"
	"log"
	"time"

	"go-vents/internal/infrastructure/gapi"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const HOW_MANY_EVENTS_CREATE = 3
const TEST_EVENT_NAME = "DynamicTestEvent"
const TEST_EVENT_SOURCE = "DynamicTestSource"
const TEST_SUBSCRIBER_NAME = "DynamicTestSubscriber"

func main() {
	serverAddr := "localhost:30000"

	conn, err := grpc.NewClient(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("❌ Error connecting to gRPC server: %v", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("⚠️ Error closing connection: %v", err)
		}
	}()

	client := gapi.NewEventserviceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	createSubscription(ctx, client, TEST_SUBSCRIBER_NAME, TEST_EVENT_NAME, TEST_EVENT_SOURCE)

	eventIDs := createEvents(ctx, client, TEST_EVENT_NAME, TEST_EVENT_SOURCE, HOW_MANY_EVENTS_CREATE)
	time.Sleep(1 * time.Second)

	messages := pullMessages(ctx, client, TEST_SUBSCRIBER_NAME, TEST_EVENT_NAME, TEST_EVENT_SOURCE)
	if len(messages) > 0 {
		ackMessage(ctx, client, messages[0].Id)
	} else {
		log.Println("⚠️ No messages pulled, skipping AckMessage.")
	}

	if len(messages) > 1 {
		var queueIDs []string
		for _, msg := range messages[1:] {
			queueIDs = append(queueIDs, msg.Id)
		}
		ackMessages(ctx, client, queueIDs)
	} else {
		log.Println("⚠️ Not enough messages pulled, skipping AckMessages.")
	}

	deleteEvents(ctx, client, eventIDs)
}

func createSubscription(ctx context.Context, client gapi.EventserviceClient, subscriberName, eventName, source string) {
	log.Println("\n=== 1. CreateSubscription ===")
	subReq := &gapi.CreateSubscriptionRequest{
		SubscriberName: subscriberName,
		EventName:      eventName,
		Source:         source,
	}
	subRes, err := client.CreateSubscription(ctx, subReq)
	if err != nil {
		log.Fatalf("❌ Error calling CreateSubscription: %v", err)
	}
	log.Printf("✅ Subscription created! Success: %t, Message: %s", subRes.Success, subRes.Message)
}

func createEvents(ctx context.Context, client gapi.EventserviceClient, eventName, source string, count int) []string {
	log.Printf("\n=== 2. CreateEvents(%d) ===", count)
	var eventIDs []string
	for range count {
		createReq := &gapi.CreateEventRequest{
			Slug:    eventName,
			Source:  source,
			Payload: `{"message": "Hello from dynamic client!"}`,
		}
		createRes, err := client.CreateEvent(ctx, createReq)
		if err != nil {
			log.Fatalf("❌ Error calling CreateEvent: %v", err)
		}
		log.Printf("✅ Event created with ID: %s", createRes.Id)
		eventIDs = append(eventIDs, createRes.Id)
	}
	return eventIDs
}

func pullMessages(ctx context.Context, client gapi.EventserviceClient, subscriberName, eventName, source string) []*gapi.QueueMessage {
	log.Println("\n=== 3. PullMessages ===")
	pullReq := &gapi.PullMessagesRequest{
		EventName:      eventName,
		Source:         source,
		SubscriberName: subscriberName,
	}
	pullRes, err := client.PullMessages(ctx, pullReq)
	if err != nil {
		log.Fatalf("❌ Error calling PullMessages: %v", err)
	}
	log.Printf("✅ Pulled %d messages", len(pullRes.Messages))
	return pullRes.Messages
}

func ackMessage(ctx context.Context, client gapi.EventserviceClient, queueID string) {
	log.Println("\n=== 4. AckMessage ===")
	ackReq := &gapi.AckMessageRequest{
		QueueId: queueID,
	}
	ackRes, err := client.AckMessage(ctx, ackReq)
	if err != nil {
		log.Printf("⚠️ Error calling AckMessage: %v", err)
	} else {
		log.Printf("✅ Message %s acknowledged successfully! Success: %t", queueID, ackRes.Success)
	}
}

func ackMessages(ctx context.Context, client gapi.EventserviceClient, queueIDs []string) {
	log.Println("\n=== 5. AckMessages ===")
	acksReq := &gapi.AckMessagesRequest{
		QueueIds: queueIDs,
	}
	acksRes, err := client.AckMessages(ctx, acksReq)
	if err != nil {
		log.Printf("⚠️ Error calling AckMessages: %v", err)
	} else {
		log.Printf("✅ %d messages acknowledged successfully! Success: %t", len(queueIDs), acksRes.Success)
	}
}

func deleteEvents(ctx context.Context, client gapi.EventserviceClient, eventIDs []string) {
	log.Println("\n=== 6. DeleteEvents ===")
	for _, id := range eventIDs {
		deleteReq := &gapi.DeleteEventRequest{
			Id: id,
		}
		deleteRes, err := client.DeleteEvent(ctx, deleteReq)
		if err != nil {
			log.Fatalf("❌ Error calling DeleteEvent for ID %s: %v", id, err)
		}
		log.Printf("✅ Event %s deleted successfully! Response: %s", id, deleteRes.Deleted)
	}
}
