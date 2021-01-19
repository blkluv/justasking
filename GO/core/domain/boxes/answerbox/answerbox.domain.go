package answerboxdomain

import (
	"errors"
	"fmt"
	"justasking/GO/common/operationresult"
	"justasking/GO/core/domain/account"
	"justasking/GO/core/domain/applogs"
	"justasking/GO/core/model/answerboxentry"
	"justasking/GO/core/model/answerboxquestion"
	"justasking/GO/core/model/boxes/answerbox"
	"justasking/GO/core/repo/boxes/answerbox"

	uuid "github.com/satori/go.uuid"
)

var domainName = "AnswerBoxDomain"

// CreateAnswerBox creates an answer box
func CreateAnswerBox(answerBox answerboxmodel.AnswerBox) (uuid.UUID, *operationresult.OperationResult) {
	functionName := "CreateAnswerBox"
	result := operationresult.New()

	// Create ID that will be used for both basebox and QuestionBox
	boxID, _ := uuid.NewV4()
	answerBox.BoxId = boxID
	answerBox.BaseBox.CreatedBy = answerBox.CreatedBy

	err := answerboxrepo.InsertAnswerBox(answerBox)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error creating Answer Box. Error: [%v]", msg), false)
	} else {
		applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("Answer Box created with code [%v] and id [%v]", answerBox.BaseBox.Code, answerBox.BoxId))
	}

	return boxID, result
}

// GetAnswerBoxByBoxId returns an answer box
func GetAnswerBoxByBoxId(guid uuid.UUID) (answerboxmodel.AnswerBox, *operationresult.OperationResult) {
	functionName := "GetAnswerBoxByBoxId"
	result := operationresult.New()

	answerBox, err := answerboxrepo.GetAnswerBoxByBoxId(guid)
	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting answer box with id [%v]. Error: [%v]", guid, msg), false)
	}

	return answerBox, result
}

// GetAnswerBoxByCode returns an answer box
func GetAnswerBoxByCode(code string) (answerboxmodel.AnswerBox, *operationresult.OperationResult) {
	functionName := "GetAnswerBoxByCode"
	result := operationresult.New()

	answerBox, err := answerboxrepo.GetAnswerBoxByBoxCode(code)
	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting answer box with code [%v]. Error: [%v]", code, msg), false)
	}

	return answerBox, result
}

// InsertAnswerBoxEntry inserts an entry to the question box entries table
func InsertAnswerBoxEntry(entry answerboxentrymodel.AnswerBoxEntry) (answerboxentrymodel.AnswerBoxEntry, *operationresult.OperationResult) {
	functionName := "InsertAnswerBoxEntry"
	result := operationresult.New()

	entryID, _ := uuid.NewV4()
	entry.EntryId = entryID

	err := answerboxrepo.InsertAnswerBoxEntry(entry)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error inserting entry for question [%v]. Error: [%v]", entry.QuestionId, msg), false)
	} else {
		applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("Entry [%v] added to question [%v] by user [%v]", entry.Entry, entry.QuestionId, entry.CreatedBy))
	}

	return entry, result
}

// GetAnswerBoxEntriesByBoxId gets all answers for a specific box
func GetAnswerBoxEntriesByBoxId(boxId uuid.UUID) ([]answerboxentrymodel.AnswerBoxEntry, *operationresult.OperationResult) {
	functionName := "GetAnswerBoxEntriesByBoxId"
	result := operationresult.New()

	responses, err := answerboxrepo.GetAnswerBoxEntriesByBoxId(boxId)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting entries for answer box [%v]. Error: [%v]", boxId, msg), false)
	}

	return responses, result
}

// GetAnswerBoxEntriesByCode gets all answers for a specific box
func GetAnswerBoxEntriesByCode(code string) ([]answerboxentrymodel.AnswerBoxEntry, *operationresult.OperationResult) {
	functionName := "GetAnswerBoxEntriesByCode"
	result := operationresult.New()

	responses, err := answerboxrepo.GetAnswerBoxEntriesByCode(code)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting entries for answer box [%v]. Error: [%v]", code, msg), false)
	}

	return responses, result
}

// GetVisibleAnswerBoxEntriesByCode gets all answers for a specific box
func GetVisibleAnswerBoxEntriesByCode(code string) ([]answerboxentrymodel.AnswerBoxEntry, *operationresult.OperationResult) {
	functionName := "GetVisibleAnswerBoxEntriesByCode"
	result := operationresult.New()

	responses, err := answerboxrepo.GetVisibleAnswerBoxEntriesByCode(code)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting entries for answer box [%v]. Error: [%v]", code, msg), false)
	}

	return responses, result
}

// ActivateQuestion activates the given question for an answer box
func ActivateQuestion(answerBoxQuestion answerboxquestionmodel.AnswerBoxQuestion, updatedBy string) *operationresult.OperationResult {
	functionName := "ActivateQuestion"
	result := operationresult.New()

	userId, err := uuid.FromString(updatedBy)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting userId from string [%v]. Error: [%v]", updatedBy, msg), false)
	} else {
		userCanUpdate := userCanUpdateQuestion(userId, answerBoxQuestion.QuestionId)
		if userCanUpdate {
			err := answerboxrepo.ActivateQuestion(answerBoxQuestion, updatedBy)

			if err != nil {
				msg := err.Error()
				result = operationresult.CreateErrorResult(msg, err)
				applogsdomain.LogError(domainName, functionName,
					fmt.Sprintf("Error activating question [%v] from answer box [%v]. Error: [%v]", answerBoxQuestion.QuestionId, answerBoxQuestion.BoxId, msg), false)
			}
		} else {
			msg := fmt.Sprintf("User [%v] attempted to activate answer box question [%v], which belongs to a different account.", updatedBy, answerBoxQuestion.QuestionId)
			result.Message = msg
			result.Status = operationresult.Error
			applogsdomain.LogError(domainName, functionName, msg, false)
		}
	}

	return result
}

// DeactivateQuestion activates the given question for an answer box
func DeactivateQuestion(answerBoxQuestion answerboxquestionmodel.AnswerBoxQuestion, updatedBy string) *operationresult.OperationResult {
	functionName := "DeactivateQuestion"
	result := operationresult.New()

	userId, err := uuid.FromString(updatedBy)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting userId from string [%v]. Error: [%v]", updatedBy, msg), false)
	} else {
		userCanUpdate := userCanUpdateQuestion(userId, answerBoxQuestion.QuestionId)
		if userCanUpdate {
			err := answerboxrepo.DeactivateQuestion(answerBoxQuestion, updatedBy)

			if err != nil {
				msg := err.Error()
				result = operationresult.CreateErrorResult(msg, err)
				applogsdomain.LogError(domainName, functionName,
					fmt.Sprintf("Error deactivating question [%v] from answer box [%v]. Error: [%v]", answerBoxQuestion.QuestionId, answerBoxQuestion.BoxId, msg), false)
			}
		} else {
			msg := fmt.Sprintf("User [%v] attempted to deactivate answer box question [%v], which belongs to a different account.", updatedBy, answerBoxQuestion.QuestionId)
			result.Message = msg
			result.Status = operationresult.Error
			applogsdomain.LogError(domainName, functionName, msg, false)
		}
	}

	return result
}

// HideEntry hides the given answer for a given box
func HideEntry(answerBoxEntry answerboxentrymodel.AnswerBoxEntry, updatedBy string) *operationresult.OperationResult {
	functionName := "HideEntry"
	result := operationresult.New()

	userId, err := uuid.FromString(updatedBy)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting userId from string [%v]. Error: [%v]", updatedBy, msg), false)
	} else {
		userCanUpdate := userCanUpdateEntry(userId, answerBoxEntry.EntryId)
		if userCanUpdate {
			err := answerboxrepo.HideEntry(answerBoxEntry, updatedBy)

			if err != nil {
				msg := err.Error()
				result = operationresult.CreateErrorResult(msg, err)
				applogsdomain.LogError(domainName, functionName,
					fmt.Sprintf("Error hiding entry [%v] from answer box question [%v]. Error: [%v]", answerBoxEntry.Entry, answerBoxEntry.QuestionId, msg), false)
			}
		} else {
			msg := fmt.Sprintf("User [%v] attempted to hide answer box entry for question [%v], which belongs to a different account.", updatedBy, answerBoxEntry.QuestionId)
			result.Message = msg
			result.Status = operationresult.Error
			applogsdomain.LogError(domainName, functionName, msg, false)
		}
	}

	return result
}

// UnhideEntry hides the given answer for a given box
func UnhideEntry(answerBoxEntry answerboxentrymodel.AnswerBoxEntry, updatedBy string) *operationresult.OperationResult {
	functionName := "UnhideEntry"
	result := operationresult.New()

	userId, err := uuid.FromString(updatedBy)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting userId from string [%v]. Error: [%v]", updatedBy, msg), false)
	} else {
		userCanUpdate := userCanUpdateEntry(userId, answerBoxEntry.EntryId)
		if userCanUpdate {
			err := answerboxrepo.UnhideEntry(answerBoxEntry, updatedBy)

			if err != nil {
				msg := err.Error()
				result = operationresult.CreateErrorResult(msg, err)
				applogsdomain.LogError(domainName, functionName,
					fmt.Sprintf("Error hiding entry [%v] from answer box question [%v]. Error: [%v]", answerBoxEntry.Entry, answerBoxEntry.QuestionId, msg), false)
			}
		} else {
			msg := fmt.Sprintf("User [%v] attempted to unhide answer box entry for question [%v], which belongs to a different account.", updatedBy, answerBoxEntry.QuestionId)
			result.Message = msg
			result.Status = operationresult.Error
			applogsdomain.LogError(domainName, functionName, msg, false)
		}
	}

	return result
}

// GetActiveQuestions gets active questions for an answer box. For use with sms.
func GetActiveQuestions(boxId uuid.UUID) (answerboxquestionmodel.AnswerBoxQuestion, *operationresult.OperationResult) {
	functionName := "GetActiveQuestion"
	result := operationresult.New()
	var answerboxQuestion answerboxquestionmodel.AnswerBoxQuestion

	response, err := answerboxrepo.GetActiveQuestions(boxId)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting active question for answer box [%v]. Error: [%v]", boxId, msg), false)
	} else {
		if len(response) == 1 {
			answerboxQuestion = response[0]
		} else {
			err = errors.New("too many answer box questions returned")
			msg := err.Error()
			result = operationresult.CreateErrorResult(msg, err)
			applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("Attempted to get active question for box with more than or less than one active question. Box: [%v]", boxId))
		}
	}

	return answerboxQuestion, result
}

// HideEntriesForQuestion hides the given answer for a question
func HideEntriesForQuestion(answerBoxQuestion answerboxquestionmodel.AnswerBoxQuestion, updatedBy string) *operationresult.OperationResult {
	functionName := "HideEntriesForQuestion"
	result := operationresult.New()

	userId, err := uuid.FromString(updatedBy)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting userId from string [%v]. Error: [%v]", updatedBy, msg), false)
	} else {
		userCanUpdate := userCanUpdateQuestion(userId, answerBoxQuestion.QuestionId)
		if userCanUpdate {

			err := answerboxrepo.HideEntriesForQuestion(answerBoxQuestion, updatedBy)

			if err != nil {
				msg := err.Error()
				result = operationresult.CreateErrorResult(msg, err)
				applogsdomain.LogError(domainName, functionName,
					fmt.Sprintf("Error hiding entries for answer box question [%v]. Error: [%v]", answerBoxQuestion.QuestionId, msg), false)
			}
		} else {
			msg := fmt.Sprintf("User [%v] attempted to hide entries for answer box question [%v], which belongs to a different account.", updatedBy, answerBoxQuestion.QuestionId)
			result.Message = msg
			result.Status = operationresult.Error
			applogsdomain.LogError(domainName, functionName, msg, false)
		}
	}

	return result
}

// UnhideEntriesForQuestion hides the given answer for a given question
func UnhideEntriesForQuestion(answerBoxQuestion answerboxquestionmodel.AnswerBoxQuestion, updatedBy string) *operationresult.OperationResult {
	functionName := "UnhideEntriesForQuestion"
	result := operationresult.New()

	userId, err := uuid.FromString(updatedBy)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting userId from string [%v]. Error: [%v]", updatedBy, msg), false)
	} else {
		userCanUpdate := userCanUpdateQuestion(userId, answerBoxQuestion.QuestionId)
		if userCanUpdate {

			err := answerboxrepo.UnhideEntriesForQuestion(answerBoxQuestion, updatedBy)

			if err != nil {
				msg := err.Error()
				result = operationresult.CreateErrorResult(msg, err)
				applogsdomain.LogError(domainName, functionName,
					fmt.Sprintf("Error hiding entries for answer box question [%v]. Error: [%v]", answerBoxQuestion.QuestionId, msg), false)
			}
		} else {
			msg := fmt.Sprintf("User [%v] attempted to hide entries for answer box question [%v], which belongs to a different account.", updatedBy, answerBoxQuestion.QuestionId)
			result.Message = msg
			result.Status = operationresult.Error
			applogsdomain.LogError(domainName, functionName, msg, false)
		}
	}

	return result
}

// HideAllEntries hides the given answer for a given box
func HideAllEntries(answerBox answerboxmodel.AnswerBox, updatedBy string) *operationresult.OperationResult {
	functionName := "HideAllEntries"
	result := operationresult.New()

	err := answerboxrepo.HideAllEntries(answerBox, updatedBy)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error hiding all entries for answer box [%v]. Error: [%v]", answerBox.BoxId, msg), false)
	}

	return result
}

// UnhideAllEntries hides the given answer for a given box
func UnhideAllEntries(answerBox answerboxmodel.AnswerBox, updatedBy string) *operationresult.OperationResult {
	functionName := "UnhideAllEntries"
	result := operationresult.New()

	err := answerboxrepo.UnhideAllEntries(answerBox, updatedBy)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName,
			fmt.Sprintf("Error unhiding all entries for answer box [%v]. Error: [%v]", answerBox.BoxId, msg), false)
	}

	return result
}

func userCanUpdateEntry(userId uuid.UUID, entryId uuid.UUID) bool {

	userCanUpdateEntry := false
	account, accountResponse := accountdomain.GetAccountByAnswerBoxEntry(entryId)
	if accountResponse.IsSuccess() {
		userBelongsToAccount := accountdomain.UserBelongsToAccount(userId, account.Id)
		if userBelongsToAccount {
			userCanUpdateEntry = true
		}
	}

	return userCanUpdateEntry
}

func userCanUpdateQuestion(userId uuid.UUID, questionId uuid.UUID) bool {

	userCanUpdateEntry := false
	account, accountResponse := accountdomain.GetAccountByAnswerBoxQuestion(questionId)
	if accountResponse.IsSuccess() {
		userBelongsToAccount := accountdomain.UserBelongsToAccount(userId, account.Id)
		if userBelongsToAccount {
			userCanUpdateEntry = true
		}
	}

	return userCanUpdateEntry
}
