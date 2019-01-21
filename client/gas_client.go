package client

import (
	"github.com/gedorinku/tsugidoko-server/infra/record"
)

// GASClient provides buildings data
type GASClient interface {
	GetBuildings() (*GetBuildingsResponse, error)
}

// GetBuildingsResponse represent response of GetBuildings
type GetBuildingsResponse struct {
	Buildings  []*record.Building  `json:"buildings"`
	ClassRooms []*record.ClassRoom `json:"class_rooms"`
	Beacons    []*record.Beacon    `json:"beacons"`
}
