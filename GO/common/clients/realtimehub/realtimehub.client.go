package realtimehubclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"justasking/GO/common/operationresult"
	"justasking/GO/core/domain/appconfigs"
	"justasking/GO/core/model/websocketmessage"
	"net/http"
	"time"
)

// BroadcastSmsMessage calls the realtimehub api to broadcast an sms message
func BroadcastSmsMessage(code string, messageType string, messageData string) *operationresult.OperationResult {
	var websocketMessage websocketmessagemodel.WebSocketMessage

	websocketMessage.MessageType = messageType
	websocketMessage.MessageData = messageData
	postData, _ := json.Marshal(websocketMessage)

	configModel, result := appconfigsdomain.GetAppConfig("justasking", "RealTimeHubBaseUrl")
	if result.IsSuccess() {
		println(configModel.ConfigValue)
		realtimeUrl := fmt.Sprintf("%v/broadcast/%v", configModel.ConfigValue, code)
		req, err := http.NewRequest("POST", realtimeUrl, bytes.NewBuffer(postData))
		if err != nil {
			result.Status = operationresult.Error
			result.Message = fmt.Sprintf("Error creating request object to send realtimehub for SMS. Error: [%v]", err.Error())
			result.Error = err
		} else {
			client := getHTTPClient()
			_, err := client.Do(req)
			if err != nil {
				result.Status = operationresult.Error
				result.Message = err.Error()
				result.Error = err
			}
		}
	}

	return result
}

func getHTTPClient() http.Client {
	var client = &http.Client{
		Timeout: time.Second * 20,
	}

	return *client
}
