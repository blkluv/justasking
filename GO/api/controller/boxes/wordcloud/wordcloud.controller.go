package wordcloudcontroller

import (
	"encoding/json"
	"justasking/GO/common/operationresult"
	"net/http"

	"justasking/GO/api/startup/middleware"
	"justasking/GO/core/domain/boxes/wordcloud"
	"justasking/GO/core/model/boxes/wordcloud"
	"justasking/GO/core/model/wordcloudresponse"
	"justasking/GO/core/startup/flight"

	"github.com/blue-jay/core/router"
	uuid "github.com/satori/go.uuid"
)

// Load the routes.
func Load() {
	router.Post("/wordcloud", CreateWordCloud, middleware.AuthorizedHandler)
	router.Put("/wordcloudresponse/hide/", HideAnswer, middleware.AuthorizedHandler)
	router.Put("/wordcloudresponse/unhide/", UnhideAnswer, middleware.AuthorizedHandler)
	router.Put("/wordcloud/hideall/", HideAllAnswers, middleware.AuthorizedHandler)
	router.Put("/wordcloud/unhideall/", UnhideAllAnswers, middleware.AuthorizedHandler)
	router.Get("/wordcloudresponse/full/id/:boxid", GetFullWordCloudResponsesByBoxId)
	router.Get("/wordcloudresponse/partial/id/:boxid", GetPartialWordCloudResponsesByBoxId)
	router.Get("/wordcloudresponse/full/code/:code", GetFullWordCloudResponsesByCode)
	router.Get("/wordcloudresponse/partial/code/:code", GetPartialWordCloudResponsesByCode)
	router.Get("/wordcloud/id/:boxid", GetWordCloudByBoxId)
	router.Get("/wordcloud/code/:code", GetWordCloudByCode)
	router.Post("/wordcloudresponse", CreateWordCloudResponse)
}

// CreateWordCloud returns the current user based on the token provided in the headers
func CreateWordCloud(w http.ResponseWriter, r *http.Request) {
	var wordCloud wordcloudmodel.WordCloud

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&wordCloud)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		tokenClaims := middleware.GetTokenClaims(r)
		wordCloud.CreatedBy = tokenClaims.ID.String()
		wordCloud.BaseBox.AccountId = tokenClaims.Account.Id
		boxId, result := wordclouddomain.CreateWordCloud(wordCloud)
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

// GetFullWordCloudResponsesByBoxId gets all answers for a specific box
func GetFullWordCloudResponsesByBoxId(w http.ResponseWriter, r *http.Request) {
	context := flight.Context(w, r)
	boxId, err := uuid.FromString(context.Param("boxid"))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		answers, result := wordclouddomain.GetFullWordCloudResponsesByBoxId(boxId)
		if result.IsSuccess() {
			responseString, err := json.Marshal(answers)
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

// GetPartialWordCloudResponsesByBoxId gets all answers for a specific box
func GetPartialWordCloudResponsesByBoxId(w http.ResponseWriter, r *http.Request) {
	context := flight.Context(w, r)
	boxId, err := uuid.FromString(context.Param("boxid"))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		answers, result := wordclouddomain.GetPartialWordCloudResponsesByBoxId(boxId)
		if result.IsSuccess() {
			responseString, err := json.Marshal(answers)
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

// GetFullWordCloudResponsesByCode gets all answers for a specific box
func GetFullWordCloudResponsesByCode(w http.ResponseWriter, r *http.Request) {
	context := flight.Context(w, r)
	code := context.Param("code")

	answers, result := wordclouddomain.GetFullWordCloudResponsesByCode(code)
	if result.IsSuccess() {
		responseString, err := json.Marshal(answers)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write([]byte(responseString))
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

}

// GetPartialWordCloudResponsesByCode gets all answers for a specific box
func GetPartialWordCloudResponsesByCode(w http.ResponseWriter, r *http.Request) {
	context := flight.Context(w, r)
	code := context.Param("code")

	answers, result := wordclouddomain.GetPartialWordCloudResponsesByCode(code)
	if result.IsSuccess() {
		responseString, err := json.Marshal(answers)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write([]byte(responseString))
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

}

// GetWordCloudByBoxId returns a word cloud
func GetWordCloudByBoxId(w http.ResponseWriter, r *http.Request) {
	context := flight.Context(w, r)
	boxid, err := uuid.FromString(context.Param("boxid"))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		wordCloud, result := wordclouddomain.GetWordCloudByBoxId(boxid)
		if result.IsSuccess() {
			responseString, err := json.Marshal(wordCloud)
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

// GetWordCloudByCode returns a word cloud
func GetWordCloudByCode(w http.ResponseWriter, r *http.Request) {
	context := flight.Context(w, r)
	code := context.Param("code")

	wordCloud, result := wordclouddomain.GetWordCloudByCode(code)
	if result.IsSuccess() {
		responseString, err := json.Marshal(wordCloud)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write([]byte(responseString))
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

}

// CreateWordCloudResponse creates a word cloud response
func CreateWordCloudResponse(w http.ResponseWriter, r *http.Request) {
	var wordCloudResponse wordcloudresponsemodel.WordCloudResponse
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&wordCloudResponse)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		entry, result := wordclouddomain.InsertWordCloudResponse(wordCloudResponse)
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

// HideAnswer hides the given answer for a given box
func HideAnswer(w http.ResponseWriter, r *http.Request) {
	var wordCloudResponse wordcloudresponsemodel.WordCloudResponse
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&wordCloudResponse)
	tokenClaims := middleware.GetTokenClaims(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		_, result := wordclouddomain.HideAnswer(wordCloudResponse.BoxId, wordCloudResponse.Response, tokenClaims.ID.String())
		if result.IsSuccess() {
			responseString, err := json.Marshal(wordCloudResponse)
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

// UnhideAnswer hides the given answer for a given box
func UnhideAnswer(w http.ResponseWriter, r *http.Request) {
	var wordCloudResponse wordcloudresponsemodel.WordCloudResponse
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&wordCloudResponse)
	tokenClaims := middleware.GetTokenClaims(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		_, result := wordclouddomain.UnhideAnswer(wordCloudResponse.BoxId, wordCloudResponse.Response, tokenClaims.ID.String())
		if result.IsSuccess() {
			responseString, err := json.Marshal(wordCloudResponse)
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

// HideAllAnswers hides the given answer for a given box
func HideAllAnswers(w http.ResponseWriter, r *http.Request) {
	var wordCloud wordcloudmodel.WordCloud
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&wordCloud)
	tokenClaims := middleware.GetTokenClaims(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		result := wordclouddomain.HideAllAnswers(wordCloud.BoxId, tokenClaims.ID.String())
		if result.IsSuccess() {
			responseString, err := json.Marshal(wordCloud)
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

// UnhideAllAnswers hides the given answer for a given box
func UnhideAllAnswers(w http.ResponseWriter, r *http.Request) {
	var wordCloud wordcloudmodel.WordCloud
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&wordCloud)
	tokenClaims := middleware.GetTokenClaims(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		result := wordclouddomain.UnhideAllAnswers(wordCloud.BoxId, tokenClaims.ID.String())
		if result.IsSuccess() {
			responseString, err := json.Marshal(wordCloud)
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
