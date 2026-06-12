package gapi

import (
	context "context"
	"log"

	"govent/internal/domain/services"
	"govent/internal/domain/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateEvent(ctx context.Context, in *UpdateEventRequest) (*UpdateEventResponse, error) {
	if _, err := types.NewEventId(in.Id); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	var service = services.NewEventUpdateService(s.repository)
	var request = services.EventUpdateRequest{
		Id:      in.Id,
		Name:    in.Name,
		Source:  in.Source,
		Payload: in.Payload,
	}
	if err := service.Do(request); err != nil {
		log.Printf("❌ ERROR (UpdateEvent): %v\n", err)
		return nil, status.Error(s.GetGRPCCode(err), err.Error())
	}

	return &UpdateEventResponse{
		Updated: request.Id,
	}, nil
}
