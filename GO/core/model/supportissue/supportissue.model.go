package supportissuemodel

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// SupportIssue is a response to a word cloud
type SupportIssue struct {
	IssueId   uuid.UUID
	Issue     string
	UserAgent string
	Resolved  bool
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt *time.Time
	UpdatedBy string
	DeletedAt *time.Time
}

// TableName returns the table name for use with ORM
func (SupportIssue) TableName() string {
	return "support_issues"
}
