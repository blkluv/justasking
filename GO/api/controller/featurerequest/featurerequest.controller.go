package featurerequestcontroller

import (
	"encoding/json"
	"net/http"

	"github.com/chande/justasking/api/startup/middleware"
	"github.com/chande/justasking/common/operationresult"
	featurerequestdomain "github.com/chande/justasking/core/domain/featurerequest"
	featurerequestmodel "github.com/chande/justasking/core/model/featurerequest"

	"github.com/blue-jay/core/router"
)

// Load the routes.
func Load() {
	router.Post("/featurerequest", CreateFeatureRequest, middleware.AuthorizedHandler)
}

// CreateFeatureRequest creates a feature request from the user
func CreateFeatureRequest(w http.ResponseWriter, r *http.Request) {
	var featureRequest featurerequestmodel.FeatureRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&featureRequest)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		tokenClaims := middleware.GetTokenClaims(r)
		featureRequest.CreatedBy = tokenClaims.ID.String()

		featureRequestId, result := featurerequestdomain.CreateFeatureRequest(featureRequest, tokenClaims)
		if result.IsSuccess() {
			responseString, err := json.Marshal(featureRequestId)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				w.Write([]byte(responseString))
			}
		} else if result.Status == operationresult.Forbidden {
			w.WriteHeader(http.StatusForbidden)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
