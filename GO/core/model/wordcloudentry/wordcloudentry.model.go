package wordcloudentrymodel

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// WordCloudEntry is an entry to a word cloud
type WordCloudEntry struct {
	BoxId      uuid.UUID
	EntryValue string
	IsHidden   bool
	CreatedAt  time.Time
	CreatedBy  int
	UpdatedAt  *time.Time
	UpdatedBy  int
	DeletedAt  *time.Time
}

// TableName returns the table name for use with ORM
func (WordCloudEntry) TableName() string {
	return "word_cloud_box_entries"
}
