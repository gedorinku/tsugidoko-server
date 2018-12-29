package server

import (
	"context"

	"github.com/izumin5210/grapi/pkg/grapiserver"

	api_pb "github.com/gedorinku/tsugidoko-server/api"
)

// UserPositionServiceServer is a composite interface of api_pb.UserPositionServiceServer and grapiserver.Server.
type UserPositionServiceServer interface {
	api_pb.UserPositionServiceServer
	grapiserver.Server
}

// NewUserPositionServiceServer creates a new UserPositionServiceServer instance.
func NewUserPositionServiceServer() UserPositionServiceServer {
	return &userPositionServiceServerImpl{}
}

type userPositionServiceServerImpl struct {
}

func (s *userPositionServiceServerImpl) GetUserPosition(context.Context, *api_pb.GetUserPositionRequest) (*api_pb.UserPosition, error) {
	panic("not implemented")
}

func (s *userPositionServiceServerImpl) UpdateUserPosition(context.Context, *api_pb.UpdateUserPositionRequest) (*api_pb.UserPosition, error) {
	panic("not implemented")
}
