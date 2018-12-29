package userstore

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"golang.org/x/crypto/bcrypt"

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

func (s *userStoreImpl) ListUsers() ([]*record.User, error) {
	res, err := record.Users().All(s.ctx, s.db)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return res, nil
}

func (s *userStoreImpl) GetUser(userID model.UserID) (*record.User, error) {
	u, err := record.Users(qm.Where("id = ?", userID)).One(s.ctx, s.db)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return u, nil
}

func (s *userStoreImpl) CreateUser(name, password string) (*record.User, error) {
	d, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	u := &record.User{
		Name:           name,
		PasswordDigest: string(d),
	}
	err = u.Insert(s.ctx, s.db, boil.Whitelist("name"))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return nil, nil
}

func (s *userStoreImpl) UpdateUser(user *record.User) error {
	panic("not implemented")
}
