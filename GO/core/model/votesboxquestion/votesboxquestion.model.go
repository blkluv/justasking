package votesboxquestionmodel

import (
	"justasking/GO/core/model/votesboxquestionanswer"
	"time"

	uuid "github.com/satori/go.uuid"
)

// VotesBoxQuestion is a votesboxquestion
type VotesBoxQuestion struct {
	QuestionId uuid.UUID
	BoxId      uuid.UUID
	Header     string
	IsActive   bool
	SortOrder  int
	CreatedAt  time.Time
	CreatedBy  string
	UpdatedAt  *time.Time
	UpdatedBy  string
	DeletedAt  *time.Time
	Answers    []votesboxquestionanswermodel.VotesBoxQuestionAnswer `gorm:"-"`
}

// TableName returns the table name for use with ORM
func (VotesBoxQuestion) TableName() string {
	return "votes_box_questions"
}
