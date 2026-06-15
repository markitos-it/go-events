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

	createReq := &gapi.CreateEventRequest{
		Slug:    "EventToDelete",
		Source:  "DeleteTest",
		Payload: `{"message": "I will be deleted"}`,
	}

	log.Printf("📡 Creating event...")
	createRes, err := client.CreateEvent(ctx, createReq)
	if err != nil {
		log.Fatalf("❌ Error calling CreateEvent: %v", err)
	}
	log.Printf("✅ Event created with ID: %s", createRes.Id)

	deleteReq := &gapi.DeleteEventRequest{
		Id: createRes.Id,
	}

	log.Printf("📡 Deleting event with ID: '%s'...", deleteReq.Id)
	deleteRes, err := client.DeleteEvent(ctx, deleteReq)
	if err != nil {
		log.Fatalf("❌ Error calling DeleteEvent: %v", err)
	}

	log.Printf("✅ Event deleted successfully! Response: %s", deleteRes.Deleted)
}
