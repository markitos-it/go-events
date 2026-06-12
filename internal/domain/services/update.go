package services

import (
	"govent/internal/domain/types"
)

type EventUpdateRequest struct {
	Id      string `json:"id" binding:"required"`
	Name    string `json:"name" binding:"required"`
	Source  string `json:"source" binding:"required"`
	Payload string `json:"payload" binding:"required"`
}

type EventUpdateService struct {
	Repository types.EventRepository
}

func NewEventUpdateService(repository types.EventRepository) EventUpdateService {
	return EventUpdateService{
		Repository: repository,
	}
}

func (s EventUpdateService) Do(request EventUpdateRequest) error {
	securedId, err := types.NewEventId(request.Id)
	if err != nil {
		return err
	}

	securedName, err := types.NewEventName(request.Name)
	if err != nil {
		return err
	}

	securedSource, err := types.NewEventSource(request.Source)
	if err != nil {
		return err
	}

	securedPayload, err := types.NewEventPayload(request.Payload)
	if err != nil {
		return err
	}

	event, err := s.Repository.One(securedId)
	if err != nil {
		return err
	}

	event.Name = securedName.Value()
	event.Source = securedSource.Value()
	event.Payload = securedPayload.Value()

	return s.Repository.Update(event)
}
