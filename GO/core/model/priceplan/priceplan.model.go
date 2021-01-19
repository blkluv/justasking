package priceplanmodel

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// PricePlan is a PricePlan
type PricePlan struct {
	Id                 uuid.UUID
	Name               string
	DisplayName        string
	Description        string
	Price              int
	PriceDescription   string
	ImagePath          string
	Responses          int
	ActiveBoxesLimit   int
	WordCloud          bool
	QuestionBox        bool
	AnswerBox          bool
	VotesBox           bool
	ToggleResponses    bool
	Sms                bool
	CustomCode         bool
	Delegates          int
	Support            bool
	FeatureName        string `json:"-"`
	FeatureDescription string `json:"-"`
	FeatureValue       string `json:"-"`
	ExpiresInDays      int
	IsPublic           bool
	IsActive           bool
	SortOrder          int
	PeriodEnd          time.Time
	CreatedAt          time.Time
	UpdatedAt          *time.Time
	DeletedAt          *time.Time
}

// TableName returns the table name for use with ORM
func (PricePlan) TableName() string {
	return "price_plans"
}
