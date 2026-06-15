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

	req := &gapi.AckMessagesRequest{
		QueueIds: []string{
			"5074b99b-57e4-419a-b70a-f96d2f5c3403",
			"7c5fd3e3-ce6b-4213-856a-f25d589cbcf3",
			"81b9de9e-104e-4ae9-b721-eb68bcca2d9f",
			"8bf5c330-a81f-43a7-a825-638d767bf082",
			"1120cadd-5b6c-4f1c-bec8-f39879b1cff4",
		},
	}

	log.Printf("📡 Sending request to %s to ack messages with Queue IDs: %v...", serverAddr, req.QueueIds)

	res, err := client.AckMessages(ctx, req)
	if err != nil {
		log.Fatalf("❌ Error calling AckMessages: %v", err)
	}

	log.Printf("✅ Messages acknowledged successfully! Success: %t", res.Success)
}
