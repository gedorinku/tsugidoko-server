package di

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq" // for database/sql
	"github.com/pkg/errors"

	"github.com/gedorinku/tsugidoko-server/app/config"
	"github.com/gedorinku/tsugidoko-server/store"
	sessionstore "github.com/gedorinku/tsugidoko-server/store/session"
	userstore "github.com/gedorinku/tsugidoko-server/store/user"
)

// StoreComponent is an interface of stores
type StoreComponent interface {
	UserStore(ctx context.Context) store.UserStore
	SessionStore(ctx context.Context) store.SessionStore
}

// NewStoreComponent creates new store component
func NewStoreComponent(cfg *config.Config) (StoreComponent, error) {
	db, err := connectRDB(cfg)
	if err != nil {
		return nil, err
	}

	return &storeComponentImpl{
		db: db,
	}, nil
}

type storeComponentImpl struct {
	db *sql.DB
}

func (s *storeComponentImpl) UserStore(ctx context.Context) store.UserStore {
	return userstore.NewUserStore(ctx, s.db)
}

func (s *storeComponentImpl) SessionStore(ctx context.Context) store.SessionStore {
	return sessionstore.NewSessionStore(ctx, s.db)
}

func connectRDB(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.DataBaseURL)
	if err != nil {
		return nil, errors.Wrap(err, "faild to connect db")
	}
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, errors.Wrap(err, "faild to ping db")
	}

	return db, nil
}
