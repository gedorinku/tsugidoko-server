package server

import (
	api_pb "github.com/gedorinku/tsugidoko-server/api"
	"github.com/gedorinku/tsugidoko-server/infra/record"
)

func tagToResponse(tag *record.Tag) *api_pb.Tag {
	return &api_pb.Tag{
		Id:   int32(tag.ID),
		Name: tag.Name,
	}
}
