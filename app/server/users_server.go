package server

import (
	"context"

	"github.com/izumin5210/grapi/pkg/grapiserver"

	api_pb "github.com/gedorinku/tsugidoko-server/api"
	"github.com/gedorinku/tsugidoko-server/app/di"
)

// UserServiceServer is a composite interface of api_pb.UserServiceServer and grapiserver.Server.
type UserServiceServer interface {
	api_pb.UserServiceServer
	grapiserver.Server
}

// NewUserServiceServer creates a new UserServiceServer instance.
func NewUserServiceServer(store di.StoreComponent) UserServiceServer {
	return &userServiceServerImpl{
		store,
	}
}

type userServiceServerImpl struct {
	di.StoreComponent
}

func (s *userServiceServerImpl) GetCurrentUser(context.Context, *api_pb.GetCurrentUserRequest) (*api_pb.User, error) {
	panic("not implemented")
}

func (s *userServiceServerImpl) UpdateUser(context.Context, *api_pb.UpdateUserRequest) (*api_pb.User, error) {
	panic("not implemented")
}

func (s *userServiceServerImpl) GetUser(context.Context, *api_pb.GetUserRequest) (*api_pb.User, error) {
	panic("not implemented")
}

func (s *userServiceServerImpl) CreateUser(context.Context, *api_pb.CreateUserRequest) (*api_pb.User, error) {
	panic("not implemented")
}
