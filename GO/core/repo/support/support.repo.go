package supportrepo

import (
	"justasking/GO/core/model/supportissue"
	"justasking/GO/core/startup/flight"
	"time"
)

// InsertSupportIssue inserts a new support issue on the system
func InsertSupportIssue(supportIssue supportissuemodel.SupportIssue) error {
	db := flight.Context(nil, nil).DB

	err := db.Exec(`
		INSERT INTO justasking.support_issues
		(issue_id,issue,user_agent,resolved,created_at,created_by)
		VALUES (?,?,?,?,?,?);`,
		supportIssue.IssueId, supportIssue.Issue, supportIssue.UserAgent, supportIssue.Resolved, time.Now(), supportIssue.CreatedBy).Error

	return err
}
