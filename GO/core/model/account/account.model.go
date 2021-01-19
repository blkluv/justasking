package accountmodel

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Account is an Account
type Account struct {
	Id        uuid.UUID
	OwnerId   uuid.UUID
	Name      string
	IsActive  bool
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt *time.Time
	UpdatedBy string
	DeletedAt *time.Time
}

// TableName returns the table name for use with ORM
func (Account) TableName() string {
	return "accounts"
}
