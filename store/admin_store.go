package store

import (
	"github.com/gedorinku/tsugidoko-server/infra/record"
)

// AdminStore provides administrative operations
type AdminStore interface {
	Overwrite(buildings []*record.Building, rooms []*record.ClassRoom, beacons []*record.Beacon) error
}
