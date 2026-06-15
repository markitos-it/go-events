package types

import (
	"log"
	"time"
)

type Subscription struct {
	Id             string    `json:"id" binding:"required,uuid"`
	SubscriberName string    `json:"subscriber_name" binding:"required"`
	EventName      string    `json:"event_name" binding:"required"`
	Source         string    `json:"source" binding:"required"`
	CreatedAt      time.Time `json:"created_at" binding:"required,datetime" default:"now"`
	UpdatedAt      time.Time `json:"updated_at" binding:"required,datetime" default:"now"`
}

func NewSubscription(id, subscriberName, eventName, source string) (*Subscription, error) {
	secureId, err := NewId(id)
	if err != nil {
		log.Printf("❌ DEBUG ERROR (Id): %v\n", err)
		return nil, err
	}

	secureEventName, err := NewSlug(eventName)
	if err != nil {
		log.Printf("❌ DEBUG ERROR (event_name): %v\n", err)
		return nil, err
	}

	secureSubscriberName, err := NewSlug(subscriberName)
	if err != nil {
		log.Printf("❌ DEBUG ERROR (subscriber_name): %v\n", err)
		return nil, err
	}

	secureSource, err := NewSlug(source)
	if err != nil {
		log.Printf("❌ DEBUG ERROR (source): %v\n", err)
		return nil, err
	}

	return &Subscription{
		Id:             secureId.Value(),
		SubscriberName: secureSubscriberName.Value(),
		EventName:      secureEventName.Value(),
		Source:         secureSource.Value(),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}, nil
}
