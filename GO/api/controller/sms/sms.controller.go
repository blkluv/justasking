package smscontroller

import (
	"net/http"

	"justasking/GO/core/model/twilio/sms"

	"justasking/GO/core/domain/sms"

	"github.com/blue-jay/core/router"
	"github.com/gorilla/schema"
)

// Load the routes.
func Load() {
	router.Post("/sms/receive", SmsReceive)
}

// SmsReceive receives SMS data and decides what to do with it
func SmsReceive(w http.ResponseWriter, r *http.Request) {
	smsMessage := new(smsmodel.Sms)
	err := r.ParseForm()

	if err != nil {
		// Handle error
	}

	decoder := schema.NewDecoder()
	err = decoder.Decode(smsMessage, r.PostForm)

	if err != nil {
		result := smsdomain.HandleSms(*smsMessage)
		if result.IsSuccess() {

		} else {

		}
	}
}
