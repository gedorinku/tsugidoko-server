package tagstore

import (
	"context"
	"database/sql"
	"sort"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"

	"github.com/gedorinku/tsugidoko-server/infra/record"
	"github.com/gedorinku/tsugidoko-server/model"
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

func (s *tagStoreImpl) ListValidTags() ([]*model.Tag, error) {
	tags := []*model.Tag{}
	m := []qm.QueryMod{
		qm.Select("t.*, count(ut.*)"),
		qm.InnerJoin("user_tags as ut on t.id = ut.tag_id"),
		qm.GroupBy("t.id"),
		qm.Having("0 < count(ut.*)"),
	}
	err := record.Tags(m...).Bind(s.ctx, s.db, &tags)
	if err != nil && err != sql.ErrNoRows {
		return nil, errors.WithStack(err)
	}

	sort.Slice(tags, func(i, j int) bool {
		return tags[i].Name < tags[j].Name
	})

	return tags, nil
}

//select t.*, count(ut.*) from tags as t inner join user_tags as ut on t.id = ut.tag_id group by t.id;
func (s *tagStoreImpl) GetTag(tagID int64) (*model.Tag, error) {
	mods := []qm.QueryMod{
		qm.Select("t.*, count(ut.*)"),
		qm.InnerJoin("user_tags as ut on t.id = ut.tag_id and t.id = ?", tagID),
		qm.GroupBy("t.id"),
		qm.Limit(1),
	}
	t := &model.Tag{}
	err := record.Tags(mods...).Bind(s.ctx, s.db, t)
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
