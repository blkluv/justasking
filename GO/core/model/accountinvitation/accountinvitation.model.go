package accountinvitationmodel

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// AccountInvitation is an Account Invitation
type AccountInvitation struct {
	Id             uuid.UUID `json:"-"`
	AccountId      uuid.UUID
	AccountName    string
	RoleId         uuid.UUID
	RoleName       string
	Email          string
	InvitationCode string
	IsActive       bool
	CreatedAt      time.Time
	CreatedBy      string
	UpdatedAt      *time.Time
	UpdatedBy      string
	DeletedAt      *time.Time
}

// TableName returns the table name for use with ORM
func (AccountInvitation) TableName() string {
	return "account_invitations"
}
