package services

import (
	"govent/internal/domain/shared"
	"govent/internal/domain/types"
)

type GoldenCreateRequest struct {
	Name       string
	Content    string
	PosterData string
	/* ___CUSTOM_STRUCT_FIELDS___*/
}

type GoldenCreateResponse struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

type GoldenCreateService struct {
	Repository types.GoldenRepository
}

func NewGoldenCreateService(repository types.GoldenRepository) GoldenCreateService {
	return GoldenCreateService{
		Repository: repository,
	}
}

func (s GoldenCreateService) Do(request GoldenCreateRequest) (*GoldenCreateResponse, error) {

	golden, err := types.NewGolden(shared.UUIDv4(), request.Name, request.Content)
	if err != nil {
		return nil, err
	}

	if err := s.Repository.Create(golden); err != nil {
		return nil, err
	}

	return &GoldenCreateResponse{
		Id:      golden.Id,
		Name:    golden.Name,
		Content: golden.Content,
	}, nil
}
