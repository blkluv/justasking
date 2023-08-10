package recaptchaclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/chande/justasking/common/operationresult"
	"github.com/chande/justasking/common/recaptcha"
	appconfigsdomain "github.com/chande/justasking/core/domain/appconfigs"
	applogsdomain "github.com/chande/justasking/core/domain/applogs"
)

var clientName = "ReCaptchaClient"

// ValidateReCaptchaToken validates a recaptcha token by calling out to Google
func ValidateReCaptchaToken(reCaptchaToken string) bool {
	functionName := "ValidateReCaptchaToken"
	var reCaptchaRequest recaptcha.ReCaptchaRequest
	var reCaptchaResponse recaptcha.ReCaptchaResponse
	isValid := false
	var body []byte

	configs, result := appconfigsdomain.GetAppConfigs("recaptcha")
	if result.IsSuccess() {
		reCaptchaUrl := configs["ReCaptchaUrl"]
		reCaptchaRequest.Secret = configs["ReCaptchaSecretKey"]
		reCaptchaRequest.Response = reCaptchaToken
		//postData, _ := json.Marshal(reCaptchaRequest)

		fmt.Println(reCaptchaUrl)
		fmt.Println(reCaptchaRequest.Secret)
		payload := strings.NewReader(fmt.Sprintf("secret=%v&response=%v", reCaptchaRequest.Secret, reCaptchaRequest.Response))

		req, err := http.NewRequest("POST", reCaptchaUrl, payload)
		if err != nil {
			result.Status = operationresult.Error
			result.Message = fmt.Sprintf("Error creating request object to send to Google for reCaptcha. Error: [%v]", err.Error())
			result.Error = err
		} else {
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

			client := getHTTPClient()
			resp, err := client.Do(req)
			if err != nil {
				result.Status = operationresult.Error
				result.Message = err.Error()
				result.Error = err
			} else if resp.StatusCode != 200 {
				applogsdomain.LogError(clientName, functionName, fmt.Sprintf("Error making request to recaptcha. Message: [%v]", err), false)
			} else {
				defer resp.Body.Close()
				body, err = ioutil.ReadAll(resp.Body)
				if err != nil {
					applogsdomain.LogError(clientName, functionName, fmt.Sprintf("Error reading response from reCaptcha. Message: [%v]", err), false)
				} else {
					err = json.Unmarshal([]byte(body), &reCaptchaResponse)
					if err != nil {
						applogsdomain.LogError(clientName, functionName, fmt.Sprintf("Error unmarshaling response from reCaptcha. Message: [%v]", err), false)
					} else {
						if reCaptchaResponse.Success == true {
							isValid = true
						} else {
							if len(reCaptchaResponse.ErrorCodes) > 0 {
								errors := strings.Join(reCaptchaResponse.ErrorCodes, ",")
								applogsdomain.LogError(clientName, functionName, fmt.Sprintf("Error getting response from recaptcha. Message: [%v]", errors), false)
							}
						}
					}
				}
			}
		}
	} else {
		applogsdomain.LogError(clientName, functionName, fmt.Sprintf("Error getting app configs. Error: [%v]", result.Message), false)
	}

	return isValid
}

func getHTTPClient() http.Client {
	var client = &http.Client{
		Timeout: time.Second * 20,
	}

	return *client
}
