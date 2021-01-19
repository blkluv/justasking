package idpjustaskingmodel

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// IdpJustAsking is for user login data
type IdpJustAsking struct {
	Id           int
	Email        string
	Password     string
	Sub          uuid.UUID `json:"-"`
	Name         string
	PhoneNumber  string
	ImageUrl     string
	GivenName    string
	FamilyName   string
	CaptchaToken string `gorm:"-"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}

// TableName returns the table name for use with ORM
func (IdpJustAsking) TableName() string {
	return "idp_justasking"
}
