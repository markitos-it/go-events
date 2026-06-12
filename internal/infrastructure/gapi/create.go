package gapi

import (
	context "context"
	"log"

	"govent/internal/domain/services"

	"google.golang.org/grpc/status"
)

func (s *Server) CreateGolden(ctx context.Context, req *CreateGoldenRequest) (*CreateGoldenResponse, error) {
	var request = services.GoldenCreateRequest{
		Name:    req.Name,
		Content: req.Content,
	}

	var service = services.NewGoldenCreateService(s.repository)
	entity, err := service.Do(request)
	if err != nil {
		log.Printf("❌ ERROR (CreateGolden): %v\n", err)
		return nil, status.Error(s.GetGRPCCode(err), err.Error())
	}

	return &CreateGoldenResponse{
		Id:      entity.Id,
		Name:    entity.Name,
		Content: entity.Content,
	}, nil
}
