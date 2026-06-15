package gapi

import (
	context "context"

	"go-vents/internal/domain/services"

	"google.golang.org/grpc/status"
)

func (s *Server) AckMessage(ctx context.Context, req *AckMessageRequest) (*AckMessageResponse, error) {

	var request = services.AckMessageRequest{
		QueueId: req.QueueId,
	}

	var service = services.NewAckMessageService(s.repository)
	err := service.Do(ctx, request)
	if err != nil {
		return nil, status.Error(s.GetGRPCCode(err), err.Error())
	}

	return &AckMessageResponse{
		Success: true,
	}, nil
}
