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

func (s *userTagStoreImpl) UpdateUserTag(userID model.UserID, tagIDs []int64) ([]*record.Tag, error) {
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

	var tags []*record.Tag
	if 0 < len(tagIDs) {
		tags, err = record.Tags(qm.WhereIn("id in ?", toInterfaceSlice(tagIDs)...)).All(s.ctx, tx)
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
