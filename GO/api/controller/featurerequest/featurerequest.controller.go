package featurerequestcontroller

import (
	"encoding/json"
	"justasking/GO/api/startup/middleware"
	"justasking/GO/common/operationresult"
	"justasking/GO/core/domain/featurerequest"
	"justasking/GO/core/model/featurerequest"
	"net/http"

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
