package passwordresetrequestmodel

import (
	"time"

	"github.com/satori/go.uuid"
)

// PasswordResetRequest is a request to reset a user's password
type PasswordResetRequest struct {
	Id        uuid.UUID
	UserId    uuid.UUID
	ResetCode string
	IsActive  bool
	ExpiresAt time.Time
	Password  string `gorm:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// TableName returns the table name for use with ORM
func (PasswordResetRequest) TableName() string {
	return "password_reset_requests"
}
