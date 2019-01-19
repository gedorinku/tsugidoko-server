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
		qm.Load(record.ClassRoomTagRels.Tag),
		qm.Select(record.ClassRoomTagColumns.ID),
		qm.Where("0 < count"),
	}
	cts, err := record.ClassRoomTags(m...).All(s.ctx, s.db)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	tags := make([]*record.Tag, 0, len(cts))
	for _, ct := range cts {
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

func (s *tagStoreImpl) CreateTag(tag *record.Tag) error {
	err := tag.Insert(s.ctx, s.db, boil.Infer())
	return errors.WithStack(err)
}
