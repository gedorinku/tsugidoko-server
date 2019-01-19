package server

import (
	"context"

	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "github.com/gedorinku/tsugidoko-server/api"
)

// TagServiceServer is a composite interface of api_pb.TagServiceServer and grapiserver.Server.
type TagServiceServer interface {
	api_pb.TagServiceServer
	grapiserver.Server
}

// NewTagServiceServer creates a new TagServiceServer instance.
func NewTagServiceServer() TagServiceServer {
	return &tagServiceServerImpl{}
}

type tagServiceServerImpl struct {
}

func (s *tagServiceServerImpl) ListTags(ctx context.Context, req *api_pb.ListTagsRequest) (*api_pb.ListTagsResponse, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *tagServiceServerImpl) GetTag(ctx context.Context, req *api_pb.GetTagRequest) (*api_pb.Tag, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *tagServiceServerImpl) CreateTag(ctx context.Context, req *api_pb.CreateTagRequest) (*api_pb.Tag, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}
