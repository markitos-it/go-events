package gapi

import (
	context "context"

	"govent/internal/domain/services"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

func (s *Server) SearchEvents(ctx context.Context, in *SearchEventsRequest) (*SearchEventsResponse, error) {
	if in.PageNumber < 1 {
		return nil, status.Error(codes.InvalidArgument, "invalid page number")
	}

	if in.PageSize < 1 {
		return nil, status.Error(codes.InvalidArgument, "invalid page size")
	}

	var service = services.NewEventsearchService(s.repository)
	var request = services.EventsearchRequest{
		SearchTerm: in.SearchTerm,
		PageNumber: int(in.PageNumber),
		PageSize:   int(in.PageSize),
	}

	response, err := service.Do(request)
	if err != nil {
		return nil, status.Error(s.GetGRPCCode(err), err.Error())
	}

	domainEvents := response.Data
	grpcCollection := s.ToProtos(domainEvents)

	return &SearchEventsResponse{
		Events: grpcCollection,
	}, nil
}
