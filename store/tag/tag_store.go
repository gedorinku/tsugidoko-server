package tagstore

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"

	"github.com/gedorinku/tsugidoko-server/infra/record"
	"github.com/gedorinku/tsugidoko-server/store"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type tagStoreImpl struct {
	ctx context.Context
	db  *sql.DB
}

// NewTagStore creates new tag store
func NewTagStore(ctx context.Context, db *sql.DB) store.TagStore {
	return &tagStoreImpl{
		ctx: ctx,
		db:  db,
	}
}

func (s *tagStoreImpl) ListValidTags() ([]*record.Tag, error) {
	m := []qm.QueryMod{
		qm.Load(record.UserTagRels.Tag),
		qm.GroupBy(record.UserTagColumns.TagID),
		qm.Select(record.UserTagColumns.TagID),
	}
	uts, err := record.UserTags(m...).All(s.ctx, s.db)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	tags := make([]*record.Tag, 0, len(uts))
	for _, ct := range uts {
		tags = append(tags, ct.R.Tag)
	}

	return tags, nil
}

func (s *tagStoreImpl) GetTag(tagID int64) (*record.Tag, error) {
	q := qm.Where("id = ?", tagID)
	t, err := record.Tags(q).One(s.ctx, s.db)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return t, nil
}

func (s *tagStoreImpl) CreateTag(tag *record.Tag) (*record.Tag, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	mods := []qm.QueryMod{
		qm.Where("name = ?", tag.Name),
	}
	et, err := record.Tags(mods...).One(s.ctx, tx)
	if err != nil {
		if err != sql.ErrNoRows {
			tx.Rollback()
			return nil, errors.WithStack(err)
		}

		err = tag.Insert(s.ctx, tx, boil.Infer())
		if err != nil {
			tx.Rollback()
			return nil, errors.WithStack(err)
		}
	} else {
		tag = et
	}

	err = tx.Commit()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return tag, nil
}
