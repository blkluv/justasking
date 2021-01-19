package customplanlicensemodel

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// CustomPlanLicense is a Custom Plan License
type CustomPlanLicense struct {
	Id          uuid.UUID `json:"-"`
	AccountId   uuid.UUID
	UserId      uuid.UUID
	PlanId      uuid.UUID
	LicenseCode string
	IsActive    bool
	CreatedAt   time.Time
	CreatedBy   string
	UpdatedAt   *time.Time
	UpdatedBy   string
	DeletedAt   *time.Time
}

// TableName returns the table name for use with ORM
func (CustomPlanLicense) TableName() string {
	return "custom_plan_licenses"
}
