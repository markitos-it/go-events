package main

import (
	"context"
	"log"
	"time"

	"go-vents/internal/infrastructure/gapi"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req := &gapi.PullMessagesRequest{
		EventName: "LoadTestEventBravo",
		Source:    "LoadTestSource",
	}

	log.Printf("📡 Sending request to %s to pull events with Slug: '%s' and Source: '%s'...", serverAddr, req.EventName, req.Source)

	res, err := client.PullMessages(ctx, req)
	if err != nil {
		log.Fatalf("❌ Error calling PullMessages: %v", err)
	}

	queue := res.Messages
	if len(queue) == 0 {
		log.Println("ℹ️ No queue messages found for the provided slug and source.")
		return
	}

	log.Printf("✅ Found %d queue:", len(queue))
	for i, queueMessage := range queue {
		log.Printf("   [%d] ID: %s | Status: %s | Subscriber: %s | Payload: %s",
			i+1,
			queueMessage.Id,
			queueMessage.Status,
			queueMessage.SubscriberName,
			queueMessage.Payload)
	}
}
