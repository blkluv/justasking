package capabilitiesmodel

// Capabilities models the list of twilio phone number capabilities
type Capabilities struct {
	Voice bool `json:"voice"`
	Sms   bool `json:"sms"`
	Mms   bool `json:"mms"`
}
