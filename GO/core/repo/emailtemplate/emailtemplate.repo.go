package emailtemplaterepo

import (
	"justasking/GO/core/model/emailtemplate"
	"justasking/GO/core/startup/flight"
)

// GetEmailTemplateByName gets email template for a given name
func GetEmailTemplateByName(name string) (emailtemplatemodel.EmailTemplate, error) {
	db := flight.Context(nil, nil).DB

	var emailTemplate emailtemplatemodel.EmailTemplate

	err := db.Raw("SELECT `id`,`name`,`is_active`,`to`,`cc`,`bcc`,`from`,`subject`,`body`,`created_at`,`updated_at`,`deleted_at` FROM email_templates WHERE name = ? ", name).Scan(&emailTemplate).Error

	return emailTemplate, err
}
