package questionboxmodel

import (
	"time"

	"justasking/GO/core/model/boxes/basebox"

	uuid "github.com/satori/go.uuid"
)

// QuestionBox is a questionbox
type QuestionBox struct {
	BoxId     uuid.UUID
	Header    string
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt *time.Time
	UpdatedBy string
	DeletedAt *time.Time
	BaseBox   baseboxmodel.BaseBox `gorm:"-"`
}

// TableName returns the table name for use with ORM
func (QuestionBox) TableName() string {
	return "question_box"
}
