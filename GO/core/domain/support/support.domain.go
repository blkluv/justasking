package supportdomain

import (
	"fmt"
	"justasking/GO/common/authenticationclaim"
	"justasking/GO/common/operationresult"
	"justasking/GO/core/domain/applogs"
	"justasking/GO/core/domain/email"
	"justasking/GO/core/model/emailtemplate"
	"justasking/GO/core/model/supportissue"
	"justasking/GO/core/repo/emailtemplate"
	"justasking/GO/core/repo/support"
	"strings"

	uuid "github.com/satori/go.uuid"
)

var domainName = "SupportDomain"

// CreateSupportIssue creates a support issue in the database and sends the user and the team a confirmation of their issue
func CreateSupportIssue(supportIssue supportissuemodel.SupportIssue, tokenClaims *authenticationclaim.AuthenticationClaim) (uuid.UUID, *operationresult.OperationResult) {
	functionName := "CreateSupportIssue"
	result := operationresult.New()

	supportIssueID, _ := uuid.NewV4()
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
