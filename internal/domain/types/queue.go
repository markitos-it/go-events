package types

import (
	"log"
	"time"
)

type MessageStatus string

const (
	StatusPending   MessageStatus = "pending"
	StatusProcessed MessageStatus = "processed"
	StatusFailed    MessageStatus = "failed"
)

type Queue struct {
	Id             string        `json:"id" binding:"required,uuid"`
	SubscriberName string        `json:"subscriber_name" binding:"required"`
	EventId        string        `json:"event_id" binding:"required,uuid"`
	Payload        string        `json:"payload" binding:"required"`
	Status         MessageStatus `json:"status" binding:"required"`
	CreatedAt      time.Time     `json:"created_at" binding:"required,datetime" default:"now"`
	UpdatedAt      time.Time     `json:"updated_at" binding:"required,datetime" default:"now"`
}

func (Queue) TableName() string {
	return "queue"
}

func NewQueueMessage(id, subscriberName, eventId, payload string) (*Queue, error) {
	secureId, err := NewId(id)
	if err != nil {
		log.Printf("❌ DEBUG ERROR (Id): %v\n", err)
		return nil, err
	}

	secureEventId, err := NewId(eventId)
	if err != nil {
		log.Printf("❌ DEBUG ERROR (event_id): %v\n", err)
		return nil, err
	}

	secureSubscriberName, err := NewSlug(subscriberName)
	if err != nil {
		log.Printf("❌ DEBUG ERROR (subscriber_name): %v\n", err)
		return nil, err
	}

	securePayload, err := NewPayload(payload)
	if err != nil {
		log.Printf("❌ DEBUG ERROR (payload): %v\n", err)
		return nil, err
	}

	return &Queue{
		Id:             secureId.Value(),
		SubscriberName: secureSubscriberName.Value(),
		EventId:        secureEventId.Value(),
		Payload:        securePayload.Value(),
		Status:         StatusPending,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}, nil
}
