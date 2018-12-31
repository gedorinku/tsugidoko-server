package server

import (
	"context"
	"database/sql"

	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc/grpclog"

	api_pb "github.com/gedorinku/tsugidoko-server/api"
	"github.com/gedorinku/tsugidoko-server/app/di"
	"github.com/gedorinku/tsugidoko-server/app/interceptor"
	"github.com/gedorinku/tsugidoko-server/infra/record"
	"github.com/gedorinku/tsugidoko-server/model"
)

// UserPositionServiceServer is a composite interface of api_pb.UserPositionServiceServer and grapiserver.Server.
type UserPositionServiceServer interface {
	api_pb.UserPositionServiceServer
	grapiserver.Server
}

// NewUserPositionServiceServer creates a new UserPositionServiceServer instance.
func NewUserPositionServiceServer(store di.StoreComponent) UserPositionServiceServer {
	return &userPositionServiceServerImpl{
		StoreComponent: store,
	}
}

type userPositionServiceServerImpl struct {
	di.StoreComponent
}

func (s *userPositionServiceServerImpl) GetUserPosition(ctx context.Context, req *api_pb.GetUserPositionRequest) (*api_pb.UserPosition, error) {
	session := interceptor.GetCurrentSession(ctx)
	if session == nil {
		return nil, ErrInvalidSession
	}

	us := s.UserPositionStore(ctx)
	pos, err := us.GetUserPosition(model.UserID(session.ID))
	if err != nil {
		if err == sql.ErrNoRows {
			return userPositionToResponse(nil), nil
		}
		grpclog.Error(err)
		return nil, err
	}

	return userPositionToResponse(pos), nil
}

func (s *userPositionServiceServerImpl) UpdateUserPosition(context.Context, *api_pb.UpdateUserPositionRequest) (*api_pb.UserPosition, error) {
	panic("not implemented")
}

func userPositionToResponse(pos *record.UserPosition) *api_pb.UserPosition {
	var room *api_pb.ClassRoom
	if pos != nil && pos.R.ClassRoom != nil {
		room = classRoomToResponse(pos.R.ClassRoom)
	}
	resp := &api_pb.UserPosition{
		ClassRoom: room,
		IsValid:   room != nil,
	}

	return resp
}
