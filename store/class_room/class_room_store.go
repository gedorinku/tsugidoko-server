package classroomstore

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/gedorinku/tsugidoko-server/app/conv"
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

func (s *classRoomStoreImpl) ListClassRoom(tagIDs []int64) ([]*record.ClassRoom, error) {
	mods := []qm.QueryMod{
		qm.Load(record.ClassRoomRels.Beacons),
	}
	if 0 < len(tagIDs) {
		q := qm.WhereIn("tags.id in ?", conv.Int64SliceToAbstractSlice(tagIDs)...)
		mods = append(mods, qm.Load(record.ClassRoomRels.ClassRoomTags+"."+record.ClassRoomTagRels.Tag, q))
	}
	res, err := record.ClassRooms(mods...).All(s.ctx, s.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.WithStack(err)
	}

	return res, nil
}

func (s *classRoomStoreImpl) GetClassRoom(classRoomID int64, tagIDs []int64) (*record.ClassRoom, error) {
	mods := []qm.QueryMod{
		qm.Load(record.ClassRoomRels.Beacons),
		qm.Where("id = ?", classRoomID),
	}
	if 0 < len(tagIDs) {
		q := qm.WhereIn("tags.id in ?", conv.Int64SliceToAbstractSlice(tagIDs)...)
		mods = append(mods, qm.Load(record.ClassRoomRels.ClassRoomTags+"."+record.ClassRoomTagRels.Tag, q))
	}
	res, err := record.ClassRooms(mods...).One(s.ctx, s.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, errors.WithStack(err)
	}

	return res, nil
}
