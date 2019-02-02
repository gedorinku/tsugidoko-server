package store

import (
	"time"

	"github.com/gedorinku/tsugidoko-server/infra/record"
	"github.com/gedorinku/tsugidoko-server/model"
)

// UserPositionStore provides user data
type UserPositionStore interface {
	GetUserPosition(userID model.UserID) (*record.UserPosition, error)
	UpdateUserPosition(userID model.UserID, bssid string, connected bool) (*record.UserPosition, error)
	ResetUserPosition() error
	Expire(lifeTime time.Duration) (updated int, err error)
}
