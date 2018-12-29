package sessionstore

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base32"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"golang.org/x/crypto/bcrypt"

	"github.com/gedorinku/tsugidoko-server/infra/record"
	"github.com/gedorinku/tsugidoko-server/store"
)

const (
	// SecretKeyBytes reoresents bytes of secret key
	SecretKeyBytes = 16
)

type sessionStoreImpl struct {
	ctx context.Context
	db  *sql.DB
}

// NewSessionStore creates new session store
func NewSessionStore(ctx context.Context, db *sql.DB) store.SessionStore {
	return &sessionStoreImpl{
		ctx: ctx,
		db:  db,
	}
}

func (s *sessionStoreImpl) GetSession(secretKey string) (*record.Session, error) {
	session, err := record.Sessions(qm.Where("secret_key = ?", secretKey)).One(s.ctx, s.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, errors.WithStack(err)
	}

	return session, nil
}

func (s *sessionStoreImpl) CreateSession(name, password string) (*record.Session, error) {
	u, err := record.Users(qm.Where("name = ?", name)).One(s.ctx, s.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrInvalidUserNameOrPassword
		}
		return nil, errors.WithStack(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordDigest), []byte(password))
	if err != nil {
		return nil, store.ErrInvalidUserNameOrPassword
	}

	key, err := generateSecretKey()
	if err != nil {
		return nil, errors.Wrap(err, "failed to create session")
	}

	session := &record.Session{
		SecretKey: key,
		UserID:    u.ID,
	}
	err = session.Insert(s.ctx, s.db, boil.Infer())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return session, nil
}

func (s *sessionStoreImpl) DeleteSession(session *record.Session) error {
	_, err := session.Delete(s.ctx, s.db)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func generateSecretKey() (string, error) {
	key := make([]byte, SecretKeyBytes)
	_, err := rand.Read(key)
	if err != nil {
		return "", errors.WithStack(err)
	}

	enc := base32.StdEncoding.WithPadding(base32.NoPadding)
	return enc.EncodeToString(key), nil
}
