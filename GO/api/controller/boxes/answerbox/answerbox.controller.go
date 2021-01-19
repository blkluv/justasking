package answerboxcontroller

import (
	"encoding/json"
	"justasking/GO/api/startup/middleware"
	"justasking/GO/common/operationresult"
	"justasking/GO/core/domain/boxes/answerbox"
	"justasking/GO/core/model/answerboxentry"
	"justasking/GO/core/model/answerboxquestion"
	"justasking/GO/core/model/boxes/answerbox"
	"justasking/GO/core/startup/flight"
	"net/http"

	"github.com/blue-jay/core/router"
	uuid "github.com/satori/go.uuid"
)

// Load the routes.
func Load() {
	router.Post("/answerbox", CreateAnswerBox, middleware.AuthorizedHandler)
	router.Post("/answerboxentry", InsertAnswerBoxEntry)
	router.Put("/answerbox/hide/", HideEntry, middleware.AuthorizedHandler)
	router.Put("/answerbox/unhide/", UnhideEntry, middleware.AuthorizedHandler)
	router.Put("/answerbox/hideall/question/", HideEntriesForQuestion, middleware.AuthorizedHandler)
	router.Put("/answerbox/unhideall/question/", UnhideEntriesForQuestion, middleware.AuthorizedHandler)
	router.Put("/answerbox/hideall/", HideAllEntries, middleware.AuthorizedHandler)
	router.Put("/answerbox/unhideall/", UnhideAllEntries, middleware.AuthorizedHandler)
	router.Put("/answerbox/activate/question/", ActivateQuestion, middleware.AuthorizedHandler)
	router.Put("/answerbox/deactivate/question/", DeactivateQuestion, middleware.AuthorizedHandler)
	router.Get("/answerbox/id/:boxid", GetAnswerBoxByBoxId)
	router.Get("/answerbox/code/:code", GetAnswerBoxByCode)
	router.Get("/answerboxentries/id/:boxid", GetAnswerBoxEntriesByBoxId)
	router.Get("/answerboxentries/code/:code", GetAnswerBoxEntriesByCode)
	router.Get("/answerboxentriesvisible/code/:code", GetVisibleAnswerBoxEntriesByCode)
}

// CreateAnswerBox creates an Answer Box
func CreateAnswerBox(w http.ResponseWriter, r *http.Request) {
	var answerBox answerboxmodel.AnswerBox

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&answerBox)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		tokenClaims := middleware.GetTokenClaims(r)
		answerBox.CreatedBy = tokenClaims.ID.String()
		answerBox.BaseBox.AccountId = tokenClaims.Account.Id
		boxId, result := answerboxdomain.CreateAnswerBox(answerBox)
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

// GetAnswerBoxByBoxId returns an answer box
func GetAnswerBoxByBoxId(w http.ResponseWriter, r *http.Request) {
	context := flight.Context(w, r)
	boxid, err := uuid.FromString(context.Param("boxid"))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		answerBox, result := answerboxdomain.GetAnswerBoxByBoxId(boxid)
		if result.IsSuccess() {
			responseString, err := json.Marshal(answerBox)
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

// GetAnswerBoxByCode returns an answer box
func GetAnswerBoxByCode(w http.ResponseWriter, r *http.Request) {
	context := flight.Context(w, r)
	code := context.Param("code")

	answerBox, result := answerboxdomain.GetAnswerBoxByCode(code)
	if result.IsSuccess() {
		responseString, err := json.Marshal(answerBox)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write([]byte(responseString))
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// ActivateQuestion activates a question for an answer box
func ActivateQuestion(w http.ResponseWriter, r *http.Request) {
	var answerBoxQuestion answerboxquestionmodel.AnswerBoxQuestion
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&answerBoxQuestion)
	tokenClaims := middleware.GetTokenClaims(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		result := answerboxdomain.ActivateQuestion(answerBoxQuestion, tokenClaims.ID.String())
		if result.IsSuccess() {
			responseString, err := json.Marshal(answerBoxQuestion)
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

// DeactivateQuestion deactivates a question for an answer box
func DeactivateQuestion(w http.ResponseWriter, r *http.Request) {
	var answerBoxQuestion answerboxquestionmodel.AnswerBoxQuestion
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&answerBoxQuestion)
	tokenClaims := middleware.GetTokenClaims(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		result := answerboxdomain.DeactivateQuestion(answerBoxQuestion, tokenClaims.ID.String())
		if result.IsSuccess() {
			responseString, err := json.Marshal(answerBoxQuestion)
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

// HideEntry hides the given entry for a given box
func HideEntry(w http.ResponseWriter, r *http.Request) {
	var answerBoxEntry answerboxentrymodel.AnswerBoxEntry
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&answerBoxEntry)
	tokenClaims := middleware.GetTokenClaims(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		result := answerboxdomain.HideEntry(answerBoxEntry, tokenClaims.ID.String())
		if result.IsSuccess() {
			responseString, err := json.Marshal(answerBoxEntry)
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

// UnhideEntry hides the given entry for a given box
func UnhideEntry(w http.ResponseWriter, r *http.Request) {
	var answerBoxEntry answerboxentrymodel.AnswerBoxEntry
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&answerBoxEntry)
	tokenClaims := middleware.GetTokenClaims(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		result := answerboxdomain.UnhideEntry(answerBoxEntry, tokenClaims.ID.String())
		if result.IsSuccess() {
			responseString, err := json.Marshal(answerBoxEntry)
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

// InsertAnswerBoxEntry returns the current user based on the token provided in the headers
func InsertAnswerBoxEntry(w http.ResponseWriter, r *http.Request) {
	var answerBoxEntry answerboxentrymodel.AnswerBoxEntry
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&answerBoxEntry)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		entry, result := answerboxdomain.InsertAnswerBoxEntry(answerBoxEntry)
		if result.IsSuccess() {
			responseString, err := json.Marshal(entry)
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

// GetAnswerBoxEntriesByBoxId returns all entries for an answer box
func GetAnswerBoxEntriesByBoxId(w http.ResponseWriter, r *http.Request) {
	context := flight.Context(w, r)
	boxId, err := uuid.FromString(context.Param("boxid"))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		entries, result := answerboxdomain.GetAnswerBoxEntriesByBoxId(boxId)
		if result.IsSuccess() {
			responseString, err := json.Marshal(entries)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				w.Write([]byte(responseString))
			}
		}
	}
}

// GetAnswerBoxEntriesByCode returns all entries for an answer box
func GetAnswerBoxEntriesByCode(w http.ResponseWriter, r *http.Request) {
	context := flight.Context(w, r)
	code := context.Param("code")

	entries, result := answerboxdomain.GetAnswerBoxEntriesByCode(code)
	if result.IsSuccess() {
		responseString, err := json.Marshal(entries)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write([]byte(responseString))
		}
	}
}

// GetVisibleAnswerBoxEntriesByCode returns all entries for a qeustion box, with their vote counts
func GetVisibleAnswerBoxEntriesByCode(w http.ResponseWriter, r *http.Request) {
	context := flight.Context(w, r)
	code := context.Param("code")

	entries, result := answerboxdomain.GetVisibleAnswerBoxEntriesByCode(code)
	if result.IsSuccess() {
		responseString, err := json.Marshal(entries)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write([]byte(responseString))
		}
	}
}

// HideEntriesForQuestion hides the given entry for a given box
func HideEntriesForQuestion(w http.ResponseWriter, r *http.Request) {
	var answerBoxQuestion answerboxquestionmodel.AnswerBoxQuestion
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&answerBoxQuestion)
	tokenClaims := middleware.GetTokenClaims(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		result := answerboxdomain.HideEntriesForQuestion(answerBoxQuestion, tokenClaims.ID.String())
		if result.IsSuccess() {
			responseString, err := json.Marshal(answerBoxQuestion)
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

// UnhideEntriesForQuestion hides the given entry for a given box
func UnhideEntriesForQuestion(w http.ResponseWriter, r *http.Request) {
	var answerBoxQuestion answerboxquestionmodel.AnswerBoxQuestion
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&answerBoxQuestion)
	tokenClaims := middleware.GetTokenClaims(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		result := answerboxdomain.UnhideEntriesForQuestion(answerBoxQuestion, tokenClaims.ID.String())
		if result.IsSuccess() {
			responseString, err := json.Marshal(answerBoxQuestion)
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

// HideAllEntries hides the given entry for a given box
func HideAllEntries(w http.ResponseWriter, r *http.Request) {
	var answerBox answerboxmodel.AnswerBox
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&answerBox)
	tokenClaims := middleware.GetTokenClaims(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		result := answerboxdomain.HideAllEntries(answerBox, tokenClaims.ID.String())
		if result.IsSuccess() {
			responseString, err := json.Marshal(answerBox)
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

// UnhideAllEntries hides the given entry for a given box
func UnhideAllEntries(w http.ResponseWriter, r *http.Request) {
	var answerBox answerboxmodel.AnswerBox
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&answerBox)
	tokenClaims := middleware.GetTokenClaims(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		result := answerboxdomain.UnhideAllEntries(answerBox, tokenClaims.ID.String())
		if result.IsSuccess() {
			responseString, err := json.Marshal(answerBox)
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
