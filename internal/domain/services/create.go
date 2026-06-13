package services

import (
	"context"

	"go-vents/internal/domain/shared"
	"go-vents/internal/domain/types"
)

type EventCreateRequest struct {
	Slug    string
	Source  string
	Payload string
}

type EventCreateResponse struct {
	Id      string `json:"id"`
	Slug    string `json:"slug"`
	Source  string `json:"source"`
	Payload string `json:"payload"`
}

type EventCreateService struct {
	Repository types.EventRepository
}

func NewEventCreateService(repository types.EventRepository) EventCreateService {
	return EventCreateService{
		Repository: repository,
	}
}

func (s EventCreateService) Do(ctx context.Context, request EventCreateRequest) (*EventCreateResponse, error) {

	event, err := types.NewEvent(shared.UUIDv4(), request.Slug, request.Source, request.Payload)
	if err != nil {
		return nil, err
	}

	if err := s.Repository.Create(ctx, event); err != nil {
		return nil, err
	}

	return &EventCreateResponse{
		Id:      event.Id,
		Slug:    event.Slug,
		Source:  event.Source,
		Payload: event.Payload,
	}, nil
}
