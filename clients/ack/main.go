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
	defer conn.Close()

	client := gapi.NewEventserviceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req := &gapi.AckMessageRequest{
		QueueId: "c997247d-e643-4a9b-aef1-6442a06b0c4a",
	}

	log.Printf("📡 Sending request to %s to ack message with Queue ID: '%s'...", serverAddr, req.QueueId)

	res, err := client.AckMessage(ctx, req)
	if err != nil {
		log.Fatalf("❌ Error calling AckMessage: %v", err)
	}

	log.Printf("✅ Message acknowledged successfully! Success: %t", res.Success)
}
