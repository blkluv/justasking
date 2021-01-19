package questionboxcontroller

import (
	"justasking/GO/common/operationresult"
	"justasking/GO/core/startup/flight"

	"encoding/json"
	"justasking/GO/api/startup/middleware"
	"justasking/GO/core/domain/boxes/questionbox"
	"justasking/GO/core/model/boxes/questionbox"
	"justasking/GO/core/model/questionboxentry"
	"justasking/GO/core/model/questionboxentryvote"
	"net/http"

	"github.com/blue-jay/core/router"
	uuid "github.com/satori/go.uuid"
)

// Load the routes.
func Load() {
	router.Post("/questionbox", CreateQuestionBox, middleware.AuthorizedHandler)
	router.Post("/questionboxentry", InsertQuestionBoxEntry)
	router.Post("/questionboxentryvote", InsertQuestionBoxEntryVote)
	router.Post("/questionboxentryvote/upvotefromdownvote", UpvoteFromDownvote)
	router.Post("/questionboxentryvote/downvotefromupvote", DownvoteFromUpvote)
	router.Post("/questionboxentryvoteundo/up", UndoUpvote)
	router.Post("/questionboxentryvoteundo/down", UndoDownvote)
	router.Put("/questionbox/hide/", HideEntry, middleware.AuthorizedHandler)
	router.Put("/questionbox/unhide/", UnhideEntry, middleware.AuthorizedHandler)
	router.Put("/questionbox/hideall/", HideAllEntries, middleware.AuthorizedHandler)
	router.Put("/questionbox/unhideall/", UnhideAllEntries, middleware.AuthorizedHandler)
	router.Get("/questionbox/id/:boxid", GetQuestionBoxByBoxId)
	router.Get("/questionbox/code/:code", GetQuestionBoxByCode)
	router.Get("/questionboxentries/id/:boxid", GetQuestionBoxEntriesByBoxId)
	router.Get("/questionboxentries/code/:code", GetQuestionBoxEntriesByCode)
	router.Get("/questionboxentriesvisible/code/:code", GetVisibleQuestionBoxEntriesByCode)
}

// CreateQuestionBox creates a Question Box
func CreateQuestionBox(w http.ResponseWriter, r *http.Request) {
	var questionBox questionboxmodel.QuestionBox

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&questionBox)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		tokenClaims := middleware.GetTokenClaims(r)
		questionBox.CreatedBy = tokenClaims.ID.String()
		questionBox.BaseBox.AccountId = tokenClaims.Account.Id
		boxId, result := questionboxdomain.CreateQuestionBox(questionBox)
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

// GetQuestionBoxByBoxId returns a question box
func GetQuestionBoxByBoxId(w http.ResponseWriter, r *http.Request) {
	context := flight.Context(w, r)
	boxid, err := uuid.FromString(context.Param("boxid"))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		questionBox, result := questionboxdomain.GetQuestionBoxByBoxId(boxid)
		if result.IsSuccess() {
			responseString, err := json.Marshal(questionBox)
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

// GetQuestionBoxByCode returns a question box
func GetQuestionBoxByCode(w http.ResponseWriter, r *http.Request) {
	context := flight.Context(w, r)
	code := context.Param("code")

	questionBox, result := questionboxdomain.GetQuestionBoxByCode(code)
	if result.IsSuccess() {
		responseString, err := json.Marshal(questionBox)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write([]byte(responseString))
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

}

// GetQuestionBoxEntriesByBoxId returns all entries for a qeustion box, with their vote counts
func GetQuestionBoxEntriesByBoxId(w http.ResponseWriter, r *http.Request) {
	context := flight.Context(w, r)
	boxId, err := uuid.FromString(context.Param("boxid"))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		entries, result := questionboxdomain.GetQuestionBoxEntriesByBoxId(boxId)
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

// GetQuestionBoxEntriesByCode returns all entries for a qeustion box, with their vote counts
func GetQuestionBoxEntriesByCode(w http.ResponseWriter, r *http.Request) {
	context := flight.Context(w, r)
	code := context.Param("code")

	entries, result := questionboxdomain.GetQuestionBoxEntriesByCode(code)
	if result.IsSuccess() {
		responseString, err := json.Marshal(entries)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write([]byte(responseString))
		}
	}

}

// GetVisibleQuestionBoxEntriesByCode returns all entries for a qeustion box, with their vote counts
func GetVisibleQuestionBoxEntriesByCode(w http.ResponseWriter, r *http.Request) {
	context := flight.Context(w, r)
	code := context.Param("code")

	entries, result := questionboxdomain.GetVisibleQuestionBoxEntriesByCode(code)
	if result.IsSuccess() {
		responseString, err := json.Marshal(entries)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write([]byte(responseString))
		}
	}
}

// InsertQuestionBoxEntry returns the current user based on the token provided in the headers
func InsertQuestionBoxEntry(w http.ResponseWriter, r *http.Request) {
	var questionBoxEntry questionboxentrymodel.QuestionBoxEntry
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&questionBoxEntry)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		entry, result := questionboxdomain.InsertQuestionBoxEntry(questionBoxEntry)
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

// HideEntry hides the given entry for a given box
func HideEntry(w http.ResponseWriter, r *http.Request) {
	var questionBoxEntry questionboxentrymodel.QuestionBoxEntry
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&questionBoxEntry)
	tokenClaims := middleware.GetTokenClaims(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		result := questionboxdomain.HideEntry(questionBoxEntry, tokenClaims.ID.String())
		if result.IsSuccess() {
			responseString, err := json.Marshal(questionBoxEntry)
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
	var questionBoxEntry questionboxentrymodel.QuestionBoxEntry
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&questionBoxEntry)
	tokenClaims := middleware.GetTokenClaims(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		result := questionboxdomain.UnhideEntry(questionBoxEntry, tokenClaims.ID.String())
		if result.IsSuccess() {
			responseString, err := json.Marshal(questionBoxEntry)
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

// InsertQuestionBoxEntryVote inserts a vote for a question box
func InsertQuestionBoxEntryVote(w http.ResponseWriter, r *http.Request) {
	var questionBoxEntryVote questionboxentryvotemodel.QuestionBoxEntryVote
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&questionBoxEntryVote)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		entryVote, result := questionboxdomain.InsertQuestionBoxEntryVote(questionBoxEntryVote)
		if result.IsSuccess() {
			responseString, err := json.Marshal(entryVote)
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

// UpvoteFromDownvote cancels out previous downvote by giving it a -1, and adds an upvote
func UpvoteFromDownvote(w http.ResponseWriter, r *http.Request) {
	var questionBoxEntryVote questionboxentryvotemodel.QuestionBoxEntryVote
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&questionBoxEntryVote)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		entryVote, result := questionboxdomain.UpvoteFromDownvote(questionBoxEntryVote)
		if result.IsSuccess() {
			responseString, err := json.Marshal(entryVote)
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

// DownvoteFromUpvote cancels out previous upvote by giving it a -1, and adds a downvote
func DownvoteFromUpvote(w http.ResponseWriter, r *http.Request) {
	var questionBoxEntryVote questionboxentryvotemodel.QuestionBoxEntryVote
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&questionBoxEntryVote)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		entryVote, result := questionboxdomain.DownvoteFromUpvote(questionBoxEntryVote)
		if result.IsSuccess() {
			responseString, err := json.Marshal(entryVote)
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

// UndoUpvote cancels out previous upvote by giving it a -1
func UndoUpvote(w http.ResponseWriter, r *http.Request) {
	var questionBoxEntryVote questionboxentryvotemodel.QuestionBoxEntryVote
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&questionBoxEntryVote)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		entryVote, result := questionboxdomain.UndoUpvote(questionBoxEntryVote)
		if result.IsSuccess() {
			responseString, err := json.Marshal(entryVote)
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

// UndoDownvote cancels out previous upvote by giving it a -1
func UndoDownvote(w http.ResponseWriter, r *http.Request) {
	var questionBoxEntryVote questionboxentryvotemodel.QuestionBoxEntryVote
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&questionBoxEntryVote)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		entryVote, result := questionboxdomain.UndoDownvote(questionBoxEntryVote)
		if result.IsSuccess() {
			responseString, err := json.Marshal(entryVote)
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
	var questionBox questionboxmodel.QuestionBox
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&questionBox)
	tokenClaims := middleware.GetTokenClaims(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		result := questionboxdomain.HideAllEntries(questionBox, tokenClaims.ID.String())
		if result.IsSuccess() {
			responseString, err := json.Marshal(questionBox)
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
	var questionBox questionboxmodel.QuestionBox
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&questionBox)
	tokenClaims := middleware.GetTokenClaims(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		result := questionboxdomain.UnhideAllEntries(questionBox, tokenClaims.ID.String())
		if result.IsSuccess() {
			responseString, err := json.Marshal(questionBox)
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
