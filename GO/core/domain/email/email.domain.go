package emaildomain

import (
	"justasking/GO/common/operationresult"
	"justasking/GO/core/model/emailtemplate"
	"justasking/GO/core/startup/flight"
	"strings"

	gomail "gopkg.in/gomail.v2"
)

// SendEmail  sends email given a body
func SendEmail(emailTemplate emailtemplatemodel.EmailTemplate) *operationresult.OperationResult {
	result := operationresult.New()

	config := flight.Context(nil, nil).Config
	emailAddress := config.Settings["NoReplyEmail"]
	emailPassword := config.Settings["NoReplyPassword"]

	result = SendEmailWithCredentials(emailTemplate, emailAddress, emailPassword)

	return result
}

// SendEmailWithCredentials  sends email given a body and creds
func SendEmailWithCredentials(emailTemplate emailtemplatemodel.EmailTemplate, fromEmailAddress string, fromEmailPassword string) *operationresult.OperationResult {
	result := operationresult.New()

	var cc []string
	var bcc []string

	to := strings.Split(emailTemplate.To, ",")
	if len(emailTemplate.Cc) > 0 {
		cc = strings.Split(emailTemplate.Cc, ",")
	}
	if len(emailTemplate.Bcc) > 0 {
		bcc = strings.Split(emailTemplate.Bcc, ",")
	}

	m := setGoMailParameters(to, cc, bcc, emailTemplate.Subject, emailTemplate.Body)

	d := gomail.NewPlainDialer("smtp.gmail.com", 587, fromEmailAddress, fromEmailPassword)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	return result
}

func setGoMailParameters(to []string, cc []string, bcc []string, subject string, body string) *gomail.Message {

	m := gomail.NewMessage()
	m.SetAddressHeader("From", "noreply@justasking.io", "Justasking.io")
	m.SetHeader("To", to...)
	if len(cc) > 0 {
		m.SetHeader("Cc", cc...)
	}
	if len(bcc) > 0 {
		m.SetHeader("Bcc", bcc...)
	}
	m.SetHeader("MIME-Version", "1.0")
	m.SetHeader("Content-type", "text/html")
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	return m
}
