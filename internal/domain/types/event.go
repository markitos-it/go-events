package types

import (
	"log"
	"time"
)

type Event struct {
	Id        string    `json:"id" binding:"required,uuid"`
	Slug      string    `json:"slug" binding:"required"`
	Source    string    `json:"source" binding:"required"`
	Payload   string    `json:"payload" binding:"required"`
	CreatedAt time.Time `json:"created_at" binding:"required,datetime" default:"now"`
	UpdatedAt time.Time `json:"updated_at" binding:"required,datetime" default:"now"`
}

func NewEvent(id, slug, source, payload string) (*Event, error) {
	secureId, err := NewId(id)
	if err != nil {
		log.Printf("❌ DEBUG ERROR (Id): %v\n", err)
		return nil, err
	}

	secureSlug, err := NewSlug(slug)
	if err != nil {
		log.Printf("❌ DEBUG ERROR (Slug): %v\n", err)
		return nil, err
	}

	secureSource, err := NewSource(source)
	if err != nil {
		log.Printf("❌ DEBUG ERROR (Source): %v\n", err)
		return nil, err
	}

	securePayload, err := NewPayload(payload)
	if err != nil {
		log.Printf("❌ DEBUG ERROR (Payload): %v\n", err)
		return nil, err
	}

	return &Event{
		Id:        secureId.Value(),
		Slug:      secureSlug.Value(),
		Source:    secureSource.Value(),
		Payload:   securePayload.Value(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (e *Event) GetSlug() *Slug {
	result, _ := NewSlug(e.Slug)

	return result
}

func (e *Event) GetSource() *Source {
	result, _ := NewSource(e.Source)

	return result
}
