package services

import "govent/internal/domain/types"

type GoldensearchResponse struct {
	Data []*types.Golden `json:"data"`
}

type GoldensearchRequest struct {
	SearchTerm string `json:"searchTerm"`
	PageNumber int    `json:"pageNumber" bindings:"min=1"`
	PageSize   int    `json:"pageSize" bindings:"min=10,max=100"`
}

type GoldensearchService struct {
	Repository types.GoldenRepository
}

func NewGoldensearchService(repository types.GoldenRepository) GoldensearchService {
	return GoldensearchService{
		Repository: repository,
	}
}

func (s GoldensearchService) Do(request GoldensearchRequest) (*GoldensearchResponse, error) {
	response, err := s.Repository.SearchAndPaginate(request.SearchTerm, request.PageNumber, request.PageSize)
	if err != nil {
		return nil, err
	}

	return &GoldensearchResponse{
		Data: response,
	}, nil
}
