package smsdomain

import (
	"encoding/json"
	"fmt"
	"regexp"
	"justasking/GO/common/clients/realtimehub"
	"justasking/GO/common/enum/boxtype"
	"justasking/GO/common/operationresult"
	"justasking/GO/core/domain/applogs"
	"justasking/GO/core/domain/boxes/answerbox"
	"justasking/GO/core/domain/boxes/basebox"
	"justasking/GO/core/domain/boxes/questionbox"
	"justasking/GO/core/domain/boxes/votesbox"
	"justasking/GO/core/domain/boxes/wordcloud"
	"justasking/GO/core/model/answerboxentry"
	"justasking/GO/core/model/questionboxentry"
	"justasking/GO/core/model/twilio/sms"
	"justasking/GO/core/model/votesboxquestionanswer"
	"justasking/GO/core/model/wordcloudresponse"
	"justasking/GO/core/repo/sms"
	"strconv"
)

var domainName = "SmsDomain"

// HandleSms takes an sms message and passes it to the correct box function
func HandleSms(smsMessage smsmodel.Sms) *operationresult.OperationResult {
	functionName := "HandleSms"
	result := operationresult.New()

	// 1. Log Sms data with Box Id (if it exists)
	logResult := InsertSmsLog(smsMessage)

	if(len(smsMessage.Body)> 0){
		if logResult.IsSuccess() {
			// 2. Look up box associated with this number
			baseBox, result := baseboxdomain.GetBaseBoxByPhoneNumber(smsMessage.To)
			if result.IsSuccess() {
				// 3. Depending on BoxTypeId, handle message body and createdBy
				if baseBox.BoxTypeId == boxtypeenum.WORDCLOUD {
					var wordcloudResponse wordcloudresponsemodel.WordCloudResponse
					wordcloudResponse.Response = smsMessage.Body
					wordcloudResponse.BoxId = baseBox.ID
					wordcloudResponse.CreatedBy = smsMessage.From
					entry, result := wordclouddomain.InsertWordCloudResponse(wordcloudResponse)
					if result.IsSuccess() {
						marshalledStruct, err := json.Marshal(entry)
						if err != nil {
							msg := err.Error()
							result = operationresult.CreateErrorResult(msg, err)
							applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error creating json for wordcloud sms entry. Error: [%v]", msg), false)
						} else {
							broadcastResponse := realtimehubclient.BroadcastSmsMessage(baseBox.Code, "WordCloudResponse", string(marshalledStruct))
							if !broadcastResponse.IsSuccess() {
								result = broadcastResponse
								applogsdomain.LogError("Error sending wordcloud entry to realtimehub for SMS. Entry: [%v]. Error: [%v]", entry.Response, broadcastResponse.Message, false)
							}
						}
					} else {
						msg := result.Error.Error()
						result = operationresult.CreateErrorResult(msg, result.Error)
						applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error inserting word cloud response for sms. SMS From: [%v] SMS To: [%v]. SMS Message: [%v]. Error: [%v]", smsMessage.From, smsMessage.To, smsMessage.Body, msg), false)
					}
				} else if baseBox.BoxTypeId == boxtypeenum.ANSWERBOX {
					activeQuestion, questionResult := answerboxdomain.GetActiveQuestions(baseBox.ID)
					if questionResult.IsSuccess() {
						var entriesArray []answerboxentrymodel.AnswerBoxEntry
						var answerboxEntry answerboxentrymodel.AnswerBoxEntry
						answerboxEntry.Entry = smsMessage.Body
						answerboxEntry.QuestionId = activeQuestion.QuestionId
						answerboxEntry.CreatedBy = smsMessage.From
						entry, result := answerboxdomain.InsertAnswerBoxEntry(answerboxEntry)
						if result.IsSuccess() {
							answerboxEntry.EntryId = entry.EntryId
							entriesArray = append(entriesArray, answerboxEntry)
							marshalledStruct, err := json.Marshal(entriesArray)
							if err != nil {
								msg := err.Error()
								result = operationresult.CreateErrorResult(msg, err)
								applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error creating json for answerbox sms entry. Error: [%v]", msg), false)
							} else {
								broadcastResponse := realtimehubclient.BroadcastSmsMessage(baseBox.Code, "AnswerBoxEntry", string(marshalledStruct))
								if !broadcastResponse.IsSuccess() {
									result = broadcastResponse
									applogsdomain.LogError("Error sending answerbox entry to realtimehub for SMS. Entry: [%v]. Error: [%v]", entry.Entry, broadcastResponse.Message, false)
								}
							}
						} else {
							msg := result.Error.Error()
							result = operationresult.CreateErrorResult(msg, result.Error)
							applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error submitting AnswerBoxEntry for sms. SMS From: [%v] SMS To: [%v]. SMS Message: [%v]. Error: [%v]", smsMessage.From, smsMessage.To, smsMessage.Body, msg), false)
						}
					} else {
						msg := questionResult.Error.Error()
						result = operationresult.CreateErrorResult(msg, questionResult.Error)
						applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting active question for sms. BoxId: [%v]. Error: [%v]", baseBox.ID, msg), false)
					}
				} else if baseBox.BoxTypeId == boxtypeenum.QUESTIONBOX {
					var questionboxEntry questionboxentrymodel.QuestionBoxEntry
					questionboxEntry.Question = smsMessage.Body
					questionboxEntry.BoxId = baseBox.ID
					questionboxEntry.Upvotes = 1
					questionboxEntry.CreatedBy = smsMessage.From
					entry, result := questionboxdomain.InsertQuestionBoxEntry(questionboxEntry)
					if result.IsSuccess() {
						marshalledQuestionBoxEntry, err := json.Marshal(entry)
						if err != nil {
							msg := err.Error()
							result = operationresult.CreateErrorResult(msg, err)
							applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error creating json for questionbox sms entry. Error: [%v]", msg), false)
						} else {
							broadcastResponse := realtimehubclient.BroadcastSmsMessage(baseBox.Code, "QuestionBoxEntry", string(marshalledQuestionBoxEntry))
							if !broadcastResponse.IsSuccess() {
								result = broadcastResponse
								applogsdomain.LogError("Error sending questionbox entry to realtimehub for SMS. Entry: [%v]. Error: [%v]", entry.Question, broadcastResponse.Message, false)
							}
						}
					} else {
						msg := result.Error.Error()
						result = operationresult.CreateErrorResult(msg, result.Error)
						applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error submitting question box entry for sms. SMS From: [%v] SMS To: [%v]. SMS Message: [%v]. Error: [%v]", smsMessage.From, smsMessage.To, smsMessage.Body, msg), false)
					}
				} else if baseBox.BoxTypeId == boxtypeenum.VOTESBOX {
					activeQuestion, questionResult := votesboxdomain.GetActiveQuestions(baseBox.ID)
					if questionResult.IsSuccess() {
						r, _ := regexp.Compile("\\d+")
						entryNumber := r.FindString(smsMessage.Body)
						selectedAnswer, err := strconv.Atoi(entryNumber)
						if err != nil {
							msg := err.Error()
							result = operationresult.CreateErrorResult(msg, err)
							applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Unable to get VotesBox answer from SMS body [%v]. Error: [%v]", smsMessage.Body, msg), false)
						} else {
							var votesArray []votesboxquestionanswermodel.VotesBoxQuestionAnswer
							var votesBoxVote votesboxquestionanswermodel.VotesBoxQuestionAnswer
							found := false

							// Get answerId of selected answer
							for _, answer := range activeQuestion.Answers {
								if answer.SortOrder == selectedAnswer {
									votesBoxVote.AnswerId = answer.AnswerId
									votesBoxVote.QuestionId = activeQuestion.QuestionId
									votesBoxVote.CreatedBy = smsMessage.From
									found = true
								}
							}

							if found {
								_, result := votesboxdomain.InsertVoteNamed(votesBoxVote)
								if result.IsSuccess() {
									votesArray = append(votesArray, votesBoxVote)
									marshalledStruct, err := json.Marshal(votesArray)
									if err != nil {
										msg := err.Error()
										result = operationresult.CreateErrorResult(msg, err)
										applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error creating json for votesbox sms entry. Error: [%v]", msg), false)
									} else {
										broadcastResponse := realtimehubclient.BroadcastSmsMessage(baseBox.Code, "VotesBoxEntrySubmit", string(marshalledStruct))
										if !broadcastResponse.IsSuccess() {
											result = broadcastResponse
											applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error sending VotesBoxVote to realtimehub for SMS. Error: [%v]", broadcastResponse.Message), false)
										}
									}
								} else {
									msg := result.Error.Error()
									result = operationresult.CreateErrorResult(msg, result.Error)
									applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error submitting fordVotesBoxVote for sms. SMS From: [%v] SMS To: [%v]. SMS Message: [%v]. Error: [%v]", smsMessage.From, smsMessage.To, smsMessage.Body, msg), false)
								}
							} else {
								msg := result.Error.Error()
								result = operationresult.CreateErrorResult(msg, result.Error)
								applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error adding VotesBoxVote. Could not map sms body to answer. SMS From: [%v] SMS To: [%v]. SMS Message: [%v]. Error: [%v]", smsMessage.From, smsMessage.To, smsMessage.Body, msg), false)
							}
						}
					} else {
						msg := questionResult.Error.Error()
						result = operationresult.CreateErrorResult(msg, questionResult.Error)
						applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting active question for sms. BoxId: [%v]. Error: [%v]", baseBox.ID, msg), false)
					}
				}
			} else {
				msg := result.Error.Error()
				result = operationresult.CreateErrorResult(msg, result.Error)
				applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting BaseBox for sms entry. Error: [%v]", msg), false)
			}
		} else {
			msg := logResult.Error.Error()
			result = operationresult.CreateErrorResult(msg, logResult.Error)
			applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error creating log for SMS message. Error: [%v]", msg), false)
		}
	} else {
		msg := "smsMessage.Body was null.";
		result.Message = msg;
		result.Status = operationresult.UnprocessableEntity;
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("SMS message cannot be null. Error: [%v]", msg), false)
	}

	return result
}

// InsertSmsLog logs an sms message
func InsertSmsLog(smsMessage smsmodel.Sms) *operationresult.OperationResult {
	functionName := "InsertSmsLog"
	result := operationresult.New()

	err := smsrepo.InsertSmsLog(smsMessage)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error creating log for SMS message. Error: [%v]", msg), false)
	}

	return result
}
