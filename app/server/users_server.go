package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "github.com/gedorinku/tsugidoko-server/api"
)

// UserServiceServer is a composite interface of api_pb.UserServiceServer and grapiserver.Server.
type UserServiceServer interface {
	api_pb.UserServiceServer
	grapiserver.Server
}

// NewUserServiceServer creates a new UserServiceServer instance.
func NewUserServiceServer() UserServiceServer {
	return &userServiceServerImpl{}
}

type userServiceServerImpl struct {
}

func (s *userServiceServerImpl) ListUsers(ctx context.Context, req *api_pb.ListUsersRequest) (*api_pb.ListUsersResponse, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *userServiceServerImpl) GetUser(ctx context.Context, req *api_pb.GetUserRequest) (*api_pb.User, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *userServiceServerImpl) CreateUser(ctx context.Context, req *api_pb.CreateUserRequest) (*api_pb.User, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *userServiceServerImpl) UpdateUser(ctx context.Context, req *api_pb.UpdateUserRequest) (*api_pb.User, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *userServiceServerImpl) DeleteUser(ctx context.Context, req *api_pb.DeleteUserRequest) (*empty.Empty, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}
