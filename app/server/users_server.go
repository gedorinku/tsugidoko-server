package server

import (
	"context"
	"regexp"

	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "github.com/gedorinku/tsugidoko-server/api"
	"github.com/gedorinku/tsugidoko-server/app/di"
	"github.com/gedorinku/tsugidoko-server/infra/record"
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

var userNameRegex = regexp.MustCompile(`^[a-zA-Z0-9_\-.]{3,15}$`)

type userServiceServerImpl struct {
	di.StoreComponent
}

func (s *userServiceServerImpl) GetCurrentUser(ctx context.Context, req *api_pb.GetCurrentUserRequest) (*api_pb.User, error) {
	panic("not implemented")
}

func (s *userServiceServerImpl) UpdateUser(ctx context.Context, req *api_pb.UpdateUserRequest) (*api_pb.User, error) {
	panic("not implemented")
}

func (s *userServiceServerImpl) GetUser(ctx context.Context, req *api_pb.GetUserRequest) (*api_pb.User, error) {
	panic("not implemented")
}

func (s *userServiceServerImpl) CreateUser(ctx context.Context, req *api_pb.CreateUserRequest) (*api_pb.User, error) {
	name := req.GetName()
	if !userNameRegex.MatchString(name) {
		return nil, status.Error(codes.InvalidArgument, `Username should be matched "^[a-zA-Z0-9_\-.]{3,15}$"`)
	}

	store := s.UserStore(ctx)
	u, err := store.CreateUser(name, req.GetPassword())
	if err != nil {
		return nil, err
	}
	return userToResponse(u), nil
}

func userToResponse(user *record.User) *api_pb.User {
	// TODO user tags
	return &api_pb.User{
		UserId: uint32(user.ID),
		Name:   user.Name,
	}
}
