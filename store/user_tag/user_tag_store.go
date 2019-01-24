package usertagstore

import (
	"context"
	"database/sql"
	"strings"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/gedorinku/tsugidoko-server/infra/record"
	"github.com/gedorinku/tsugidoko-server/model"
	"github.com/gedorinku/tsugidoko-server/store"
)

type userTagStoreImpl struct {
	ctx context.Context
	db  *sql.DB
}

// NewUserTagStore creates new user's tag store
func NewUserTagStore(ctx context.Context, db *sql.DB) store.UserTagStore {
	return &userTagStoreImpl{
		ctx: ctx,
		db:  db,
	}
}

func (s *userTagStoreImpl) UpdateUserTag(userID model.UserID, tagIDs []int64) ([]*model.Tag, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	mods := []qm.QueryMod{
		qm.Where("user_id = ?", userID),
	}
	_, err = record.UserTags(mods...).DeleteAll(s.ctx, tx)
	if err != nil {
		tx.Rollback()
		return nil, errors.WithStack(err)
	}

	// saiaku na ko-do
	for _, id := range tagIDs {
		t := &record.UserTag{
			UserID: int64(userID),
			TagID:  id,
		}
		err = t.Insert(s.ctx, tx, boil.Infer())
		if err != nil {
			tx.Rollback()
			return nil, errors.WithStack(err)
		}
	}

	tags := []*model.Tag{}
	if 0 < len(tagIDs) {
		mods := []qm.QueryMod{
			qm.Select("tags.*, count(ut.*) as total"),
			qm.InnerJoin("user_tags as ut on tags.id = ut.tag_id"),
			qm.WhereIn("tags.id in ?", toInterfaceSlice(tagIDs)...),
			qm.GroupBy("tags.id"),
		}
		err = record.Tags(mods...).Bind(s.ctx, tx, &tags)
		if err != nil {
			tx.Rollback()
			return nil, errors.WithStack(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, errors.WithStack(err)
	}
	return tags, nil
}

func (s *userTagStoreImpl) AddUserTag(userID model.UserID, tagID int64) error {
	tx, err := s.db.Begin()
	if err != nil {
		return errors.WithStack(err)
	}

	mods := []qm.QueryMod{
		qm.Where("user_id = ?", userID),
		qm.Where("tag_id = ?", tagID),
	}
	exists, err := record.UserTags(mods...).Exists(s.ctx, tx)
	if err != nil {
		tx.Rollback()
		return errors.WithStack(err)
	}

	if !exists {
		ut := &record.UserTag{
			UserID: int64(userID),
			TagID:  tagID,
		}
		err = ut.Insert(s.ctx, tx, boil.Infer())
		if err != nil {
			tx.Rollback()
			return errors.WithStack(err)
		}
	}

	return errors.WithStack(tx.Commit())
}

func (s *userTagStoreImpl) bulkInsertQuery(tags []*record.Tag) (string, error) {
	b := strings.Builder{}
	b.WriteString(`INSERT INTO user_tags ("user_id", "tag_id", "created_at", "updated_at") VALUES`)

	for i := range tags {
		b.WriteString(`()`)
		if i != len(tags)-1 {
			b.WriteString(",")
		}
	}

	//TODO
	return "", errors.New("not impl")
}

func toInterfaceSlice(s []int64) []interface{} {
	res := make([]interface{}, 0, len(s))
	for _, i := range s {
		res = append(res, i)
	}
	return res
}
