package server

import (
	"context"

	"github.com/izumin5210/grapi/pkg/grapiserver"
	"github.com/pkg/errors"

	api_pb "github.com/gedorinku/tsugidoko-server/api"
	"github.com/gedorinku/tsugidoko-server/app/conv"
	"github.com/gedorinku/tsugidoko-server/app/di"
	"github.com/gedorinku/tsugidoko-server/app/interceptor"
	"github.com/gedorinku/tsugidoko-server/model"
)

// UserTagServiceServer is a composite interface of api_pb.UserTagServiceServer and grapiserver.Server.
type UserTagServiceServer interface {
	api_pb.UserTagServiceServer
	grapiserver.Server
}

// NewUserTagServiceServer creates a new UserTagServiceServer instance.
func NewUserTagServiceServer(store di.StoreComponent) UserTagServiceServer {
	return &userTagServiceServerImpl{
		StoreComponent: store,
	}

}

type userTagServiceServerImpl struct {
	di.StoreComponent
}

func (s *userTagServiceServerImpl) UpdateUserTag(ctx context.Context, req *api_pb.UpdateUserTagRequest) (*api_pb.UpdateUserTagResponse, error) {
	session := interceptor.GetCurrentSession(ctx)
	if session == nil {
		return nil, ErrInvalidSession
	}

	us := s.UserTagStore(ctx)
	_, err := us.UpdateUserTag(model.UserID(session.UserID), conv.Int32SliceToInt64Slice(req.GetTagIds()))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	resp := &api_pb.UpdateUserTagResponse{
		// TODO
		// Tags: tagsToResponse(tags),
	}
	return resp, nil
}
