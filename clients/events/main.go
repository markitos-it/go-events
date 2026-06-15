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

	req := &gapi.AllEventsBySlugAndSourceRequest{
		Slug:   "LoadTestEventBravo",
		Source: "LoadTestSource",
	}

	log.Printf("📡 Sending request to %s to get events with Slug: '%s' and Source: '%s'...", serverAddr, req.Slug, req.Source)

	res, err := client.AllBySlugAndSource(ctx, req)
	if err != nil {
		log.Fatalf("❌ Error calling AllBySlugAndSource: %v", err)
	}

	events := res.GetEvents()
	if len(events) == 0 {
		log.Println("ℹ️ No events found for the provided slug and source.")
		return
	}

	log.Printf("✅ Found %d events:", len(events))
	for i, ev := range events {
		log.Printf("   [%d] ID: %s | Slug: %s | Source: %s | Payload: %s", i+1, ev.Id, ev.Slug, ev.Source, ev.Payload)
	}
}
