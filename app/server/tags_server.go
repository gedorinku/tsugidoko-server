package server

import (
	"context"
	"database/sql"

	"google.golang.org/grpc/grpclog"

	"github.com/izumin5210/grapi/pkg/grapiserver"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "github.com/gedorinku/tsugidoko-server/api"
	"github.com/gedorinku/tsugidoko-server/app/di"
	"github.com/gedorinku/tsugidoko-server/infra/record"
)

// TagServiceServer is a composite interface of api_pb.TagServiceServer and grapiserver.Server.
type TagServiceServer interface {
	api_pb.TagServiceServer
	grapiserver.Server
}

// NewTagServiceServer creates a new TagServiceServer instance.
func NewTagServiceServer(store di.StoreComponent) TagServiceServer {
	return &tagServiceServerImpl{
		StoreComponent: store,
	}
}

type tagServiceServerImpl struct {
	di.StoreComponent
}

func (s *tagServiceServerImpl) ListTags(ctx context.Context, req *api_pb.ListTagsRequest) (*api_pb.ListTagsResponse, error) {
	ts := s.TagStore(ctx)
	tags, err := ts.ListValidTags()
	if err != nil && errors.Cause(err) != sql.ErrNoRows {
		grpclog.Error(err)
		return nil, err
	}

	resp := &api_pb.ListTagsResponse{
		Tags: tagsToResponse(tags),
	}
	return resp, nil
}

func (s *tagServiceServerImpl) GetTag(ctx context.Context, req *api_pb.GetTagRequest) (*api_pb.Tag, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *tagServiceServerImpl) CreateTag(ctx context.Context, req *api_pb.CreateTagRequest) (*api_pb.Tag, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func tagsToResponse(tags []*record.Tag) []*api_pb.Tag {
	res := make([]*api_pb.Tag, 0, len(tags))
	for _, t := range tags {
		res = append(res, tagToResponse(t))
	}
	return res
}

func tagToResponse(tag *record.Tag) *api_pb.Tag {
	return &api_pb.Tag{
		Id:   int32(tag.ID),
		Name: tag.Name,
	}
}
