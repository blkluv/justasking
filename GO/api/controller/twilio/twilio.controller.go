package twiliocontroller

import (
	"encoding/json"
	"net/http"

	"justasking/GO/core/domain/phonenumbers"

	"github.com/blue-jay/core/router"
)

// Load the routes.
func Load() {
	router.Get("/twilio/availablenumbers", GetAvailableNumbers)
	router.Get("/twilio/ournumbers", GetOurNumbers)
	router.Get("/twilio/ournumbers/import", ImportOurNumbers)
}

// GetAvailableNumbers gets numbers available for purchase from twilio
func GetAvailableNumbers(w http.ResponseWriter, r *http.Request) {
	phoneNumbers, result := phonenumbersdomain.GetPhoneNumbersForPurchase()
	if result.IsSuccess() {
		responseString, err := json.Marshal(phoneNumbers)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write([]byte(responseString))
		}

	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// GetOurNumbers gets numbers available for purchase from twilio
func GetOurNumbers(w http.ResponseWriter, r *http.Request) {
	phoneNumbers, result := phonenumbersdomain.GetOurTwilioPhoneNumbers()
	if result.IsSuccess() {
		responseString, err := json.Marshal(phoneNumbers)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write([]byte(responseString))
		}

	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// ImportOurNumbers gets numbers available for purchase from twilio
func ImportOurNumbers(w http.ResponseWriter, r *http.Request) {
	phoneNumbers, result := phonenumbersdomain.ImportTwilioNumbers()
	if result.IsSuccess() {
		responseString, err := json.Marshal(phoneNumbers)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write([]byte(responseString))
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
