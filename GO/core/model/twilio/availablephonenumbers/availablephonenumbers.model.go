package availablephonenumbersmodel

import (
	twiliophonenumbermodel "github.com/chande/justasking/core/model/twilio/phonenumber"
)

// AvailablePhoneNumbers models the twilio response for available phone numbers
type AvailablePhoneNumbers struct {
	Uri          string                                     `json:"uri"`
	PhoneNumbers []twiliophonenumbermodel.TwilioPhoneNumber `json:"available_phone_numbers"`
}
