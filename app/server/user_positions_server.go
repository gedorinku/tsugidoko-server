package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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

func (s *userPositionServiceServerImpl) ListUserPositions(ctx context.Context, req *api_pb.ListUserPositionsRequest) (*api_pb.ListUserPositionsResponse, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *userPositionServiceServerImpl) GetUserPosition(ctx context.Context, req *api_pb.GetUserPositionRequest) (*api_pb.UserPosition, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *userPositionServiceServerImpl) CreateUserPosition(ctx context.Context, req *api_pb.CreateUserPositionRequest) (*api_pb.UserPosition, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *userPositionServiceServerImpl) UpdateUserPosition(ctx context.Context, req *api_pb.UpdateUserPositionRequest) (*api_pb.UserPosition, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *userPositionServiceServerImpl) DeleteUserPosition(ctx context.Context, req *api_pb.DeleteUserPositionRequest) (*empty.Empty, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}
