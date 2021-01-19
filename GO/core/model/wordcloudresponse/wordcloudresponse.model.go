package wordcloudresponsemodel

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// WordCloudResponse is a response to a word cloud
type WordCloudResponse struct {
	BoxId     uuid.UUID
	Response  string
	IsHidden  bool
	Count     int `gorm:"-"`
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt *time.Time
	UpdatedBy string
	DeletedAt *time.Time
}

// TableName returns the table name for use with ORM
func (WordCloudResponse) TableName() string {
	return "word_cloud_box_responses"
}
