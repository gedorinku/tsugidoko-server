package userstore

import (
	"context"
	"database/sql"

	"github.com/gedorinku/tsugidoko-server/infra/record"
	"github.com/gedorinku/tsugidoko-server/model"
	"github.com/gedorinku/tsugidoko-server/store"
)

type userStoreImpl struct {
	ctx context.Context
	db  *sql.DB
}

// NewUserStore creates new user store
func NewUserStore(ctx context.Context, db *sql.DB) store.UserStore {
	return &userStoreImpl{
		ctx: ctx,
		db:  db,
	}
}

func (s *userStoreImpl) ListUsers() ([]record.User, error) {
	panic("not implemented")
}

func (s *userStoreImpl) GetUser(userID model.UserID) (*record.User, error) {
	panic("not implemented")
}

func (s *userStoreImpl) CreateUser(user *record.User) error {
	panic("not implemented")
}

func (s *userStoreImpl) UpdateUser(user *record.User) error {
	panic("not implemented")
}
