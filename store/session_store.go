package store

import (
	"github.com/pkg/errors"

	"github.com/gedorinku/tsugidoko-server/infra/record"
)

var (
	// ErrInvalidUserNameOrPassword is returned when username or password is invalid
	ErrInvalidUserNameOrPassword = errors.New("invalid username or password")
)

// SessionStore provides user data
type SessionStore interface {
	GetSession(secretKey string) (*record.Session, error)
	CreateSession(name, password string) (*record.Session, error)
	DeleteSession(session *record.Session) error
}
