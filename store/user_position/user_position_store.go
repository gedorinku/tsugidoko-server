package userpositionstore

import (
	"context"
	"database/sql"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"google.golang.org/grpc/grpclog"

	"github.com/gedorinku/tsugidoko-server/app/conv"
	"github.com/gedorinku/tsugidoko-server/infra/record"
	"github.com/gedorinku/tsugidoko-server/model"
	"github.com/gedorinku/tsugidoko-server/store"
)

type userPositionStoreImpl struct {
	ctx context.Context
	db  *sql.DB
}

// NewUserPositionStore creates new user position store
func NewUserPositionStore(ctx context.Context, db *sql.DB) store.UserPositionStore {
	return &userPositionStoreImpl{
		ctx: ctx,
		db:  db,
	}
}

func (s *userPositionStoreImpl) GetUserPosition(userID model.UserID) (*record.UserPosition, error) {
	mods := []qm.QueryMod{
		qm.Load(record.UserPositionRels.ClassRoom),
		qm.Where("user_id = ?", userID),
	}
	res, err := record.UserPositions(mods...).One(s.ctx, s.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, errors.WithStack(err)
	}

	return res, nil
}

func (s *userPositionStoreImpl) UpdateUserPosition(userID model.UserID, bssid string, connected bool) (*record.UserPosition, error) {
	tx, err := s.db.Begin()
	if err != nil {
		tx.Rollback()
		return nil, errors.Wrap(err, "failed to begin tx")
	}

	tagIDs, err := s.userTagIDs(tx, userID)
	if err != nil {
		tx.Rollback()
		return nil, errors.WithStack(err)
	}

	up, err := record.UserPositions(qm.Where("user_id = ?", userID)).One(s.ctx, tx)
	if err != nil && err != sql.ErrNoRows {
		tx.Rollback()
		return nil, errors.WithStack(err)
	}
	if up != nil {
		err := s.updateClassRoomTags(tx, up.ClassRoomID, tagIDs, -1)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		_, err = up.Delete(s.ctx, tx)
		if err != nil {
			tx.Rollback()
			return nil, errors.WithStack(err)
		}
	}

	if !connected {
		return nil, errors.WithStack(tx.Commit())
	}

	b, err := record.Beacons(qm.Where("bssid = ?", bssid)).One(s.ctx, tx)
	if err != nil {
		if err == sql.ErrNoRows {
			tx.Rollback()
			return nil, errors.Wrap(err, "beacon not found")
		}

		return nil, errors.WithStack(err)
	}

	newPos := &record.UserPosition{
		UserID:      int64(userID),
		ClassRoomID: b.ClassRoomID,
	}
	err = newPos.Insert(s.ctx, tx, boil.Infer())
	if err != nil {
		tx.Rollback()
		return nil, errors.WithStack(err)
	}
	// newPos.R.ClassRoomに入れる方法がなくてつらい
	newPos, err = record.UserPositions(qm.Load(record.UserPositionRels.ClassRoom), qm.Where("id = ?", newPos.ID)).One(s.ctx, tx)
	if err != nil {
		tx.Rollback()
		return nil, errors.WithStack(err)
	}

	err = s.updateClassRoomTags(tx, b.ClassRoomID, tagIDs, 1)
	if err != nil {
		tx.Rollback()
		return nil, errors.WithStack(err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return newPos, nil
}

type userPositionRes struct {
	ID          int64 `boil:"id"`
	ClassRoomID int64 `boil:"class_room_id"`
	TagID       int64 `boil:"tag_id"`
}

type classRoomTag struct {
	ClassRoomID int64
	TagID       int64
}

func (s *userPositionStoreImpl) Expire(lifeTime time.Duration) (int, error) {
	old := time.Now().Add(-lifeTime)
	grpclog.Infof("%v %v\n", time.Now(), old)
	tx, err := s.db.Begin()
	defer func() {
		if err := recover(); err != nil {
			grpclog.Error(err)
			tx.Rollback()
		}
	}()

	mods := []qm.QueryMod{
		qm.Select("user_positions.id as id, user_positions.class_room_id as class_room_id, ut.tag_id"),
		qm.InnerJoin("user_tags as ut on user_positions.user_id = ut.user_id"),
		qm.Where("user_positions.updated_at <= ?", old.In(boil.GetLocation())),
	}

	var ups []*userPositionRes
	err = record.UserPositions(mods...).Bind(s.ctx, tx, &ups)
	if err != nil {
		tx.Rollback()
		return 0, errors.WithStack(err)
	}

	diff := make(map[classRoomTag]int)
	for _, up := range ups {
		ct := classRoomTag{
			ClassRoomID: up.ClassRoomID,
			TagID:       up.TagID,
		}
		diff[ct]++
	}

	for ct, d := range diff {
		q := `update class_room_tags
			set count = count - $1
			where class_room_id = $2 and tag_id = $3`
		_, err := tx.ExecContext(s.ctx, q, d, ct.ClassRoomID, ct.TagID)
		if err != nil {
			tx.Rollback()
			return 0, errors.WithStack(err)
		}
	}

	if 0 < len(ups) {
		ids := make([]int64, 0, len(ups))
		for _, up := range ups {
			ids = append(ids, up.ID)
		}
		_, err = record.UserPositions(qm.WhereIn("id in ?", conv.Int64SliceToAbstractSlice(ids)...)).DeleteAll(s.ctx, tx)
		if err != nil {
			tx.Rollback()
			return 0, errors.WithStack(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return 0, errors.WithStack(err)
	}

	return len(ups), nil
}

func (s *userPositionStoreImpl) userTagIDs(exec boil.ContextExecutor, userID model.UserID) ([]int64, error) {
	mods := []qm.QueryMod{
		qm.Select(record.UserTagColumns.TagID),
		qm.Where("user_id = ?", userID),
	}
	tags, err := record.UserTags(mods...).All(s.ctx, exec)
	if err != nil && err != sql.ErrNoRows {
		return nil, errors.WithStack(err)
	}

	ids := make([]int64, 0, len(tags))
	for _, t := range tags {
		ids = append(ids, t.TagID)
	}

	return ids, nil
}

func (s *userPositionStoreImpl) updateClassRoomTags(exec boil.ContextExecutor, classRoomID int64, tagIDs []int64, diff int64) error {
	if len(tagIDs) == 0 {
		return nil
	}

	builder := strings.Builder{}
	builder.Grow(len(",$10") * len(tagIDs))
	for i := 5; i <= 3+len(tagIDs); i++ {
		builder.WriteString(",$" + strconv.Itoa(i))
	}
	q := "UPDATE class_room_tags SET count = count + $1, updated_at = $2 WHERE class_room_id = $3 AND tag_id in ($4" + builder.String() + ")"
	args := make([]interface{}, 0, len(tagIDs)+3)
	args = append(args, diff, time.Now().In(boil.GetLocation()), classRoomID)
	for _, id := range tagIDs {
		args = append(args, id)
	}
	_, err := queries.Raw(q, args...).ExecContext(s.ctx, exec)
	if err != nil && err != sql.ErrNoRows {
		return errors.WithStack(err)
	}

	return nil
}

func (s *userPositionStoreImpl) ResetUserPosition() error {
	_, err := record.ClassRoomTags().DeleteAll(s.ctx, s.db)
	if err != nil && err != sql.ErrNoRows {
		return errors.WithStack(err)
	}

	_, err = record.UserPositions().DeleteAll(s.ctx, s.db)
	if err != nil && err != sql.ErrNoRows {
		return errors.WithStack(err)
	}

	rooms, err := record.ClassRooms(qm.Select("id")).All(s.ctx, s.db)
	if err != nil && err != sql.ErrNoRows {
		return errors.WithStack(err)
	}
	tags, err := record.Tags(qm.Select("id")).All(s.ctx, s.db)
	if err != nil && err != sql.ErrNoRows {
		return errors.WithStack(err)
	}

	for _, r := range rooms {
		for _, t := range tags {
			rt := &record.ClassRoomTag{
				ClassRoomID: r.ID,
				TagID:       t.ID,
			}
			err = rt.Insert(s.ctx, s.db, boil.Infer())
			if err != nil {
				return errors.WithStack(err)
			}
		}
	}

	return nil
}
