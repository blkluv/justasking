package recaptcha

import (
	"time"
)

// ReCaptchaResponse stores IDP information
type ReCaptchaResponse struct {
	Success            bool
	ChallengeTimestamp time.Time `json:"challenge_ts"`
	Hostname           string
	ErrorCodes         []string `json:"error-codes"`
}
