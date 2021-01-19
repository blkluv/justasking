package userstripemodel

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// UserStripe is for data mapping justasking users to stripe customers
type UserStripe struct {
	UserId             uuid.UUID
	AccountId          uuid.UUID
	StripeUserId       string
	StripePaymentToken string
	CreditCardLastFour string
	LastPayment        *time.Time
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          *time.Time
}

// TableName returns the table name for use with ORM
func (UserStripe) TableName() string {
	return "users_stripe"
}
