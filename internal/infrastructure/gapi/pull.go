package gapi

import (
	context "context"

	"go-vents/internal/domain/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const DEFAULT_LIMIT_PULL = 10

func (s *Server) PullMessages(ctx context.Context, req *PullMessagesRequest) (*PullMessagesResponse, error) {
	subscriberName, err := types.NewSlug(req.GetSubscriberName())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	slug, err := types.NewSlug(req.GetEventName())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	source, err := types.NewSlug(req.GetSource())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	dbMessages, err := s.repository.PullMessages(ctx, subscriberName, slug, source)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error en la base de datos al hacer pull: %v", err)
	}

	var protoMessages []*QueueMessage
	for _, msg := range dbMessages {
		protoMessages = append(protoMessages, &QueueMessage{
			Id:             msg.Id,
			EventId:        msg.EventId,
			SubscriberName: msg.SubscriberName,
			Status:         string(msg.Status),
			Payload:        msg.Payload,
		})
	}

	return &PullMessagesResponse{
		Messages: protoMessages,
	}, nil
}
