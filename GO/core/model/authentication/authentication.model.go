package authenticationmodel

import "time"

// IdpGoogle is the model for the idp_google table
type IdpGoogle struct {
	ID         int
	Sub        string
	Name       *string
	Email      *string
	ImageUrl   *string
	GivenName  *string
	FamilyName *string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}

// TableName returns the table name for use with ORM
func (IdpGoogle) TableName() string {
	return "idp_google"
}
