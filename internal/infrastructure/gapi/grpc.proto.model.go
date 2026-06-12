package gapi

import (
	"govent/internal/domain/types"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewGRPCGolden(in *types.Golden) *Golden {
	return &Golden{
		Id:        in.Id,
		Name:      in.Name,
		CreatedAt: timestamppb.New(in.CreatedAt),
		UpdatedAt: timestamppb.New(in.UpdatedAt),
	}
}
