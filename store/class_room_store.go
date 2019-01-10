package store

import (
	"github.com/gedorinku/tsugidoko-server/infra/record"
)

// ClassRoomStore provides class room data
type ClassRoomStore interface {
	ListClassRoom(tagIDs []int64) ([]*record.ClassRoom, error)
	GetClassRoom(classRoomID int64, tagIDs []int64) (*record.ClassRoom, error)
}
