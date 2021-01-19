package idpmappingmodel

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// IdpMapping maps users to the IDPs they are registered with
type IdpMapping struct {
	UserId    uuid.UUID
	IdpId     int
	Sub       string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// TableName returns the table name for use with ORM
func (IdpMapping) TableName() string {
	return "idp_users"
}
