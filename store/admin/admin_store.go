package adminstore

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"google.golang.org/grpc/grpclog"

	"github.com/gedorinku/tsugidoko-server/infra/record"
	"github.com/gedorinku/tsugidoko-server/store"
)

type adminStoreImpl struct {
	ctx context.Context
	db  *sql.DB
}

// NewAdminStore creates new admin store
func NewAdminStore(ctx context.Context, db *sql.DB) store.AdminStore {
	return &adminStoreImpl{
		ctx: ctx,
		db:  db,
	}
}

func (s *adminStoreImpl) Overwrite(buildings []*record.Building, rooms []*record.ClassRoom, beacons []*record.Beacon) error {
	tx, err := s.db.Begin()
	if err != nil {
		return errors.WithStack(err)
	}
	defer func() {
		if err := recover(); err != nil {
			rerr := tx.Rollback()
			grpclog.Error(err)
			if rerr != nil {
				grpclog.Error(rerr)
			}
		}
	}()

	err = s.deleteOldRecords(tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = s.insertNewRecords(tx, buildings, rooms, beacons)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return errors.WithStack(err)
	}
	return nil
}

func (s *adminStoreImpl) deleteOldRecords(tx *sql.Tx) error {
	_, err := record.ClassRoomTags().DeleteAll(s.ctx, tx)
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = record.UserPositions().DeleteAll(s.ctx, tx)
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = record.Beacons().DeleteAll(s.ctx, tx)
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = record.ClassRooms().DeleteAll(s.ctx, tx)
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = record.Buildings().DeleteAll(s.ctx, tx)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s *adminStoreImpl) insertNewRecords(tx *sql.Tx, buildings []*record.Building, rooms []*record.ClassRoom, beacons []*record.Beacon) error {
	for _, b := range buildings {
		err := b.Insert(s.ctx, tx, boil.Infer())
		if err != nil {
			return errors.WithStack(err)
		}
	}
	for _, r := range rooms {
		err := r.Insert(s.ctx, tx, boil.Infer())
		if err != nil {
			return errors.WithStack(err)
		}
	}
	for _, b := range beacons {
		err := b.Insert(s.ctx, tx, boil.Infer())
		if err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}
