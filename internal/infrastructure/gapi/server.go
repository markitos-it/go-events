package gapi

import (
	"errors"
	"strings"

	"govent/internal/domain/shared"
	"govent/internal/domain/types"
	"govent/internal/infrastructure/configuration"

	codes "google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	UnimplementedGoldenserviceServer
	address    string
	repository types.GoldenRepository
	config     configuration.GoldenConfiguration
}

func NewServer(address string, repository types.GoldenRepository, config configuration.GoldenConfiguration) *Server {
	apiGRPC := &Server{
		address:    address,
		repository: repository,
		config:     config,
	}

	return apiGRPC
}

func (s *Server) Repository() types.GoldenRepository {
	return s.repository
}

func (s *Server) GetGRPCCode(err error) codes.Code {
	var code = codes.Internal

	switch {
	case errors.Is(err, shared.ErrGoldenNotFound):
		code = codes.NotFound
	case errors.Is(err, shared.ErrInvalidGoldenId),
		errors.Is(err, shared.ErrInvalidGoldenName),
		errors.Is(err, shared.ErrInvalidPageNumber),
		errors.Is(err, shared.ErrInvalidPageSize),
		strings.Contains(err.Error(), "invalid"),
		strings.Contains(err.Error(), "Invalid"),
		strings.Contains(err.Error(), "illegal"),
		strings.Contains(err.Error(), "bad request"):
		code = codes.InvalidArgument
	}

	return code
}

func (s *Server) ToProtos(domainGoldens []*types.Golden) []*Golden {
	var protoGoldens []*Golden
	for _, golden := range domainGoldens {
		protoGoldens = append(protoGoldens, s.ToProto(golden))
	}

	return protoGoldens
}

func (s *Server) ToProto(domainGolden *types.Golden) *Golden {
	return &Golden{
		Id:        domainGolden.Id,
		Name:      domainGolden.Name,
		Content:   domainGolden.Content,
		CreatedAt: timestamppb.New(domainGolden.CreatedAt),
		UpdatedAt: timestamppb.New(domainGolden.UpdatedAt),
	}
}
