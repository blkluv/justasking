package questionboxentryvotemodel

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// QuestionBoxEntryVote records a vote for a question box entry
type QuestionBoxEntryVote struct {
	EntryId   uuid.UUID
	VoteType  string
	VoteValue int
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt *time.Time
	UpdatedBy string
	DeletedAt *time.Time
}

// TableName returns the table name for use with ORM
func (QuestionBoxEntryVote) TableName() string {
	return "question_box_entries_votes"
}
