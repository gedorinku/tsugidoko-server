package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "github.com/gedorinku/tsugidoko-server/api"
)

// ClassRoomServiceServer is a composite interface of api_pb.ClassRoomServiceServer and grapiserver.Server.
type ClassRoomServiceServer interface {
	api_pb.ClassRoomServiceServer
	grapiserver.Server
}

// NewClassRoomServiceServer creates a new ClassRoomServiceServer instance.
func NewClassRoomServiceServer() ClassRoomServiceServer {
	return &classRoomServiceServerImpl{}
}

type classRoomServiceServerImpl struct {
}

func (s *classRoomServiceServerImpl) ListClassRooms(ctx context.Context, req *api_pb.ListClassRoomsRequest) (*api_pb.ListClassRoomsResponse, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *classRoomServiceServerImpl) GetClassRoom(ctx context.Context, req *api_pb.GetClassRoomRequest) (*api_pb.ClassRoom, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *classRoomServiceServerImpl) CreateClassRoom(ctx context.Context, req *api_pb.CreateClassRoomRequest) (*api_pb.ClassRoom, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *classRoomServiceServerImpl) UpdateClassRoom(ctx context.Context, req *api_pb.UpdateClassRoomRequest) (*api_pb.ClassRoom, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *classRoomServiceServerImpl) DeleteClassRoom(ctx context.Context, req *api_pb.DeleteClassRoomRequest) (*empty.Empty, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}
