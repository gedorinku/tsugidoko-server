package userpositionstore

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
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

func (s *userPositionStoreImpl) UpdateUserPosition(userID model.UserID, bssid string) (*record.UserPosition, error) {
	panic("not implemented")
}
