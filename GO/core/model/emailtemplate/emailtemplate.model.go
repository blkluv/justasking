package emailtemplatemodel

import (
	"time"
)

// EmailTemplate is an email template
type EmailTemplate struct {
	Id        int
	Name      string
	IsActive  bool
	To        string
	Cc        string
	Bcc       string
	From      string
	Subject   string
	Body      string
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

// TableName returns the table name for use with ORM
func (EmailTemplate) TableName() string {
	return "email_templates"
}
