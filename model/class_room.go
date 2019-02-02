package model

import (
	"time"

	"github.com/gedorinku/tsugidoko-server/infra/record"
	"github.com/volatiletech/null"
)

// ClassRoom represents class room model
type ClassRoom struct {
	ID            int64              `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name          string             `boil:"name" json:"name" toml:"name" yaml:"name"`
	Latitude      null.Float64       `boil:"latitude" json:"latitude,omitempty" toml:"latitude" yaml:"latitude,omitempty"`
	Longitude     null.Float64       `boil:"longitude" json:"longitude,omitempty" toml:"longitude" yaml:"longitude,omitempty"`
	CreatedAt     time.Time          `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt     time.Time          `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	BuildingID    null.Int64         `boil:"building_id" json:"building_id,omitempty" toml:"building_id" yaml:"building_id,omitempty"`
	Floor         int                `boil:"floor" json:"floor" toml:"floor" yaml:"floor"`
	LocalX        float64            `boil:"local_x" json:"local_x" toml:"local_x" yaml:"local_x"`
	LocalY        float64            `boil:"local_y" json:"local_y" toml:"local_y" yaml:"local_y"`
	Building      *record.Building   `boil:",bind"`
	Beacons       record.BeaconSlice `boil:",bind"`
	ClassRoomTags []*ClassRoomTag    `boil:"-"`
}
