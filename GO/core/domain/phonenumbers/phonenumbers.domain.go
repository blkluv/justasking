package phonenumbersdomain

import (
	"fmt"
	"justasking/GO/common/clients/twilio"
	"justasking/GO/common/operationresult"
	"justasking/GO/core/domain/appconfigs"
	"justasking/GO/core/domain/applogs"
	"justasking/GO/core/model/phonenumber"
	"justasking/GO/core/model/twilio/availablephonenumbers"
	"justasking/GO/core/model/twilio/incomingphonenumbers"
	"justasking/GO/core/model/twilio/phonenumber"
	"justasking/GO/core/repo/phonenumber"
)

var domainName = "PhoneNumbersDomain"

// GetOurTwilioPhoneNumbers returns a list of ALL (regardless of messaging service) our phone numbers in Twilio
func GetOurTwilioPhoneNumbers() (incomingphonenumbersmodel.IncomingPhoneNumbers, *operationresult.OperationResult) {
	functionName := "GetOurTwilioPhoneNumbers"
	result := operationresult.New()
	var phoneNumbers incomingphonenumbersmodel.IncomingPhoneNumbers
	var err error

	twilioConfigs, configsResult := appconfigsdomain.GetAppConfigs("twilio")
	if configsResult.IsSuccess() {
		phoneNumbers, err = twilioclient.GetOurTwilioPhoneNumbers(twilioConfigs)
		if err != nil {
			msg := err.Error()
			result = operationresult.CreateErrorResult(msg, err)
			applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting phone numbers from Twilio. Error: [%v]", msg), false)
		}
	} else {
		msg := configsResult.Error.Error()
		result = operationresult.CreateErrorResult(msg, configsResult.Error)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting app configs. Error: [%v]", msg), false)
	}

	return phoneNumbers, result
}

// ImportTwilioNumbers imports all our phone numbers from Twilio
func ImportTwilioNumbers() (incomingphonenumbersmodel.IncomingPhoneNumbers, *operationresult.OperationResult) {
	functionName := "ImportTwilioNumbers"
	result := operationresult.New()
	var err error
	var twilioPhoneNumbers incomingphonenumbersmodel.IncomingPhoneNumbers

	twilioConfigs, configsResult := appconfigsdomain.GetAppConfigs("twilio")
	if configsResult.IsSuccess() {
		twilioPhoneNumbers, err = twilioclient.GetOurTwilioPhoneNumbers(twilioConfigs)

		if err != nil {
			msg := err.Error()
			result = operationresult.CreateErrorResult(msg, err)
			applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting phone numbers from Twilio. Error: [%v]", msg), false)
		} else {
			for _, phoneNumber := range twilioPhoneNumbers.PhoneNumbers {
				err = phonenumberrepo.InsertPhoneNumber(phoneNumber)
				if err != nil {
					msg := err.Error()
					result = operationresult.CreateErrorResult(msg, err)
					applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error inserting phone number into database. Sid: [%v]. Phone Number: [%v]. Error: [%v]", phoneNumber.Sid, phoneNumber.PhoneNumber, msg), false)
				} else {
					applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("Successfully inserted phone number [%v]", phoneNumber.PhoneNumber))
				}
			}
		}
	} else {
		msg := configsResult.Error.Error()
		result = operationresult.CreateErrorResult(msg, configsResult.Error)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting app configs. Error: [%v]", msg), false)
	}

	return twilioPhoneNumbers, result
}

// GetAllActiveJustAskingPhoneNumbers gets all active phone numbers from our database
func GetAllActiveJustAskingPhoneNumbers() ([]phonenumbermodel.PhoneNumber, *operationresult.OperationResult) {
	functionName := "GetAllActiveJustAskingPhoneNumbers"
	result := operationresult.New()

	phoneNumbers, err := phonenumberrepo.GetAllActiveJustAskingPhoneNumbers()
	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Unable to retrieve active JustAsking phone numbers."), false)
	}

	return phoneNumbers, result
}

// GetPhoneNumbersForPurchase returns a list of available numbers for purchase in Twilio
func GetPhoneNumbersForPurchase() (availablephonenumbersmodel.AvailablePhoneNumbers, *operationresult.OperationResult) {
	functionName := "GetPhoneNumbersForPurchase"
	result := operationresult.New()
	var phoneNumbers availablephonenumbersmodel.AvailablePhoneNumbers
	var err error

	twilioConfigs, configsResult := appconfigsdomain.GetAppConfigs("twilio")
	if configsResult.IsSuccess() {
		phoneNumbers, err = twilioclient.GetPhoneNumbersForPurchase(twilioConfigs)
		if err != nil {
			msg := err.Error()
			result = operationresult.CreateErrorResult(msg, err)
			applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting available phone numbers from Twilio. Error: [%v]", msg), false)
		}
	} else {
		msg := configsResult.Error.Error()
		result = operationresult.CreateErrorResult(msg, configsResult.Error)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting app configs. Error: [%v]", msg), false)
	}

	return phoneNumbers, result
}

// PurchasePhoneNumber purchases a given phone number from Twilio
func PurchasePhoneNumber(phoneNumber twiliophonenumbermodel.TwilioPhoneNumber) (twiliophonenumbermodel.TwilioPhoneNumber, *operationresult.OperationResult) {
	functionName := "PurchasePhoneNumber"
	result := operationresult.New()
	var err error

	twilioConfigs, configsResult := appconfigsdomain.GetAppConfigs("twilio")
	if configsResult.IsSuccess() {
		phoneNumber, err = twilioclient.PurchasePhoneNumber(twilioConfigs, phoneNumber.PhoneNumber)
		if err != nil {
			msg := err.Error()
			result = operationresult.CreateErrorResult(msg, err)
			applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error purchasing a phone number from Twilio. Error: [%v]", msg), false)
		}
	} else {
		msg := configsResult.Error.Error()
		result = operationresult.CreateErrorResult(msg, configsResult.Error)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting app configs. Error: [%v]", msg), false)
	}

	return phoneNumber, result
}

// ReleasePhoneNumber releases a phone number from Twilio
func ReleasePhoneNumber() (phonenumbermodel.PhoneNumber, *operationresult.OperationResult) {
	functionName := "ReleasePhoneNumber"
	result := operationresult.New()
	var releasedPhoneNumber phonenumbermodel.PhoneNumber
	var err error

	twilioConfigs, configsResult := appconfigsdomain.GetAppConfigs("twilio")
	if configsResult.IsSuccess() {
		releasedPhoneNumber, err = phonenumberrepo.ReleasePhoneNumber(twilioConfigs)
		if err != nil {
			msg := err.Error()
			result = operationresult.CreateErrorResult(msg, err)
			applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error releasing a phone number from Twilio. Error: [%v]", msg), false)
		}
	} else {
		msg := configsResult.Error.Error()
		result = operationresult.CreateErrorResult(msg, configsResult.Error)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting app configs. Error: [%v]", msg), false)
	}

	return releasedPhoneNumber, result
}
