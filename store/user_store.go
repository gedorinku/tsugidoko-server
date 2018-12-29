package store

import (
	"github.com/gedorinku/tsugidoko-server/infra/record"
	"github.com/gedorinku/tsugidoko-server/model"
)

// UserStore provides user data
type UserStore interface {
	ListUsers() ([]*record.User, error)
	GetUser(userID model.UserID) (*record.User, error)
	CreateUser(name, password string) (*record.User, error)
	UpdateUser(user *record.User) error
}
