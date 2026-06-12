package services

import "govent/internal/domain/types"

type EventsearchResponse struct {
	Data []*types.Event `json:"data"`
}

type EventsearchRequest struct {
	SearchTerm string `json:"searchTerm"`
	PageNumber int    `json:"pageNumber" bindings:"min=1"`
	PageSize   int    `json:"pageSize" bindings:"min=10,max=100"`
}

type EventsearchService struct {
	Repository types.EventRepository
}

func NewEventsearchService(repository types.EventRepository) EventsearchService {
	return EventsearchService{
		Repository: repository,
	}
}

func (s EventsearchService) Do(request EventsearchRequest) (*EventsearchResponse, error) {
	response, err := s.Repository.SearchAndPaginate(request.SearchTerm, request.PageNumber, request.PageSize)
	if err != nil {
		return nil, err
	}

	return &EventsearchResponse{
		Data: response,
	}, nil
}
