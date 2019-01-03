package server

import (
	type_pb "github.com/gedorinku/tsugidoko-server/api/type"
	"github.com/gedorinku/tsugidoko-server/infra/record"
)

func tagToResponse(tag *record.Tag) *type_pb.Tag {
	return &type_pb.Tag{
		Id:   int32(tag.ID),
		Name: tag.Name,
	}
}
