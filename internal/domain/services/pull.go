package services

import (
	"context"
	"go-vents/internal/domain/types"
)

type PullResponse struct {
	Data []*types.Queue `json:"data"`
}

type PullRequest struct {
	SubscriberName string `json:"subscriber_name"`
	Source         string `json:"source"`
	Slug           string `json:"slug"`
}

type PullService struct {
	Repository types.EventRepository
}

func NewPullService(repository types.EventRepository) PullService {
	return PullService{
		Repository: repository,
	}
}

func (s PullService) Do(ctx context.Context, request PullRequest) (*PullResponse, error) {
	subscriberName, err := types.NewName(request.SubscriberName)
	if err != nil {
		return nil, err
	}

	eventSlug, err := types.NewSlug(request.Slug)
	if err != nil {
		return nil, err
	}
	eventSource, err := types.NewSource(request.Source)
	if err != nil {
		return nil, err
	}

	events, err := s.Repository.PullMessages(ctx, subscriberName, eventSlug, eventSource)
	if err != nil {
		return nil, err
	}

	return &PullResponse{
		Data: events,
	}, nil
}
