package classroomstore

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/gedorinku/tsugidoko-server/app/conv"
	"github.com/gedorinku/tsugidoko-server/app/util"
	"github.com/gedorinku/tsugidoko-server/infra/record"
	"github.com/gedorinku/tsugidoko-server/store"
)

const (
	// MinTagCount represents min number of tag count
	MinTagCount = 2
	// TagCountThresholdRate represents rate of tag count threshold
	TagCountThresholdRate = 0.2
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

type countRes struct {
	TagID int64 `boil:"tag_id"`
	Total int   `boil:"total"`
}

func (s *classRoomStoreImpl) ListClassRoom(tagIDs []int64) ([]*record.ClassRoom, error) {
	mods := s.loadQueries()
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

	th, err := s.thresholds(s.db)
	if err != nil {
		return nil, err
	}
	s.filterClassRooms(res, th)

	return res, nil
}

func (s *classRoomStoreImpl) GetClassRoom(classRoomID int64, tagIDs []int64) (*record.ClassRoom, error) {
	mods := s.loadQueries()
	mods = append(mods, qm.Where("id = ?", classRoomID))
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

func (s *classRoomStoreImpl) loadQueries() []qm.QueryMod {
	return []qm.QueryMod{
		qm.Load(record.ClassRoomRels.Beacons),
		qm.Load(record.ClassRoomRels.Building),
	}
}

func (s *classRoomStoreImpl) thresholds(exec boil.ContextExecutor) (map[int64]int, error) {
	totals := make([]countRes, 0, 0)
	mods := []qm.QueryMod{
		qm.Select("tag_id, COUNT(*) as total"),
		qm.GroupBy("tag_id"),
	}
	err := record.UserTags(mods...).Bind(s.ctx, exec, &totals)
	if err != nil && err != sql.ErrNoRows {
		return nil, errors.WithStack(err)
	}

	res := make(map[int64]int, 0)
	for _, t := range totals {
		res[t.TagID] = util.MaxInt(MinTagCount, int(float64(t.Total)*TagCountThresholdRate))
	}

	return res, nil
}

func (s *classRoomStoreImpl) filterClassRooms(rooms []*record.ClassRoom, thresholds map[int64]int) {
	for _, r := range rooms {
		s.filterClassRoom(r, thresholds)
	}
}

func (s *classRoomStoreImpl) filterClassRoom(room *record.ClassRoom, thresholds map[int64]int) {
	if room.R == nil {
		return
	}

	for _, t := range room.R.ClassRoomTags {
		if thresholds[t.TagID] <= t.Count {
			continue
		}
		t.Count = 0
	}
}
