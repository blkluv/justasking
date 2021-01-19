package smsrepo

import (
	"justasking/GO/core/startup/flight"

	"justasking/GO/core/model/twilio/sms"

	"time"

	uuid "github.com/satori/go.uuid"
)

// InsertSmsLog logs an sms message
func InsertSmsLog(smsMessage smsmodel.Sms) error {
	db := flight.Context(nil, nil).DB

	newId, _ := uuid.NewV4()
	err := db.Exec(`INSERT INTO sms_logs (log_id,message_sid,sms_sid,account_sid,message_service_sid,sms_from,sms_to,sms_body,num_media,sms_from_city,sms_from_state,
					sms_from_zip,sms_from_country,sms_to_city,sms_to_state,sms_to_zip,sms_to_country,created_at)
					VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
		newId, smsMessage.MessageSid, smsMessage.SmsSid, smsMessage.AccountSid, smsMessage.MessagingServiceSid,
		smsMessage.From, smsMessage.To, smsMessage.Body, smsMessage.NumMedia, smsMessage.FromCity, smsMessage.FromState,
		smsMessage.FromZip, smsMessage.FromCountry, smsMessage.ToCity, smsMessage.ToState, smsMessage.ToZip, smsMessage.ToCountry, time.Now()).Error

	return err
}
