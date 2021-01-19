package baseboxmodel

import (
	"time"

	"github.com/satori/go.uuid"
)

// BaseBox stores the common box fields
type BaseBox struct {
	ID                      uuid.UUID
	Code                    string
	OriginalCode            string
	AccountId               uuid.UUID
	BoxTypeId               int
	BoxType                 string
	ThemeId                 int
	Theme                   string
	IsLive                  bool
	EntryPageEnabled        bool
	PresentationPageEnabled bool
	LoginRequired           bool
	SmsEnabled              bool
	PhoneNumber             string
	CreatedAt               time.Time
	CreatedBy               string //ID of user who created the box
	UpdatedAt               *time.Time
	UpdatedBy               string //ID of user who updated the box
	DeletedAt               *time.Time
}

// TableName returns the table name for use with ORM
func (BaseBox) TableName() string {
	return "base_box"
}
