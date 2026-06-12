package services

import (
	"govent/internal/domain/types"
)

type GoldenUpdateRequest struct {
	Id         string `json:"id" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Content    string `json:"content"`
	PosterData string `json:"poster_data"`
}

type GoldenUpdateService struct {
	Repository types.GoldenRepository
}

func NewGoldenUpdateService(repository types.GoldenRepository) GoldenUpdateService {
	return GoldenUpdateService{
		Repository: repository,
	}
}

func (s GoldenUpdateService) Do(request GoldenUpdateRequest) error {
	securedId, err := types.NewGoldenId(request.Id)
	if err != nil {
		return err
	}

	securedName, err := types.NewGoldenName(request.Name)
	if err != nil {
		return err
	}

	securedContent, err := types.NewGoldenContent(request.Content)
	if err != nil {
		return err
	}

	golden, err := s.Repository.One(securedId)
	if err != nil {
		return err
	}

	golden.Name = securedName.Value()
	golden.Content = securedContent.Value()

	return s.Repository.Update(golden)
}
