package wordcloudmodel

import (
	"time"

	"justasking/GO/core/model/boxes/basebox"

	uuid "github.com/satori/go.uuid"
)

// WordCloud is a wordcloud
type WordCloud struct {
	BoxId       uuid.UUID
	Header      string
	DefaultWord string
	CreatedAt   time.Time
	CreatedBy   string
	UpdatedAt   *time.Time
	UpdatedBy   string
	DeletedAt   *time.Time
	BaseBox     baseboxmodel.BaseBox `gorm:"-"`
}

// TableName returns the table name for use with ORM
func (WordCloud) TableName() string {
	return "word_cloud_box"
}
