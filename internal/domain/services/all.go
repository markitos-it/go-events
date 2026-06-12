package services

import "govent/internal/domain/types"

type EventAllResponse struct {
	Data []*types.Event `json:"data"`
}

type EventAllService struct {
	Repository types.EventRepository
}

func NewEventAllService(repository types.EventRepository) EventAllService {
	return EventAllService{
		Repository: repository,
	}
}

func (s EventAllService) Do() (*EventAllResponse, error) {
	events, err := s.Repository.All()
	if err != nil {
		return nil, err
	}

	return &EventAllResponse{
		Data: events,
	}, nil
}
