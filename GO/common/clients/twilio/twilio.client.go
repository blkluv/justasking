package twilioclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"justasking/GO/core/domain/applogs"
	"justasking/GO/core/model/twilio/availablephonenumbers"
	"justasking/GO/core/model/twilio/incomingphonenumbers"
	"justasking/GO/core/model/twilio/phonenumber"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var clientName = "TwilioClient"

// GetPhoneNumbersForPurchase returns a list of available numbers from Twilio
func GetPhoneNumbersForPurchase(twilioConfigs map[string]string) (availablephonenumbersmodel.AvailablePhoneNumbers, error) {
	functionName := "GetPhoneNumbersForPurchase"
	var numbers = new(availablephonenumbersmodel.AvailablePhoneNumbers)
	var err error
	var req *http.Request
	var resp *http.Response
	var body []byte

	twilioBaseUrl := twilioConfigs["TwilioApiBaseUrl"]
	twilioGetNumbersUri := twilioConfigs["TwilioAvailableNumbersUri"]
	twilioSid := twilioConfigs["TwilioLiveAccountSid"]
	twilioAuthToken := twilioConfigs["TwilioLiveAuthToken"]

	urlToCall := fmt.Sprintf("%v%v/?%v", twilioBaseUrl, twilioGetNumbersUri, "SmsEnabled=true")
	getNumbersUrl := urlToCall

	req, err = http.NewRequest("GET", getNumbersUrl, nil)
	if err != nil {
		applogsdomain.LogError(clientName, functionName, fmt.Sprintf("Error retrieving phone numbers for purchase. Message: [%v]", err), false)
	} else {
		req.SetBasicAuth(twilioSid, twilioAuthToken)

		client := getHTTPClient()
		resp, err = client.Do(req)
		if err != nil {
			applogsdomain.LogError(clientName, functionName, fmt.Sprintf("Error making request to Twilio. Endpoint called: [%v]. Message: [%v]", getNumbersUrl, err), false)
		} else if resp.StatusCode != 200 {
			err = fmt.Errorf("received [%v] status code from Twilio", resp.StatusCode)
			applogsdomain.LogError(clientName, functionName, fmt.Sprintf("Error making request to Twilio. Endpoint called: [%v]. Message: [%v]", getNumbersUrl, err), false)
		} else {
			defer resp.Body.Close()
			body, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				applogsdomain.LogError(clientName, functionName, fmt.Sprintf("Error reading response from Twilio. Message: [%v]", err), false)
			} else {
				err = json.Unmarshal([]byte(body), &numbers)
				if err != nil {
					applogsdomain.LogError(clientName, functionName, fmt.Sprintf("Error unmarshaling response from Twilio. Message: [%v]", err), false)
				}
			}
		}
	}
	return *numbers, err
}

// GetOurTwilioPhoneNumbers returns list of all phone numbers on our account
func GetOurTwilioPhoneNumbers(twilioConfigs map[string]string) (incomingphonenumbersmodel.IncomingPhoneNumbers, error) {
	functionName := "GetOurTwilioPhoneNumbers"
	var numbers = new(incomingphonenumbersmodel.IncomingPhoneNumbers)
	var err error
	var req *http.Request
	var resp *http.Response
	var body []byte

	twilioBaseUrl := twilioConfigs["TwilioApiBaseUrl"]
	twilioOurPhoneNumbersUri := twilioConfigs["TwilioOurPhoneNumbersUri"]
	twilioSid := twilioConfigs["TwilioLiveAccountSid"]
	twilioAuthToken := twilioConfigs["TwilioLiveAuthToken"]

	getNumbersUrl := fmt.Sprintf("%v%v", twilioBaseUrl, twilioOurPhoneNumbersUri)

	req, err = http.NewRequest("GET", getNumbersUrl, nil)
	if err != nil {
		applogsdomain.LogError(clientName, functionName, fmt.Sprintf("Error retrieving our phone numbers from Twilio. Message: [%v]", err), false)
	} else {
		req.SetBasicAuth(twilioSid, twilioAuthToken)

		client := getHTTPClient()
		resp, err = client.Do(req)
		if err != nil {
			applogsdomain.LogError(clientName, functionName, fmt.Sprintf("Error making request to Twilio. Message: [%v]", err), false)
		} else if resp.StatusCode != 200 {
			err = fmt.Errorf("received [%v] status code from Twilio", resp.StatusCode)
			applogsdomain.LogError(clientName, functionName, fmt.Sprintf("Error making request to Twilio. Message: [%v]", err), false)
		} else {
			defer resp.Body.Close()
			body, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				applogsdomain.LogError(clientName, functionName, fmt.Sprintf("Error reading response from Twilio. Message: [%v]", err), false)
			} else {
				err = json.Unmarshal([]byte(body), &numbers)
				if err != nil {
					applogsdomain.LogError(clientName, functionName, fmt.Sprintf("Error unmarshaling response from Twilio. Message: [%v]", err), false)
				}
			}
		}
	}
	return *numbers, err
}

// PurchasePhoneNumber purchases a phone number from Twilio
func PurchasePhoneNumber(twilioConfigs map[string]string, numberToBuy string) (twiliophonenumbermodel.TwilioPhoneNumber, error) {
	functionName := "PurchasePhoneNumber"
	var phoneNumber = new(twiliophonenumbermodel.TwilioPhoneNumber)
	var twilioAuthToken string
	var twilioSid string
	var err error
	var req *http.Request
	var resp *http.Response

	twilioBaseUrl := twilioConfigs["TwilioApiBaseUrl"]
	twilioPurchasePhoneNumberUri := twilioConfigs["TwilioPurchasePhoneNumberUri"]
	testMode, _ := strconv.ParseBool(twilioConfigs["TwilioTestModeFlag"])
	smsUrl := twilioConfigs["TwilioSmsUrl"]

	if testMode {
		twilioSid = twilioConfigs["TwilioTestAccountSid"]
		twilioAuthToken = twilioConfigs["TwilioTestAuthToken"]
		twilioPurchasePhoneNumberUri = strings.Replace(twilioPurchasePhoneNumberUri, "{TwilioAccountSid}", twilioSid, -1)
	} else {
		twilioSid = twilioConfigs["TwilioLiveAccountSid"]
		twilioAuthToken = twilioConfigs["TwilioLiveAuthToken"]
		twilioPurchasePhoneNumberUri = strings.Replace(twilioPurchasePhoneNumberUri, "{TwilioAccountSid}", twilioSid, -1)
	}

	payload := strings.NewReader(fmt.Sprintf("PhoneNumber=%v&SmsUrl=%v", numberToBuy, smsUrl))

	purchasePhoneNumberUrl := fmt.Sprintf("%v%v", twilioBaseUrl, twilioPurchasePhoneNumberUri)

	req, err = http.NewRequest("POST", purchasePhoneNumberUrl, payload)
	if err != nil {
		applogsdomain.LogError(clientName, functionName, fmt.Sprintf("Error purchasing phone number. Message: [%v]", err), false)
	} else {
		req.SetBasicAuth(twilioSid, twilioAuthToken)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		client := getHTTPClient()
		resp, err = client.Do(req)
		if err != nil {
			applogsdomain.LogError(clientName, functionName, fmt.Sprintf("Error making request to Twilio. Message: [%v]", err), false)
		} else if resp.StatusCode != 200 && resp.StatusCode != 201 {
			err = fmt.Errorf("received [%v] status code from Twilio", resp.StatusCode)
			applogsdomain.LogError(clientName, functionName, fmt.Sprintf("Error making request to Twilio. Message: [%v]", err), false)
		}
	}

	return *phoneNumber, err
}

// ReleasePhoneNumber releases a twilio phone number
func ReleasePhoneNumber(twilioConfigs map[string]string, numberToReleaseSid string) error {
	functionName := "ReleasePhoneNumber"
	var twilioAuthToken string
	var twilioSid string
	var err error
	var req *http.Request
	var resp *http.Response

	twilioBaseUrl := twilioConfigs["TwilioApiBaseUrl"]
	twilioReleasePhoneNumberUri := twilioConfigs["TwilioReleasePhoneNumberUri"]
	testMode, _ := strconv.ParseBool(twilioConfigs["TwilioTestModeFlag"])

	if testMode {
		twilioSid = twilioConfigs["TwilioTestAccountSid"]
		twilioAuthToken = twilioConfigs["TwilioTestAuthToken"]
		twilioReleasePhoneNumberUri = strings.Replace(twilioReleasePhoneNumberUri, "{TwilioAccountSid}", twilioSid, -1)
	} else {
		twilioSid = twilioConfigs["TwilioLiveAccountSid"]
		twilioAuthToken = twilioConfigs["TwilioLiveAuthToken"]
		twilioReleasePhoneNumberUri = strings.Replace(twilioReleasePhoneNumberUri, "{TwilioAccountSid}", twilioSid, -1)
	}

	twilioReleasePhoneNumberUri = strings.Replace(twilioReleasePhoneNumberUri, "{PhoneNumberSid}", numberToReleaseSid, -1)

	releasePhoneNumberUrl := fmt.Sprintf("%v%v", twilioBaseUrl, twilioReleasePhoneNumberUri)

	req, err = http.NewRequest("DELETE", releasePhoneNumberUrl, nil)
	if err != nil {
		applogsdomain.LogError(clientName, functionName, fmt.Sprintf("Error releasing phone number. Message: [%v]", err), false)
	} else {
		req.SetBasicAuth(twilioSid, twilioAuthToken)

		client := getHTTPClient()
		resp, err = client.Do(req)
		if err != nil {
			applogsdomain.LogError(clientName, functionName, fmt.Sprintf("Error making request to Twilio. Message: [%v]", err), false)
		} else if resp.StatusCode != 200 && resp.StatusCode != 201 && resp.StatusCode != 204 {
			err = fmt.Errorf("received [%v] status code from Twilio", resp.StatusCode)
			applogsdomain.LogError(clientName, functionName, fmt.Sprintf("Error making request to Twilio. Message: [%v]", err), false)
		}
	}

	return err
}

func getHTTPClient() http.Client {
	var client = &http.Client{
		Timeout: time.Second * 20,
	}

	return *client
}
