package answerboxquestionmodel

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// AnswerBoxQuestion is an answerbox question
type AnswerBoxQuestion struct {
	QuestionId uuid.UUID
	BoxId      uuid.UUID
	Question   string
	IsActive   bool
	SortOrder  int
	CreatedAt  time.Time
	CreatedBy  string
	UpdatedAt  *time.Time
	UpdatedBy  string
	DeletedAt  *time.Time
}

// TableName returns the table name for use with ORM
func (AnswerBoxQuestion) TableName() string {
	return "answer_box_questions"
}
