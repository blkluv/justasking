package applogsdomain

import (
	"bytes"
	"os"
	"strings"

	emaildomain "github.com/chande/justasking/core/domain/email"
	applogsmodel "github.com/chande/justasking/core/model/applogs"
	applogsrepo "github.com/chande/justasking/core/repo/applogs"
	emailtemplaterepo "github.com/chande/justasking/core/repo/emailtemplate"
)

// LogInfo creates a record in the log table
func LogInfo(domainName, functionName, errorMessage string) {
	Log(domainName, functionName, errorMessage, "Info", false)
}

// LogWarn creates a record in the log table
func LogWarn(domainName, functionName, errorMessage string) {
	Log(domainName, functionName, errorMessage, "Warn", false)
}

// LogError creates a record in the log table
func LogError(domainName, functionName, errorMessage string, sendEmail bool) {
	Log(domainName, functionName, errorMessage, "Error", sendEmail)
}

// Log creates the log
func Log(domainName, functionName, errorMessage, logType string, sendEmail bool) {
	var sourceName bytes.Buffer
	sourceName.WriteString(domainName)
	sourceName.WriteString(" - ")
	sourceName.WriteString(functionName)

	var machineName, _ = os.Hostname()

	log := applogsmodel.AppLog{LogType: logType, SourceName: sourceName.String(), Message: errorMessage, MachineName: machineName}
	applogsrepo.InsertLog(log)

	if sendEmail {
		errorEmailTemplate, _ := emailtemplaterepo.GetEmailTemplateByName("error_email")
		errorEmailTemplate.Subject = strings.Replace(errorEmailTemplate.Subject, "{DomainName}", domainName, -1)
		errorEmailTemplate.Body = strings.Replace(errorEmailTemplate.Body, "{ErrorDetails}", errorMessage, -1)
		emaildomain.SendEmail(errorEmailTemplate)
	}
}
