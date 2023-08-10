package featurerequestdomain

import (
	"fmt"
	"strings"

	"github.com/chande/justasking/common/authenticationclaim"
	"github.com/chande/justasking/common/operationresult"
	applogsdomain "github.com/chande/justasking/core/domain/applogs"
	emaildomain "github.com/chande/justasking/core/domain/email"
	emailtemplatemodel "github.com/chande/justasking/core/model/emailtemplate"
	featurerequestmodel "github.com/chande/justasking/core/model/featurerequest"
	emailtemplaterepo "github.com/chande/justasking/core/repo/emailtemplate"
	featurerequestrepo "github.com/chande/justasking/core/repo/featurerequest"

	uuid "github.com/satori/go.uuid"
)

var domainName = "FeatureRequestDomain"

// CreateFeatureRequest creates a feature request from the user
func CreateFeatureRequest(featureRequest featurerequestmodel.FeatureRequest, tokenClaims *authenticationclaim.AuthenticationClaim) (uuid.UUID, *operationresult.OperationResult) {
	functionName := "CreateFeatureRequest"
	result := operationresult.New()

	featureRequestID := uuid.NewV4()
	featureRequest.FeatureRequestId = featureRequestID

	err := featurerequestrepo.InsertFeatureRequest(featureRequest)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error creating Feature Request. Error: [%v]", msg), false)
	} else {
		applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("Feature Request created for user: [%v]", featureRequest.CreatedBy))

		var featureRequestTemplate emailtemplatemodel.EmailTemplate
		featureRequestTemplate, err = emailtemplaterepo.GetEmailTemplateByName("feature_request_thank_you")

		if err != nil {
			applogsdomain.LogError(domainName, functionName, "Unable to retrieve feature request template.", false)
		} else {
			featureRequestTemplate.To = tokenClaims.Email
			featureRequestTemplate.Body = strings.Replace(featureRequestTemplate.Body, "{FeatureRequest}", featureRequest.FeatureRequest, -1)
			emailSendResult := emaildomain.SendEmail(featureRequestTemplate)
			if emailSendResult.IsSuccess() {
				applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("Feature Request email sent to user [%v].", tokenClaims.ID))
			} else {
				applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Unable to send Feature Request email to user [%v]. Error: [%v]", tokenClaims.ID, emailSendResult.Message), false)
			}
		}
	}

	return featureRequestID, result
}
