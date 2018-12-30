package classroomstore

import (
	"context"
	"database/sql"

	"github.com/gedorinku/tsugidoko-server/infra/record"
	"github.com/gedorinku/tsugidoko-server/store"
)

type classRoomStoreImpl struct {
	ctx context.Context
	db  *sql.DB
}

// NewClassRoomStore creates new class room store
func NewClassRoomStore(ctx context.Context, db *sql.DB) store.ClassRoomStore {
	return &classRoomStoreImpl{
		ctx: ctx,
		db:  db,
	}
}

func (s *classRoomStoreImpl) ListClassRoom() ([]*record.ClassRoom, error) {
	panic("not implemented")
}

func (s *classRoomStoreImpl) GetClassRoom(classRoomID int64) (*record.ClassRoom, error) {
	panic("not implemented")
}
