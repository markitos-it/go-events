package services

import (
	"govent/internal/domain/types"
)

type GoldenOneRequest struct {
	Id string `json:"id"`
}

type GoldenOneResponse struct {
	Data *types.Golden `json:"data"`
}

type GoldenOneService struct {
	Repository types.GoldenRepository
}

func NewGoldenOneService(repository types.GoldenRepository) GoldenOneService {
	return GoldenOneService{
		Repository: repository,
	}
}

func (s GoldenOneService) Do(request GoldenOneRequest) (*GoldenOneResponse, error) {
	securedId, err := types.NewGoldenId(request.Id)
	if err != nil {
		return nil, err
	}

	golden, err := s.Repository.One(securedId)
	if err != nil {
		return nil, err
	}

	return &GoldenOneResponse{golden}, nil
}
