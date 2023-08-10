package supportdomain

import (
	"fmt"
	"strings"

	"github.com/chande/justasking/common/authenticationclaim"
	"github.com/chande/justasking/common/operationresult"
	applogsdomain "github.com/chande/justasking/core/domain/applogs"
	emaildomain "github.com/chande/justasking/core/domain/email"
	emailtemplatemodel "github.com/chande/justasking/core/model/emailtemplate"
	supportissuemodel "github.com/chande/justasking/core/model/supportissue"
	emailtemplaterepo "github.com/chande/justasking/core/repo/emailtemplate"
	supportrepo "github.com/chande/justasking/core/repo/support"

	uuid "github.com/satori/go.uuid"
)

var domainName = "SupportDomain"

// CreateSupportIssue creates a support issue in the database and sends the user and the team a confirmation of their issue
func CreateSupportIssue(supportIssue supportissuemodel.SupportIssue, tokenClaims *authenticationclaim.AuthenticationClaim) (uuid.UUID, *operationresult.OperationResult) {
	functionName := "CreateSupportIssue"
	result := operationresult.New()

	supportIssueID := uuid.NewV4()
	supportIssue.IssueId = supportIssueID

	err := supportrepo.InsertSupportIssue(supportIssue)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error creating Support Issue. Error: [%v]", msg), false)
	} else {
		applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("Support Issue created for user: [%v]", supportIssue.CreatedBy))

		var supportIssueTempalte emailtemplatemodel.EmailTemplate
		supportIssueTempalte, err = emailtemplaterepo.GetEmailTemplateByName("support_issue")

		if err != nil {
			applogsdomain.LogError(domainName, functionName, "Unable to retrieve support issue email template.", false)
		} else {
			supportIssueTempalte.To = tokenClaims.Email
			supportIssueTempalte.Body = strings.Replace(supportIssueTempalte.Body, "{IssueId}", supportIssue.IssueId.String(), -1)
			supportIssueTempalte.Body = strings.Replace(supportIssueTempalte.Body, "{Issue}", supportIssue.Issue, -1)
			emailSendResult := emaildomain.SendEmail(supportIssueTempalte)
			if emailSendResult.IsSuccess() {
				applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("Support issue email sent to user [%v].", tokenClaims.ID))
			} else {
				applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Unable to send support issue email to user [%v]. Error: [%v]", tokenClaims.ID, emailSendResult.Message), false)
			}
		}
	}

	return supportIssueID, result
}
