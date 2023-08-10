package supportcontroller

import (
	"encoding/json"
	"net/http"

	"github.com/chande/justasking/api/startup/middleware"
	"github.com/chande/justasking/common/operationresult"
	supportdomain "github.com/chande/justasking/core/domain/support"
	supportissuemodel "github.com/chande/justasking/core/model/supportissue"

	"github.com/blue-jay/core/router"
)

// Load the routes.
func Load() {
	router.Post("/supportissue", CreateSupportIssue, middleware.AuthorizedHandler)
}

// CreateSupportIssue creates a support issue for the user
func CreateSupportIssue(w http.ResponseWriter, r *http.Request) {
	var supportIssue supportissuemodel.SupportIssue

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&supportIssue)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		tokenClaims := middleware.GetTokenClaims(r)
		supportIssue.CreatedBy = tokenClaims.ID.String()
		supportIssue.Resolved = false

		supportIssueId, result := supportdomain.CreateSupportIssue(supportIssue, tokenClaims)
		if result.IsSuccess() {
			responseString, err := json.Marshal(supportIssueId)
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
