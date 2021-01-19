package answerboxmodel

import (
	"justasking/GO/core/model/answerboxquestion"
	"justasking/GO/core/model/boxes/basebox"
	"time"

	uuid "github.com/satori/go.uuid"
)

// AnswerBox is an answerbox
type AnswerBox struct {
	BoxId     uuid.UUID
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt *time.Time
	UpdatedBy string
	DeletedAt *time.Time
	Questions []answerboxquestionmodel.AnswerBoxQuestion `gorm:"-"`
	BaseBox   baseboxmodel.BaseBox                       `gorm:"-"`
}

// TableName returns the table name for use with ORM
func (AnswerBox) TableName() string {
	return "answer_box"
}
