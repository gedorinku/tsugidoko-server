package model

import (
	"github.com/gedorinku/tsugidoko-server/infra/record"
)

// ClassRoomTag represents class room tag
type ClassRoomTag struct {
	ClassRoomID int64       `boil:"class_room_id" json:"class_room_id" toml:"class_room_id" yaml:"class_room_id"`
	TagID       int64       `boil:"tag_id" json:"tag_id" toml:"tag_id" yaml:"tag_id"`
	Tag         *record.Tag `boil:",bind"`
	Count       int         `boil:"count" json:"count" toml:"count" yaml:"count"`
}
