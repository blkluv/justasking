package questionboxdomain

import (
	"fmt"
	"justasking/GO/common/operationresult"
	"justasking/GO/core/domain/applogs"
	"justasking/GO/core/domain/boxes/basebox"
	"justasking/GO/core/model/boxes/questionbox"
	"justasking/GO/core/model/questionboxentry"
	"justasking/GO/core/model/questionboxentryvote"
	"justasking/GO/core/repo/boxes/questionbox"

	uuid "github.com/satori/go.uuid"
)

var domainName = "QuestionBoxDomain"

// CreateQuestionBox creates a question box
func CreateQuestionBox(questionBox questionboxmodel.QuestionBox) (uuid.UUID, *operationresult.OperationResult) {
	functionName := "CreateQuestionBox"
	result := operationresult.New()

	// Create ID that will be used for both basebox and QuestionBox
	boxID, _ := uuid.NewV4()
	questionBox.BoxId = boxID
	questionBox.BaseBox.CreatedBy = questionBox.CreatedBy

	err := questionboxrepo.InsertQuestionBox(questionBox)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error creating Question Box. Error: [%v]", msg), false)
	} else {
		applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("Question Box created with code [%v] and id [%v]", questionBox.BaseBox.Code, questionBox.BoxId))
	}

	return boxID, result
}

// GetQuestionBoxByBoxId returns a question box
func GetQuestionBoxByBoxId(guid uuid.UUID) (questionboxmodel.QuestionBox, *operationresult.OperationResult) {
	functionName := "GetQuestionBoxByBoxId"
	result := operationresult.New()

	questionBox, err := questionboxrepo.GetQuestionBoxByBoxId(guid)
	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting question box with id [%v]. Error: [%v]", guid, msg), false)
	}

	return questionBox, result
}

// GetQuestionBoxByCode returns a question box
func GetQuestionBoxByCode(code string) (questionboxmodel.QuestionBox, *operationresult.OperationResult) {
	functionName := "GetQuestionBoxByCode"
	result := operationresult.New()

	questionBox, err := questionboxrepo.GetQuestionBoxByBoxCode(code)
	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting question box with code [%v]. Error: [%v]", code, msg), false)
	}

	return questionBox, result
}

// InsertQuestionBoxEntry inserts an entry to the question box entries table
func InsertQuestionBoxEntry(entry questionboxentrymodel.QuestionBoxEntry) (questionboxentrymodel.QuestionBoxEntry, *operationresult.OperationResult) {
	functionName := "InsertQuestionBoxEntry"
	result := operationresult.New()

	entryID, _ := uuid.NewV4()
	entry.EntryId = entryID

	err := questionboxrepo.InsertQuestionBoxEntry(entry)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error inserting entry into question box [%v]. Error: [%v]", entry.BoxId, msg), false)
	} else {
		applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("Entry [%v] added to question box [%v] by user [%v]", entry.Question, entry.BoxId, entry.CreatedBy))

		//entry was added, now give it an upvote
		var entryVote questionboxentryvotemodel.QuestionBoxEntryVote
		entryVote.EntryId = entry.EntryId
		entryVote.VoteType = "upvote"
		entryVote.VoteValue = 1
		entryVote.CreatedBy = entry.CreatedBy
		err = questionboxrepo.InsertQuestionBoxEntryVote(entryVote)

		if err != nil {
			msg := err.Error()
			result = operationresult.CreateErrorResult(msg, err)
			applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error inserting entry vote for  EntryId [%v]. Error: [%v]", entryVote.EntryId, msg), false)
		} else {
			applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("Vote Type [%v] with Vote Value [%v] added for entry [%v] by user [%v] in entry votes table", entryVote.VoteType, entryVote.VoteValue, entryVote.EntryId, entryVote.CreatedBy))
		}
	}

	return entry, result
}

// InsertQuestionBoxEntryVote inserts an entry to the question box entries table
func InsertQuestionBoxEntryVote(entryVote questionboxentryvotemodel.QuestionBoxEntryVote) (questionboxentryvotemodel.QuestionBoxEntryVote, *operationresult.OperationResult) {
	functionName := "InsertQuestionBoxEntryVote"
	result := operationresult.New()

	err := questionboxrepo.InsertQuestionBoxEntryVote(entryVote)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error inserting entry vote for  EntryId [%v]. Error: [%v]", entryVote.EntryId, msg), false)
	} else {
		applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("Vote Type [%v] with Vote Value [%v] added for entry [%v] by user [%v] in entry votes table", entryVote.VoteType, entryVote.VoteValue, entryVote.EntryId, entryVote.CreatedBy))
	}

	return entryVote, result
}

// GetQuestionBoxEntriesByBoxId gets all answers for a specific box
func GetQuestionBoxEntriesByBoxId(boxId uuid.UUID) ([]questionboxentrymodel.QuestionBoxEntry, *operationresult.OperationResult) {
	functionName := "GetQuestionBoxEntriesByBoxId"
	result := operationresult.New()

	responses, err := questionboxrepo.GetQuestionBoxEntriesByBoxId(boxId)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting entries for question box [%v]. Error: [%v]", boxId, msg), false)
	}

	return responses, result
}

// GetQuestionBoxEntriesByCode gets all answers for a specific box
func GetQuestionBoxEntriesByCode(code string) ([]questionboxentrymodel.QuestionBoxEntry, *operationresult.OperationResult) {
	functionName := "GetQuestionBoxEntriesByCode"
	result := operationresult.New()

	responses, err := questionboxrepo.GetQuestionBoxEntriesByCode(code)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting entries for question box [%v]. Error: [%v]", code, msg), false)
	}

	return responses, result
}

// GetVisibleQuestionBoxEntriesByCode gets all answers for a specific box
func GetVisibleQuestionBoxEntriesByCode(code string) ([]questionboxentrymodel.QuestionBoxEntry, *operationresult.OperationResult) {
	functionName := "GetVisibleQuestionBoxEntriesByCode"
	result := operationresult.New()

	responses, err := questionboxrepo.GetVisibleQuestionBoxEntriesByCode(code)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting entries for question box [%v]. Error: [%v]", code, msg), false)
	}

	return responses, result
}

// HideEntry hides the given answer for a given box
func HideEntry(questionBoxEntry questionboxentrymodel.QuestionBoxEntry, updatedBy string) *operationresult.OperationResult {
	functionName := "HideEntry"
	result := operationresult.New()

	userId, err := uuid.FromString(updatedBy)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting userId from string [%v]. Error: [%v]", updatedBy, msg), false)
	} else {
		userCanUpdate := baseboxdomain.UserCanUpdateBox(userId, questionBoxEntry.BoxId)
		if userCanUpdate {
			err := questionboxrepo.HideEntry(questionBoxEntry, updatedBy)

			if err != nil {
				msg := err.Error()
				result = operationresult.CreateErrorResult(msg, err)
				applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error hiding entry [%v] from question box [%v]. Error: [%v]", questionBoxEntry.Question, questionBoxEntry.BoxId, msg), false)
			}
		} else {
			msg := fmt.Sprintf("User [%v] attempted to hide question box entry [%v], which belongs to a different account.", updatedBy, questionBoxEntry.EntryId)
			result.Message = msg
			result.Status = operationresult.Error
			applogsdomain.LogError(domainName, functionName, msg, false)
		}
	}

	return result
}

// UnhideEntry hides the given answer for a given box
func UnhideEntry(questionBoxEntry questionboxentrymodel.QuestionBoxEntry, updatedBy string) *operationresult.OperationResult {
	functionName := "UnhideEntry"
	result := operationresult.New()

	userId, err := uuid.FromString(updatedBy)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting userId from string [%v]. Error: [%v]", updatedBy, msg), false)
	} else {
		userCanUpdate := baseboxdomain.UserCanUpdateBox(userId, questionBoxEntry.BoxId)
		if userCanUpdate {
			err := questionboxrepo.UnhideEntry(questionBoxEntry, updatedBy)

			if err != nil {
				msg := err.Error()
				result = operationresult.CreateErrorResult(msg, err)
				applogsdomain.LogError(domainName, functionName,
					fmt.Sprintf("Error hiding entry [%v] from question box [%v]. Error: [%v]", questionBoxEntry.Question, questionBoxEntry.BoxId, msg), false)
			}
		} else {
			msg := fmt.Sprintf("User [%v] attempted to unhide question box entry [%v], which belongs to a different account.", updatedBy, questionBoxEntry.EntryId)
			result.Message = msg
			result.Status = operationresult.Error
			applogsdomain.LogError(domainName, functionName, msg, false)
		}
	}

	return result
}

// UpvoteFromDownvote cancels out previous downvote by giving it a -1, and adds an upvote
func UpvoteFromDownvote(entryVote questionboxentryvotemodel.QuestionBoxEntryVote) (questionboxentryvotemodel.QuestionBoxEntryVote, *operationresult.OperationResult) {
	functionName := "UpvoteFromDownvote"
	result := operationresult.New()

	err := questionboxrepo.UpvoteFromDownvote(entryVote)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error upvoting from downvote for EntryId [%v]. Error: [%v]", entryVote.EntryId, msg), false)
	} else {
		applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("Upvote from downvote successfully added for entry [%v] by user [%v] in entry votes table", entryVote.VoteType, entryVote.VoteValue))
	}

	return entryVote, result
}

// DownvoteFromUpvote cancels out previous upvote by giving it a -1, and adds a downvote
func DownvoteFromUpvote(entryVote questionboxentryvotemodel.QuestionBoxEntryVote) (questionboxentryvotemodel.QuestionBoxEntryVote, *operationresult.OperationResult) {
	functionName := "DownvoteFromUpvote"
	result := operationresult.New()

	err := questionboxrepo.DownvoteFromUpvote(entryVote)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error downvoting from upvote for EntryId [%v]. Error: [%v]", entryVote.EntryId, msg), false)
	} else {
		applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("Downvote from upvote successfully added for entry [%v] by user [%v] in entry votes table", entryVote.VoteType, entryVote.VoteValue))
	}

	return entryVote, result
}

// UndoUpvote adds a -1 to the upvotes for a question
func UndoUpvote(entryVote questionboxentryvotemodel.QuestionBoxEntryVote) (questionboxentryvotemodel.QuestionBoxEntryVote, *operationresult.OperationResult) {
	functionName := "UndoUpvote"
	result := operationresult.New()

	err := questionboxrepo.UndoUpvote(entryVote)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error reverting upvote for EntryId [%v]. Error: [%v]", entryVote.EntryId, msg), false)
	} else {
		applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("Upvote successfully reverted added for entry [%v] by user [%v] in entry votes table", entryVote.VoteType, entryVote.VoteValue))
	}

	return entryVote, result
}

// UndoDownvote adds a -1 to the downvotes for a question
func UndoDownvote(entryVote questionboxentryvotemodel.QuestionBoxEntryVote) (questionboxentryvotemodel.QuestionBoxEntryVote, *operationresult.OperationResult) {
	functionName := "UndoDownvote"
	result := operationresult.New()

	err := questionboxrepo.UndoDownvote(entryVote)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error reverting downvote for EntryId [%v]. Error: [%v]", entryVote.EntryId, msg), false)
	} else {
		applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("Downvote successfully reverted added for entry [%v] by user [%v] in entry votes table", entryVote.VoteType, entryVote.VoteValue))
	}

	return entryVote, result
}

// HideAllEntries hides the given answer for a given box
func HideAllEntries(questionBox questionboxmodel.QuestionBox, updatedBy string) *operationresult.OperationResult {
	functionName := "HideAllEntries"
	result := operationresult.New()

	userId, err := uuid.FromString(updatedBy)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting userId from string [%v]. Error: [%v]", updatedBy, msg), false)
	} else {
		userCanUpdate := baseboxdomain.UserCanUpdateBox(userId, questionBox.BoxId)
		if userCanUpdate {
			err := questionboxrepo.HideAllEntries(questionBox, updatedBy)

			if err != nil {
				msg := err.Error()
				result = operationresult.CreateErrorResult(msg, err)
				applogsdomain.LogError(domainName, functionName,
					fmt.Sprintf("Error hiding all entries for question box [%v]. Error: [%v]", questionBox.BoxId, msg), false)
			}
		} else {
			msg := fmt.Sprintf("User [%v] attempted to unhide all entries for question box [%v], which belongs to a different account.", updatedBy, questionBox.BoxId)
			result.Message = msg
			result.Status = operationresult.Error
			applogsdomain.LogError(domainName, functionName, msg, false)
		}
	}

	return result
}

// UnhideAllEntries hides the given answer for a given box
func UnhideAllEntries(questionBox questionboxmodel.QuestionBox, updatedBy string) *operationresult.OperationResult {
	functionName := "UnhideAllEntries"
	result := operationresult.New()

	userId, err := uuid.FromString(updatedBy)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting userId from string [%v]. Error: [%v]", updatedBy, msg), false)
	} else {
		userCanUpdate := baseboxdomain.UserCanUpdateBox(userId, questionBox.BoxId)
		if userCanUpdate {
			err := questionboxrepo.UnhideAllEntries(questionBox, updatedBy)

			if err != nil {
				msg := err.Error()
				result = operationresult.CreateErrorResult(msg, err)
				applogsdomain.LogError(domainName, functionName,
					fmt.Sprintf("Error hiding all entries for question box [%v]. Error: [%v]", questionBox.BoxId, msg), false)
			}
		} else {
			msg := fmt.Sprintf("User [%v] attempted to unhide all entries for question box [%v], which belongs to a different account.", updatedBy, questionBox.BoxId)
			result.Message = msg
			result.Status = operationresult.Error
			applogsdomain.LogError(domainName, functionName, msg, false)
		}
	}

	return result
}
