package store

import (
	"github.com/gedorinku/tsugidoko-server/model"
)

// UserTagStore provides user's tag data
type UserTagStore interface {
	UpdateUserTag(userID model.UserID, tagIDs []int64) ([]*model.Tag, error)
	AddUserTag(userID model.UserID, tagID int64) error
}
