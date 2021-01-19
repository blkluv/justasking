package votesboxquestionanswervotemodel

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// VotesBoxQuestionAnswerVote is a VotesBoxQuestionAnswerVote
type VotesBoxQuestionAnswerVote struct {
	AnswerId        uuid.UUID
	QuestionId      uuid.UUID
	RecordsInserted int
	CreatedAt       time.Time
	CreatedBy       string
	UpdatedAt       *time.Time
	UpdatedBy       string
	DeletedAt       *time.Time
}

// TableName returns the table name for use with ORM
func (VotesBoxQuestionAnswerVote) TableName() string {
	return "votes_box_question_answers_votes"
}
