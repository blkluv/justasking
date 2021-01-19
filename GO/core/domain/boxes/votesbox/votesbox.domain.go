package votesboxdomain

import (
	"errors"
	"fmt"
	"justasking/GO/common/operationresult"
	"justasking/GO/core/domain/applogs"
	"justasking/GO/core/domain/boxes/basebox"
	"justasking/GO/core/model/boxes/votesbox"
	"justasking/GO/core/model/votesboxquestion"
	"justasking/GO/core/model/votesboxquestionanswer"
	"justasking/GO/core/repo/boxes/votesbox"

	uuid "github.com/satori/go.uuid"
)

var domainName = "VotesBoxDomain"

// CreateVotesBox creates a votes box
func CreateVotesBox(votesBox votesboxmodel.VotesBox) (uuid.UUID, *operationresult.OperationResult) {
	functionName := "CreateVotesBox"
	result := operationresult.New()

	// Create ID that will be used for both basebox and QuestionBox
	boxID, _ := uuid.NewV4()
	votesBox.BoxId = boxID
	votesBox.BaseBox.CreatedBy = votesBox.CreatedBy

	err := votesboxrepo.InsertVotesBox(votesBox)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error creating Votes Box. Error: [%v]", msg), false)
	} else {
		applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("Votes Box created with code [%v] and id [%v]", votesBox.BaseBox.Code, votesBox.BoxId))
	}

	return boxID, result
}

// GetVotesBoxByBoxId returns a votes box
func GetVotesBoxByBoxId(guid uuid.UUID) (votesboxmodel.VotesBox, *operationresult.OperationResult) {
	functionName := "GetVotesBoxByBoxId"
	result := operationresult.New()

	votesBox, err := votesboxrepo.GetVotesBoxByBoxId(guid)
	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting votes box with id [%v]. Error: [%v]", guid, msg), false)
	}

	return votesBox, result
}

// GetVotesBoxByCode returns a votes box
func GetVotesBoxByCode(code string) (votesboxmodel.VotesBox, *operationresult.OperationResult) {
	functionName := "GetVotesBoxByCode"
	result := operationresult.New()

	votesBox, err := votesboxrepo.GetVotesBoxByBoxCode(code)
	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting votes box with code [%v]. Error: [%v]", code, msg), false)
	}

	return votesBox, result
}

// InsertVoteAnonymous inserts a vote for a given answer
func InsertVoteAnonymous(entryVote votesboxquestionanswermodel.VotesBoxQuestionAnswer) (votesboxquestionanswermodel.VotesBoxQuestionAnswer, *operationresult.OperationResult) {
	functionName := "InsertVoteAnonymous"
	result := operationresult.New()

	err := votesboxrepo.InsertVoteAnonymous(entryVote)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error inserting anonymous entry vote for votes box answerId [%v]. Error: [%v]", entryVote.AnswerId, msg), false)
	} else {
		applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("Vote added for answer [%v] by anonymous user in votes box question answer table", entryVote.AnswerId))
	}

	return entryVote, result
}

// InsertVoteNamed inserts a vote for a given answer
func InsertVoteNamed(entryVote votesboxquestionanswermodel.VotesBoxQuestionAnswer) (votesboxquestionanswermodel.VotesBoxQuestionAnswer, *operationresult.OperationResult) {
	functionName := "InsertVoteNamed"
	result := operationresult.New()

	inserted, err := votesboxrepo.InsertVoteNamed(entryVote)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error inserting named entry vote for votes box answerId [%v]. Error: [%v]", entryVote.AnswerId, msg), false)
	} else if inserted.RecordsInserted != 1 {
		msg := fmt.Sprintf("Duplicate vote attempt by user [%v] for question [%v] and answerid [%v]. Vote not added.", entryVote.CreatedBy, entryVote.QuestionId, entryVote.AnswerId)
		result.Status = operationresult.Error
		result.Error = errors.New(msg)
		applogsdomain.LogInfo(domainName, functionName, msg)
	} else {
		applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("Vote added for answer [%v] by user [%v] in votes box question answer table", entryVote.AnswerId, entryVote.CreatedBy))
	}

	return entryVote, result
}

// GetActiveQuestions gets active questions for a votes box. For use with sms.
func GetActiveQuestions(boxId uuid.UUID) (votesboxquestionmodel.VotesBoxQuestion, *operationresult.OperationResult) {
	functionName := "GetActiveQuestion"
	result := operationresult.New()
	var votesboxQuestion votesboxquestionmodel.VotesBoxQuestion

	response, err := votesboxrepo.GetActiveQuestions(boxId)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting active question for votes box [%v]. Error: [%v]", boxId, msg), false)
	} else {
		if len(response) == 1 {
			votesboxQuestion = response[0]
		} else {
			err = errors.New("too many votes box questions returned")
			msg := err.Error()
			result = operationresult.CreateErrorResult(msg, err)
			applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("Attempted to get active question for box with more than or less than one active question. Box: [%v]", boxId))
		}
	}

	return votesboxQuestion, result
}

// ActivateQuestion activates the given question for a votes box
func ActivateQuestion(votesboxQuestion votesboxquestionmodel.VotesBoxQuestion, updatedBy string) *operationresult.OperationResult {
	functionName := "ActivateQuestion"
	result := operationresult.New()

	userId, err := uuid.FromString(updatedBy)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting userId from string [%v]. Error: [%v]", updatedBy, msg), false)
	} else {
		userCanUpdate := baseboxdomain.UserCanUpdateBox(userId, votesboxQuestion.BoxId)
		if userCanUpdate {
			err := votesboxrepo.ActivateQuestion(votesboxQuestion, updatedBy)

			if err != nil {
				msg := err.Error()
				result = operationresult.CreateErrorResult(msg, err)
				applogsdomain.LogError(domainName, functionName,
					fmt.Sprintf("Error activating question [%v] for votes box [%v]. Error: [%v]", votesboxQuestion.QuestionId, votesboxQuestion.BoxId, msg), false)
			}
		} else {
			msg := fmt.Sprintf("User [%v] attempted to activate question [%v] for votes box [%v], which belongs to a different account.", updatedBy, votesboxQuestion.QuestionId, votesboxQuestion.BoxId)
			result.Message = msg
			result.Status = operationresult.Error
			applogsdomain.LogError(domainName, functionName, msg, false)
		}
	}

	return result
}

// DeactivateQuestion activates the given question for an answer box
func DeactivateQuestion(votesboxQuestion votesboxquestionmodel.VotesBoxQuestion, updatedBy string) *operationresult.OperationResult {
	functionName := "DeactivateQuestion"
	result := operationresult.New()

	userId, err := uuid.FromString(updatedBy)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting userId from string [%v]. Error: [%v]", updatedBy, msg), false)
	} else {
		userCanUpdate := baseboxdomain.UserCanUpdateBox(userId, votesboxQuestion.BoxId)
		if userCanUpdate {
			err := votesboxrepo.DeactivateQuestion(votesboxQuestion, updatedBy)

			if err != nil {
				msg := err.Error()
				result = operationresult.CreateErrorResult(msg, err)
				applogsdomain.LogError(domainName, functionName,
					fmt.Sprintf("Error deactivating question [%v] for votes box [%v]. Error: [%v]", votesboxQuestion.QuestionId, votesboxQuestion.BoxId, msg), false)
			}
		} else {
			msg := fmt.Sprintf("User [%v] attempted to deactivate question [%v] for votes box [%v], which belongs to a different account.", updatedBy, votesboxQuestion.QuestionId, votesboxQuestion.BoxId)
			result.Message = msg
			result.Status = operationresult.Error
			applogsdomain.LogError(domainName, functionName, msg, false)
		}
	}
	return result
}
