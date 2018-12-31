package userpositionstore

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"

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

func (s *userPositionStoreImpl) userTagIDs(exec boil.ContextExecutor, userID model.UserID) ([]int64, error) {
	mods := []qm.QueryMod{
		qm.Select(record.UserTagColumns.UserID),
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

func (s *userPositionStoreImpl) updateClassRoomTags(tx *sql.Tx, classRoomID int64, tagIDs []int64, diff int64) error {
	if len(tagIDs) == 0 {
		return nil
	}

	q := "UPDATE class_room_tags SET count = count + ? " +
		"WHERE class_room_id = ? AND tag_id = ANY(?)"
	stmt, err := tx.PrepareContext(s.ctx, q)
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = stmt.ExecContext(s.ctx, q, diff, classRoomID, pq.Array(tagIDs))
	if err != nil && err != sql.ErrNoRows {
		return errors.WithStack(err)
	}

	return nil
}
