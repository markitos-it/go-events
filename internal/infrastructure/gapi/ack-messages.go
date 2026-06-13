package gapi

import (
	context "context"
	"go-vents/internal/domain/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) AckMessages(ctx context.Context, req *AckMessagesRequest) (*AckMessagesResponse, error) {
	if len(req.QueueIds) == 0 {
		return nil, status.Error(codes.InvalidArgument, "queue_ids list is empty")
	}

	for _, idStr := range req.QueueIds {
		id, err := types.NewSharedId(idStr)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "invalid id format: "+idStr)
		}

		err = s.repository.AckMessage(ctx, id)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to ack message %s: %v", idStr, err)
		}
	}

	return &AckMessagesResponse{
		Success: true,
	}, nil
}
