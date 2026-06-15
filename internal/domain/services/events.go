package services

import (
	"context"
	"go-vents/internal/domain/types"
)

type EventAllResponse struct {
	Data []*types.Event `json:"data"`
}

type EventAllRequest struct {
	Slug   string `json:"slug"`
	Source string `json:"source"`
}

type EventAllService struct {
	Repository types.EventRepository
}

func NewEventAllService(repository types.EventRepository) EventAllService {
	return EventAllService{
		Repository: repository,
	}
}

func (s EventAllService) Do(ctx context.Context, request EventAllRequest) (*EventAllResponse, error) {
	eventSlug, err := types.NewSlug(request.Slug)
	if err != nil {
		return nil, err
	}
	eventSource, err := types.NewSlug(request.Source)
	if err != nil {
		return nil, err
	}

	events, err := s.Repository.AllBySlugAndSource(ctx, eventSlug, eventSource)
	if err != nil {
		return nil, err
	}

	return &EventAllResponse{
		Data: events,
	}, nil
}
