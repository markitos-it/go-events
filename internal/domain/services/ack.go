package services

import (
	"context"
	"go-vents/internal/domain/types"
)

type AckMessageResponse struct {
	Success bool `json:"success"`
}

type AckMessageRequest struct {
	QueueId string `json:"queue_id"`
}

type AckMessageService struct {
	Repository types.EventRepository
}

func NewAckMessageService(repository types.EventRepository) AckMessageService {
	return AckMessageService{
		Repository: repository,
	}
}

func (s AckMessageService) Do(ctx context.Context, request AckMessageRequest) error {
	id, err := types.NewId(request.QueueId)
	if err != nil {
		return err
	}

	err = s.Repository.AckMessage(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
