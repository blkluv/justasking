package featurerequestmodel

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// FeatureRequest is a response to a word cloud
type FeatureRequest struct {
	FeatureRequestId uuid.UUID
	FeatureRequest   string
	CreatedAt        time.Time
	CreatedBy        string
}

// TableName returns the table name for use with ORM
func (FeatureRequest) TableName() string {
	return "feature_requests"
}
