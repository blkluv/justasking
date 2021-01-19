package smsmodel

// Sms models the object for an Sms message coming from Twilio
type Sms struct {
	MessageSid          string
	SmsSid              string
	AccountSid          string
	MessagingServiceSid string
	From                string
	To                  string
	Body                string
	NumMedia            string
	FromCity            string
	FromState           string
	FromZip             string
	FromCountry         string
	ToCity              string
	ToState             string
	ToZip               string
	ToCountry           string
}
