package gapi

import (
	context "context"

	"go-vents/internal/domain/services"

	"google.golang.org/grpc/status"
)

func (s *Server) CreateEvent(ctx context.Context, req *CreateEventRequest) (*CreateEventResponse, error) {
	var request = services.EventCreateRequest{
		Slug:    req.Slug,
		Source:  req.Source,
		Payload: req.Payload,
	}

	var service = services.NewEventCreateService(s.repository)
	entity, err := service.Do(ctx, request)
	if err != nil {
		return nil, status.Error(s.GetGRPCCode(err), err.Error())
	}

	return &CreateEventResponse{
		Id:      entity.Id,
		Slug:    entity.Slug,
		Source:  entity.Source,
		Payload: entity.Payload,
	}, nil
}
