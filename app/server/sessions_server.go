package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"

	api_pb "github.com/gedorinku/tsugidoko-server/api"
	"github.com/gedorinku/tsugidoko-server/app/di"
	"github.com/gedorinku/tsugidoko-server/app/interceptor"
	"github.com/gedorinku/tsugidoko-server/infra/record"
	"github.com/gedorinku/tsugidoko-server/store"
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

func (s *sessionServiceServerImpl) DeleteSession(ctx context.Context, req *api_pb.DeleteSessionRequest) (*empty.Empty, error) {
	session, ok := interceptor.GetCurrentSession(ctx)
	if !ok {
		return nil, ErrInvalidSession
	}

	ss := s.SessionStore(ctx)
	err := ss.DeleteSession(session)
	if err != nil {
		grpclog.Error(err)
		return nil, err
	}

	return &empty.Empty{}, nil
}

func sessionToResponse(session *record.Session) *api_pb.Session {
	return &api_pb.Session{
		SessionId: session.SecretKey,
		UesrId:    uint32(session.UserID),
	}
}
