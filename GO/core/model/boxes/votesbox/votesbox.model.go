package votesboxmodel

import (
	"time"

	baseboxmodel "github.com/chande/justasking/core/model/boxes/basebox"
	votesboxquestionmodel "github.com/chande/justasking/core/model/votesboxquestion"

	uuid "github.com/satori/go.uuid"
)

// VotesBox is a votesbox
type VotesBox struct {
	BoxId     uuid.UUID
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt *time.Time
	UpdatedBy string
	DeletedAt *time.Time
	Questions []votesboxquestionmodel.VotesBoxQuestion `gorm:"-"`
	BaseBox   baseboxmodel.BaseBox                     `gorm:"-"`
}

// TableName returns the table name for use with ORM
func (VotesBox) TableName() string {
	return "votes_box"
}
