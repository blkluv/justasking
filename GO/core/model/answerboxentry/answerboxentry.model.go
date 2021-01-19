package answerboxentrymodel

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// AnswerBoxEntry is an answerbox question
type AnswerBoxEntry struct {
	EntryId    uuid.UUID
	QuestionId uuid.UUID
	Entry      string
	IsHidden   bool
	CreatedAt  time.Time
	CreatedBy  string
	UpdatedAt  *time.Time
	UpdatedBy  string
	DeletedAt  *time.Time
}

// TableName returns the table name for use with ORM
func (AnswerBoxEntry) TableName() string {
	return "answer_box_entries"
}
