package accountusermodel

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// AccountUser is an AccountUser
type AccountUser struct {
	AccountId      uuid.UUID
	AccountName    string
	UserId         uuid.UUID
	RoleId         uuid.UUID
	RoleName       string
	IsActive       bool
	CurrentAccount bool
	TokenVersion   uuid.UUID
	IsPending      bool
	Email          string
	CreatedAt      time.Time
	CreatedBy      string
	UpdatedAt      *time.Time
	UpdatedBy      string
	DeletedAt      *time.Time
}

// TableName returns the table name for use with ORM
func (AccountUser) TableName() string {
	return "account_users"
}
