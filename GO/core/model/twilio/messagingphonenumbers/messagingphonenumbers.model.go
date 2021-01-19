package messagingphonenumbersmodel

import (
	"justasking/GO/core/model/twilio/phonenumber"
)

// MessagingPhoneNumbers models the twilio response for incoming phone numbers
type MessagingPhoneNumbers struct {
	PhoneNumbers []twiliophonenumbermodel.TwilioPhoneNumber `json:"phone_numbers"`
}
