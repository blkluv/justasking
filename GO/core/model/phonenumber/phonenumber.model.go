package phonenumbermodel

import (
	"github.com/satori/go.uuid"
)

// PhoneNumber stores the common box fields
type PhoneNumber struct {
	ID           uuid.UUID
	Sid          string
	FriendlyName string
	PhoneNumber  string
	Region       string
	IsoCountry   string
	Voice        string
	Sms          string
	Mms          string
}

// TableName returns the table name for use with ORM
func (PhoneNumber) TableName() string {
	return "phone_numbers"
}
