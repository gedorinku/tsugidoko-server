package server

import (
	"context"
	"regexp"

	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"

	api_pb "github.com/gedorinku/tsugidoko-server/api"
	type_pb "github.com/gedorinku/tsugidoko-server/api/type"
	"github.com/gedorinku/tsugidoko-server/app/di"
	"github.com/gedorinku/tsugidoko-server/app/interceptor"
	"github.com/gedorinku/tsugidoko-server/infra/record"
	"github.com/gedorinku/tsugidoko-server/model"
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
	session := interceptor.GetCurrentSession(ctx)
	if session == nil {
		return nil, ErrInvalidSession
	}

	us := s.UserStore(ctx)
	u, err := us.GetUser(model.UserID(session.UserID))
	if err != nil {
		grpclog.Error(err)
		return nil, err
	}

	return userToResponse(u), nil
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
		grpclog.Error(err)
		return nil, err
	}
	return userToResponse(u), nil
}

func userToResponse(user *record.User) *api_pb.User {
	return &api_pb.User{
		UserId: uint32(user.ID),
		Name:   user.Name,
		Tags:   userTagsToResponse(user.R.UserTags),
	}
}

func userTagsToResponse(userTags []*record.UserTag) []*type_pb.Tag {
	resp := make([]*type_pb.Tag, 0, len(userTags))
	for _, t := range userTags {
		resp = append(resp, userTagToResponse(t))
	}
	return resp
}

func userTagToResponse(userTag *record.UserTag) *type_pb.Tag {
	t := userTag.R.Tag
	return &type_pb.Tag{
		Id:   int32(t.ID),
		Name: t.Name,
	}
}
