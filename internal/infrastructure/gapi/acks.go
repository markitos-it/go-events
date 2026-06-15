package gapi

import (
	context "context"
	"fmt"
	"go-vents/internal/domain/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) AckMessages(ctx context.Context, req *AckMessagesRequest) (*AckMessagesResponse, error) {
	if len(req.QueueIds) == 0 {
		return nil, status.Error(codes.InvalidArgument, "queue_ids list is empty")
	}

	if len(req.QueueIds) > 100 {
		return nil, status.Error(codes.InvalidArgument, "too many ids, maximum is 100, your request have "+fmt.Sprint(len(req.QueueIds))+" ids")
	}

	var ids []types.Id
	for _, idStr := range req.QueueIds {
		id, err := types.NewId(idStr)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "invalid id format: "+idStr)
		}
		ids = append(ids, *id)
	}

	for _, id := range ids {
		err := s.repository.AckMessage(ctx, &id)
		if err != nil {
			s.logger.Error(fmt.Sprintf("failed to ack message %s: %v", id.Value(), err))
			continue
		}
	}

	return &AckMessagesResponse{
		Success: true,
	}, nil
}
