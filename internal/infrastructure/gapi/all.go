package gapi

import (
	context "context"

	"go-vents/internal/domain/services"

	status "google.golang.org/grpc/status"
)

func (s *Server) AllBySlugAndSource(ctx context.Context, in *AllEventsBySlugAndSourceRequest) (*AllEventsBySlugAndSourceResponse, error) {
	var request = services.EventAllRequest{
		Slug:   in.Slug,
		Source: in.Source,
	}
	var service = services.NewEventAllService(s.repository)
	response, err := service.Do(ctx, request)
	if err != nil {
		return nil, status.Error(s.GetGRPCCode(err), err.Error())
	}

	return &AllEventsBySlugAndSourceResponse{
		Events: s.ToProtos(response.Data),
	}, nil
}
