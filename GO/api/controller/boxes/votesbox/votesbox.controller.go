package votesboxcontroller

import (
	"encoding/json"
	"justasking/GO/api/startup/middleware"
	"justasking/GO/common/operationresult"
	"justasking/GO/core/domain/boxes/votesbox"
	"justasking/GO/core/model/boxes/votesbox"
	"justasking/GO/core/model/votesboxquestion"
	"justasking/GO/core/model/votesboxquestionanswer"
	"justasking/GO/core/startup/flight"
	"net/http"

	"github.com/blue-jay/core/router"
	uuid "github.com/satori/go.uuid"
)

// Load the routes.
func Load() {
	router.Post("/votesbox", CreateVotesBox, middleware.AuthorizedHandler)
	router.Post("/votesbox/vote/", InsertVoteAnonymous)
	router.Put("/votesbox/activate/question/", ActivateQuestion, middleware.AuthorizedHandler)
	router.Put("/votesbox/deactivate/question/", DeactivateQuestion, middleware.AuthorizedHandler)
	router.Get("/votesbox/id/:boxid", GetVotesBoxByBoxId)
	router.Get("/votesbox/code/:code", GetVotesBoxByCode)
}

// CreateVotesBox creates an Answer Box
func CreateVotesBox(w http.ResponseWriter, r *http.Request) {
	var votesBox votesboxmodel.VotesBox

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&votesBox)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		tokenClaims := middleware.GetTokenClaims(r)
		votesBox.CreatedBy = tokenClaims.ID.String()
		votesBox.BaseBox.AccountId = tokenClaims.Account.Id
		boxId, result := votesboxdomain.CreateVotesBox(votesBox)
		if result.IsSuccess() {
			responseString, err := json.Marshal(boxId)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				w.Write([]byte(responseString))
			}
		} else if result.Status == operationresult.Forbidden {
			w.WriteHeader(http.StatusForbidden)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

// GetVotesBoxByBoxId returns an answer box
func GetVotesBoxByBoxId(w http.ResponseWriter, r *http.Request) {
	context := flight.Context(w, r)
	boxid, err := uuid.FromString(context.Param("boxid"))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		votesBox, result := votesboxdomain.GetVotesBoxByBoxId(boxid)
		if result.IsSuccess() {
			responseString, err := json.Marshal(votesBox)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				w.Write([]byte(responseString))
			}
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

// GetVotesBoxByCode returns a votes box
func GetVotesBoxByCode(w http.ResponseWriter, r *http.Request) {
	context := flight.Context(w, r)
	code := context.Param("code")

	votesBox, result := votesboxdomain.GetVotesBoxByCode(code)
	if result.IsSuccess() {
		responseString, err := json.Marshal(votesBox)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write([]byte(responseString))
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// InsertVoteAnonymous inserts a vote for a votes box answer
func InsertVoteAnonymous(w http.ResponseWriter, r *http.Request) {
	var answer votesboxquestionanswermodel.VotesBoxQuestionAnswer
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&answer)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		entryVote, result := votesboxdomain.InsertVoteAnonymous(answer)
		if result.IsSuccess() {
			responseString, err := json.Marshal(entryVote)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				w.Write([]byte(responseString))
			}
		} else if result.Status == operationresult.PaymentRequired {
			w.WriteHeader(http.StatusPaymentRequired)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

// ActivateQuestion activates a question for an answer box
func ActivateQuestion(w http.ResponseWriter, r *http.Request) {
	var votesboxQuestion votesboxquestionmodel.VotesBoxQuestion
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&votesboxQuestion)
	tokenClaims := middleware.GetTokenClaims(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		result := votesboxdomain.ActivateQuestion(votesboxQuestion, tokenClaims.ID.String())
		if result.IsSuccess() {
			responseString, err := json.Marshal(votesboxQuestion)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				w.Write([]byte(responseString))
			}
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

// DeactivateQuestion activates a question for an answer box
func DeactivateQuestion(w http.ResponseWriter, r *http.Request) {
	var votesboxQuestion votesboxquestionmodel.VotesBoxQuestion
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&votesboxQuestion)
	tokenClaims := middleware.GetTokenClaims(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		result := votesboxdomain.DeactivateQuestion(votesboxQuestion, tokenClaims.ID.String())
		if result.IsSuccess() {
			responseString, err := json.Marshal(votesboxQuestion)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				w.Write([]byte(responseString))
			}
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
