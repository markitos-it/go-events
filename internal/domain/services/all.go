package services

import "govent/internal/domain/types"

type GoldenAllResponse struct {
	Data []*types.Golden `json:"data"`
}

type GoldenAllService struct {
	Repository types.GoldenRepository
}

func NewGoldenAllService(repository types.GoldenRepository) GoldenAllService {
	return GoldenAllService{
		Repository: repository,
	}
}

func (s GoldenAllService) Do() (*GoldenAllResponse, error) {
	goldens, err := s.Repository.All()
	if err != nil {
		return nil, err
	}

	return &GoldenAllResponse{
		Data: goldens,
	}, nil
}
