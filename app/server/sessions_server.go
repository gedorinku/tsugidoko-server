package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/izumin5210/grapi/pkg/grapiserver"

	api_pb "github.com/gedorinku/tsugidoko-server/api"
)

// SessionServiceServer is a composite interface of api_pb.SessionServiceServer and grapiserver.Server.
type SessionServiceServer interface {
	api_pb.SessionServiceServer
	grapiserver.Server
}

// NewSessionServiceServer creates a new SessionServiceServer instance.
func NewSessionServiceServer() SessionServiceServer {
	return &sessionServiceServerImpl{}
}

type sessionServiceServerImpl struct {
}

func (s *sessionServiceServerImpl) GetSession(context.Context, *api_pb.GetSessionRequest) (*api_pb.Session, error) {
	panic("not implemented")
}

func (s *sessionServiceServerImpl) CreateSession(context.Context, *api_pb.CreateSessionRequest) (*api_pb.Session, error) {
	panic("not implemented")
}

func (s *sessionServiceServerImpl) DeleteSession(context.Context, *api_pb.DeleteSessionRequest) (*empty.Empty, error) {
	panic("not implemented")
}
