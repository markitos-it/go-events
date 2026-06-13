package gapi

import (
	"go-vents/internal/domain/types"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewGRPCEvent(in *types.Event) *Event {
	return &Event{
		Id:        in.Id,
		Slug:      in.Slug,
		CreatedAt: timestamppb.New(in.CreatedAt),
		UpdatedAt: timestamppb.New(in.UpdatedAt),
	}
}
