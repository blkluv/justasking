package featurerequestrepo

import (
	"time"

	featurerequestmodel "github.com/chande/justasking/core/model/featurerequest"
	"github.com/chande/justasking/core/startup/flight"
)

// InsertFeatureRequest inserts a new feature request
func InsertFeatureRequest(featurerequest featurerequestmodel.FeatureRequest) error {
	db := flight.Context(nil, nil).DB

	err := db.Exec(`
		INSERT INTO justasking.feature_requests
		(feature_request_id,feature_request,created_at,created_by)
		VALUES (?,?,?,?);`,
		featurerequest.FeatureRequestId, featurerequest.FeatureRequest, time.Now(), featurerequest.CreatedBy).Error

	return err
}
