package twiliophonenumbermodel

import (
	"justasking/GO/core/model/twilio/capabilities"
)

// TwilioPhoneNumber models a twilio phone number
type TwilioPhoneNumber struct {
	FriendlyName string                         `json:"friendly_name"`
	PhoneNumber  string                         `json:"phone_number"`
	Sid          string                         `json:"sid"`
	Lata         string                         `json:"lata"`
	RateCenter   string                         `json:"rate_center"`
	Latitude     string                         `json:"latitude"`
	Longitude    string                         `json:"longitude"`
	Region       string                         `json:"region"`
	PostalCode   string                         `json:"postal_code"`
	IsoCountry   string                         `json:"iso_country"`
	Capabilities capabilitiesmodel.Capabilities `json:"capabilities"`
	Beta         bool                           `json:"beta"`
}
