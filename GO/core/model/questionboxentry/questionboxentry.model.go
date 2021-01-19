package questionboxentrymodel

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// QuestionBoxEntry is a response to a question box
type QuestionBoxEntry struct {
	EntryId    uuid.UUID
	BoxId      uuid.UUID
	Question   string
	IsHidden   bool
	IsFavorite bool
	Upvotes    int
	Downvotes  int
	CreatedAt  time.Time
	CreatedBy  string
	UpdatedAt  *time.Time
	UpdatedBy  string
	DeletedAt  *time.Time
}

// TableName returns the table name for use with ORM
func (QuestionBoxEntry) TableName() string {
	return "question_box_entries"
}
