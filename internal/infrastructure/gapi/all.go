package gapi

import (
	context "context"

	"govent/internal/domain/services"

	status "google.golang.org/grpc/status"
)

func (s *Server) ListEvents(ctx context.Context, in *ListEventsRequest) (*ListEventsResponse, error) {
	var service = services.NewEventAllService(s.repository)
	response, err := service.Do()
	if err != nil {
		return nil, status.Error(s.GetGRPCCode(err), err.Error())
	}

	return &ListEventsResponse{
		Events: s.ToProtos(response.Data),
	}, nil
}
