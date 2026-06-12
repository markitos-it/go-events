package gapi

import (
	context "context"

	"govent/internal/domain/services"
	"govent/internal/domain/types"

	"google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

func (s *Server) GetGolden(ctx context.Context, in *GetGoldenRequest) (*GetGoldenResponse, error) {
	if _, err := types.NewGoldenId(in.Id); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	request := services.GoldenOneRequest{Id: in.Id}

	var service = services.NewGoldenOneService(s.repository)
	response, err := service.Do(request)
	if err != nil {
		return nil, status.Error(s.GetGRPCCode(err), err.Error())

	}

	return &GetGoldenResponse{
		Id:      response.Data.Id,
		Name:    response.Data.Name,
		Content: response.Data.Content,
	}, nil
}
