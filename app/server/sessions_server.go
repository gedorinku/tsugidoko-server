package server

import (
	"context"

	"github.com/gedorinku/tsugidoko-server/infra/record"

	"google.golang.org/grpc/grpclog"

	"github.com/gedorinku/tsugidoko-server/store"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/izumin5210/grapi/pkg/grapiserver"

	api_pb "github.com/gedorinku/tsugidoko-server/api"
	"github.com/gedorinku/tsugidoko-server/app/di"
)

// SessionServiceServer is a composite interface of api_pb.SessionServiceServer and grapiserver.Server.
type SessionServiceServer interface {
	api_pb.SessionServiceServer
	grapiserver.Server
}

// NewSessionServiceServer creates a new SessionServiceServer instance.
func NewSessionServiceServer(store di.StoreComponent) SessionServiceServer {
	return &sessionServiceServerImpl{
		store,
	}
}

type sessionServiceServerImpl struct {
	di.StoreComponent
}

func (s *sessionServiceServerImpl) CreateSession(ctx context.Context, req *api_pb.CreateSessionRequest) (*api_pb.Session, error) {
	ss := s.SessionStore(ctx)
	session, err := ss.CreateSession(req.GetName(), req.GetPassword())
	if err != nil {
		if err == store.ErrInvalidUserNameOrPassword {
			return nil, status.Error(codes.Unauthenticated, "Invalid username or password")
		}
		grpclog.Error(err)
		return nil, err
	}

	return sessionToResponse(session), nil
}

func (s *sessionServiceServerImpl) DeleteSession(context.Context, *api_pb.DeleteSessionRequest) (*empty.Empty, error) {
	panic("not implemented")
}

func sessionToResponse(session *record.Session) *api_pb.Session {
	return &api_pb.Session{
		SessionId: session.SecretKey,
		UesrId:    uint32(session.UserID),
	}
}
