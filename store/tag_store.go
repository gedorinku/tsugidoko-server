package store

import (
	"github.com/gedorinku/tsugidoko-server/infra/record"
)

// TagStore provides tag data
type TagStore interface {
	ListValidTags() ([]*record.Tag, error)
	GetTag(tagID int64) (*record.Tag, error)
	CreateTag(tag *record.Tag) (*record.Tag, error)
}
