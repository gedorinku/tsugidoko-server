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
		qm.Select("tags.*, count(ut.*) as total"),
		qm.InnerJoin("user_tags as ut on tags.id = ut.tag_id"),
		qm.GroupBy("tags.id"),
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

func (s *tagStoreImpl) GetTag(tagID int64) (*model.Tag, error) {
	mods := []qm.QueryMod{
		qm.Select("tags.*, count(ut.*) as total"),
		qm.InnerJoin("user_tags as ut on tags.id = ut.tag_id and tags.id = ?", tagID),
		qm.GroupBy("tags.id"),
		qm.Limit(1),
	}
	t := &model.Tag{}
	err := record.Tags(mods...).Bind(s.ctx, s.db, t)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return t, nil
}

func (s *tagStoreImpl) CreateTag(name string) (*model.Tag, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	tag := &record.Tag{
		Name: name,
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

	total, err := s.tagTotalUsers(tag.ID, tx)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	res := &model.Tag{
		ID:        tag.ID,
		Name:      tag.Name,
		CreatedAt: tag.CreatedAt,
		UpdatedAt: tag.UpdatedAt,
		Total:     int(total),
	}
	return res, nil
}

func (s *tagStoreImpl) tagTotalUsers(tagID int64, exec boil.ContextExecutor) (int64, error) {
	mods := []qm.QueryMod{
		qm.Where("tag_id = ?", tagID),
		qm.GroupBy("tag_id"),
	}
	cnt, err := record.UserTags(mods...).Count(s.ctx, exec)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	return cnt, nil
}
