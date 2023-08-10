package incomingphonenumbersmodel

import (
	twiliophonenumbermodel "github.com/chande/justasking/core/model/twilio/phonenumber"
)

// IncomingPhoneNumbers models the twilio response for incoming phone numbers
type IncomingPhoneNumbers struct {
	Uri          string                                     `json:"uri"`
	PhoneNumbers []twiliophonenumbermodel.TwilioPhoneNumber `json:"incoming_phone_numbers"`
}
