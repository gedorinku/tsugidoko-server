package store

import (
	"github.com/gedorinku/tsugidoko-server/infra/record"
	"github.com/gedorinku/tsugidoko-server/model"
)

// UserPositionStore provides user data
type UserPositionStore interface {
	GetUserPosition(userID model.UserID) (*record.UserPosition, error)
	UpdateUserPosition(userID model.UserID, bssid string) (*record.UserPosition, error)
}
