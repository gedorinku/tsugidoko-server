package store

import (
	"github.com/gedorinku/tsugidoko-server/infra/record"
	"github.com/gedorinku/tsugidoko-server/model"
)

// TagStore provides tag data
type TagStore interface {
	ListValidTags() ([]*model.Tag, error)
	GetTag(tagID int64) (*model.Tag, error)
	CreateTag(tag *record.Tag) (*record.Tag, error)
}
