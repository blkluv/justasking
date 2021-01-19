package usermodel

import (
	"justasking/GO/core/model/account"
	"justasking/GO/core/model/priceplan"
	"justasking/GO/core/model/role"
	"time"

	uuid "github.com/satori/go.uuid"
)

// User is for user data contained within JustAsking, which may be different from IDP data
type User struct {
	ID                uuid.UUID
	FirstName         string
	LastName          string
	Email             string
	ImageUrl          string
	LastLoggedInAt    time.Time
	IsActive          bool
	MembershipDetails priceplanmodel.PricePlan `gorm:"-"`
	Account           accountmodel.Account     `gorm:"-"`
	RolePermissions   rolemodel.Role           `gorm:"-"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         *time.Time
}
