package votesboxquestionanswermodel

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// VotesBoxQuestionAnswer is a votesboxquestionanswer
type VotesBoxQuestionAnswer struct {
	AnswerId   uuid.UUID
	QuestionId uuid.UUID
	Answer     string
	SortOrder  int
	Votes      int `gorm:"-"`
	CreatedAt  time.Time
	CreatedBy  string
	UpdatedAt  *time.Time
	UpdatedBy  string
	DeletedAt  *time.Time
}

// TableName returns the table name for use with ORM
func (VotesBoxQuestionAnswer) TableName() string {
	return "votes_box_question_answers"
}
