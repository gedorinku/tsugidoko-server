package store

import "github.com/gedorinku/tsugidoko-server/model"

// ClassRoomStore provides class room data
type ClassRoomStore interface {
	ListClassRoom(tagIDs []int64) ([]*model.ClassRoom, error)
	GetClassRoom(classRoomID int64, tagIDs []int64) (*model.ClassRoom, error)
}
