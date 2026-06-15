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

	req := &gapi.CreateEventRequest{
		Slug:    "LoadTestEventBravo",
		Source:  "LoadTestSource",
		Payload: `{"message": "Hello World!"}`,
	}

	log.Printf("📡 Sending request to %s to create event with Slug: '%s' and Source: '%s'...", serverAddr, req.Slug, req.Source)

	res, err := client.CreateEvent(ctx, req)
	if err != nil {
		log.Fatalf("❌ Error calling CreateEvent: %v", err)
	}

	log.Printf("✅ Event created successfully!")
	log.Printf("   ID: %s | Slug: %s | Source: %s | Payload: %s", res.Id, res.Slug, res.Source, res.Payload)
}
