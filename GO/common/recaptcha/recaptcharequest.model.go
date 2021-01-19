package recaptcha

// ReCaptchaRequest models the post data Google expects for a captcha validation
type ReCaptchaRequest struct {
	Secret   string
	Response string //this is the captcha token
}
