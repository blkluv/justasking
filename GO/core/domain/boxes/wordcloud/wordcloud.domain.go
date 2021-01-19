package wordclouddomain

import (
	"fmt"
	"justasking/GO/common/operationresult"
	"justasking/GO/core/domain/applogs"
	"justasking/GO/core/domain/boxes/basebox"
	"justasking/GO/core/model/boxes/wordcloud"
	"justasking/GO/core/model/responses"
	"justasking/GO/core/model/wordcloudresponse"
	"justasking/GO/core/repo/boxes/wordcloud"

	uuid "github.com/satori/go.uuid"
)

var domainName = "WordCloudDomain"

// CreateWordCloud creates a word cloud
func CreateWordCloud(wordCloud wordcloudmodel.WordCloud) (uuid.UUID, *operationresult.OperationResult) {
	functionName := "CreateWordCloud"
	result := operationresult.New()

	// Create ID that will be used for both basebox and WordCloud
	boxID, _ := uuid.NewV4()
	wordCloud.BoxId = boxID
	wordCloud.BaseBox.CreatedBy = wordCloud.CreatedBy

	err := wordcloudrepo.InsertWordCloud(wordCloud)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error creating word cloud. Error: [%v]", msg), false)
	} else {
		applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("Word Cloud created with code [%v] and id [%v]", wordCloud.BaseBox.Code, wordCloud.BoxId))
	}

	return boxID, result
}

// GetWordCloudByBoxId returns a word cloud
func GetWordCloudByBoxId(guid uuid.UUID) (wordcloudmodel.WordCloud, *operationresult.OperationResult) {
	functionName := "GetWordCloudByBoxId"
	result := operationresult.New()

	wordCloud, err := wordcloudrepo.GetWordCloudByBoxId(guid)
	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting word cloud with id [%v]. Error: [%v]", guid, msg), false)
	}

	return wordCloud, result
}

// GetWordCloudByCode returns a word cloud
func GetWordCloudByCode(code string) (wordcloudmodel.WordCloud, *operationresult.OperationResult) {
	functionName := "GetWordCloudByCode"
	result := operationresult.New()

	wordCloud, err := wordcloudrepo.GetWordCloudByCode(code)
	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting word cloud with code [%v]. Error: [%v]", code, msg), false)
	}

	return wordCloud, result
}

// InsertWordCloudResponse adds a response to a word cloud
func InsertWordCloudResponse(response wordcloudresponsemodel.WordCloudResponse) (wordcloudresponsemodel.WordCloudResponse, *operationresult.OperationResult) {
	functionName := "InsertWordCloudResponse"
	result := operationresult.New()
	var entry wordcloudresponsemodel.WordCloudResponse
	var err error

	entry, err = wordcloudrepo.InsertWordCloudResponse(response)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error inserting response to word cloud [%v]. Error: [%v]", response.BoxId, msg), false)
	} else {
		applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("Response [%v] added to word cloud [%v] by user [%v]", response.Response, response.BoxId, response.CreatedBy))
	}

	return entry, result
}

// GetFullWordCloudResponsesByBoxId gets all answers for a specific box
func GetFullWordCloudResponsesByBoxId(boxId uuid.UUID) ([]wordcloudresponsemodel.WordCloudResponse, *operationresult.OperationResult) {
	functionName := "GetFullWordCloudResponsesByBoxId"
	result := operationresult.New()

	responses, err := wordcloudrepo.GetWordCloudResponsesByBoxId(boxId)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting responses from word cloud [%v]. Error: [%v]", boxId, msg), false)
	}

	return responses, result
}

// GetPartialWordCloudResponsesByBoxId gets all answers for a specific box
func GetPartialWordCloudResponsesByBoxId(boxId uuid.UUID) ([]map[string]bool, *operationresult.OperationResult) {
	functionName := "GetPartialWordCloudResponsesByBoxId"
	result := operationresult.New()
	var answers []map[string]bool

	responses, err := wordcloudrepo.GetWordCloudResponsesByBoxId(boxId)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting responses from word cloud [%v]. Error: [%v]", boxId, msg), false)
	} else {
		for _, response := range responses {
			answer := make(map[string]bool)
			answer[response.Response] = response.IsHidden
			answers = append(answers, answer)
		}
	}

	return answers, result
}

// GetFullWordCloudResponsesByCode gets all answers for a specific box
func GetFullWordCloudResponsesByCode(code string) ([]wordcloudresponsemodel.WordCloudResponse, *operationresult.OperationResult) {
	functionName := "GetFullWordCloudResponsesByCode"
	result := operationresult.New()

	responses, err := wordcloudrepo.GetWordCloudResponsesByCode(code)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting responses from word cloud [%v]. Error: [%v]", code, msg), false)
	}

	return responses, result
}

// GetPartialWordCloudResponsesByCode gets all answers for a specific box
func GetPartialWordCloudResponsesByCode(code string) ([]partialresponse.PartialResponse, *operationresult.OperationResult) {
	functionName := "GetPartialWordCloudResponsesByCode"
	result := operationresult.New()
	var answers []partialresponse.PartialResponse

	responses, err := wordcloudrepo.GetWordCloudResponsesByCode(code)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting responses from word cloud [%v]. Error: [%v]", code, msg), false)
	} else {
		for _, response := range responses {
			answer := *partialresponse.New()
			answer.Response = response.Response
			answer.Hidden = response.IsHidden
			answers = append(answers, answer)
		}
	}

	return answers, result
}

// HideAnswer hides the given answer for a given box
func HideAnswer(boxId uuid.UUID, answerToHide string, updatedBy string) (string, *operationresult.OperationResult) {
	functionName := "HideAnswer"
	result := operationresult.New()

	userId, err := uuid.FromString(updatedBy)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting userId from string [%v]. Error: [%v]", updatedBy, msg), false)
	} else {
		userCanUpdate := baseboxdomain.UserCanUpdateBox(userId, boxId)
		if userCanUpdate {
			err := wordcloudrepo.HideAnswer(boxId, answerToHide, updatedBy)

			if err != nil {
				msg := err.Error()
				result = operationresult.CreateErrorResult(msg, err)
				applogsdomain.LogError(domainName, functionName,
					fmt.Sprintf("Error hiding response [%v] from word cloud [%v]. Error: [%v]", answerToHide, boxId, msg), false)
			}
		} else {
			msg := fmt.Sprintf("User [%v] attempted to hide entry [%v] for word cloud [%v], which belongs to a different account.", updatedBy, answerToHide, boxId)
			result.Message = msg
			result.Status = operationresult.Error
			applogsdomain.LogError(domainName, functionName, msg, false)
		}
	}

	return answerToHide, result
}

// UnhideAnswer hides the given answer for a given box
func UnhideAnswer(boxId uuid.UUID, answerToHide string, updatedBy string) (string, *operationresult.OperationResult) {
	functionName := "UnhideAnswer"
	result := operationresult.New()

	userId, err := uuid.FromString(updatedBy)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting userId from string [%v]. Error: [%v]", updatedBy, msg), false)
	} else {
		userCanUpdate := baseboxdomain.UserCanUpdateBox(userId, boxId)
		if userCanUpdate {
			err := wordcloudrepo.UnhideAnswer(boxId, answerToHide, updatedBy)

			if err != nil {
				msg := err.Error()
				result = operationresult.CreateErrorResult(msg, err)
				applogsdomain.LogError(domainName, functionName,
					fmt.Sprintf("Error hiding response [%v] from word cloud [%v]. Error: [%v]", answerToHide, boxId, msg), false)
			}
		} else {
			msg := fmt.Sprintf("User [%v] attempted to unhide entry [%v] for word cloud [%v], which belongs to a different account.", updatedBy, answerToHide, boxId)
			result.Message = msg
			result.Status = operationresult.Error
			applogsdomain.LogError(domainName, functionName, msg, false)
		}
	}

	return answerToHide, result
}

// HideAllAnswers hides the given answer for a given box
func HideAllAnswers(boxId uuid.UUID, updatedBy string) *operationresult.OperationResult {
	functionName := "HideAllAnswers"
	result := operationresult.New()

	userId, err := uuid.FromString(updatedBy)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting userId from string [%v]. Error: [%v]", updatedBy, msg), false)
	} else {
		userCanUpdate := baseboxdomain.UserCanUpdateBox(userId, boxId)
		if userCanUpdate {
			err := wordcloudrepo.HideAllAnswers(boxId, updatedBy)

			if err != nil {
				msg := err.Error()
				result = operationresult.CreateErrorResult(msg, err)
				applogsdomain.LogError(domainName, functionName,
					fmt.Sprintf("Error hiding all responses from word cloud [%v]. Error: [%v]", boxId, msg), false)
			}
		} else {
			msg := fmt.Sprintf("User [%v] attempted to hide all entries for word cloud [%v], which belongs to a different account.", updatedBy, boxId)
			result.Message = msg
			result.Status = operationresult.Error
			applogsdomain.LogError(domainName, functionName, msg, false)
		}
	}

	return result
}

// UnhideAllAnswers hides the given answer for a given box
func UnhideAllAnswers(boxId uuid.UUID, updatedBy string) *operationresult.OperationResult {
	functionName := "UnhideAllAnswers"
	result := operationresult.New()

	userId, err := uuid.FromString(updatedBy)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting userId from string [%v]. Error: [%v]", updatedBy, msg), false)
	} else {
		userCanUpdate := baseboxdomain.UserCanUpdateBox(userId, boxId)
		if userCanUpdate {
			err := wordcloudrepo.UnhideAllAnswers(boxId, updatedBy)

			if err != nil {
				msg := err.Error()
				result = operationresult.CreateErrorResult(msg, err)
				applogsdomain.LogError(domainName, functionName,
					fmt.Sprintf("Error hiding all responses from word cloud [%v]. Error: [%v]", boxId, msg), false)
			}
		} else {
			msg := fmt.Sprintf("User [%v] attempted to hide all entries for word cloud [%v], which belongs to a different account.", updatedBy, boxId)
			result.Message = msg
			result.Status = operationresult.Error
			applogsdomain.LogError(domainName, functionName, msg, false)
		}
	}

	return result
}
