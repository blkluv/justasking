package baseboxcontroller

import (
	"encoding/json"
	"justasking/GO/api/startup/middleware"
	"justasking/GO/common/operationresult"
	"justasking/GO/common/utility"
	"justasking/GO/core/domain/boxes/basebox"
	"justasking/GO/core/model/boxes/basebox"
	"justasking/GO/core/startup/flight"
	"net/http"

	"github.com/blue-jay/core/router"
	uuid "github.com/satori/go.uuid"
)

// Load the routes.
func Load() {
	router.Get("/basebox/account", GetBaseBoxesByAccountId, middleware.AuthorizedHandler)
	router.Get("/baseboxbycode/:code", GetBaseBoxByCode)
	router.Get("/basebox/details/:code", GetBaseBoxByCodeAuthorized, middleware.AuthorizedHandler)
	router.Get("/basebox/exists/:code", BoxCodeExists)
	router.Get("/basebox/code", GetRandomAccessCode)
	router.Post("/basebox/activate/:boxid", ActivateBaseBox, middleware.AuthorizedHandler)
	router.Post("/basebox/deactivate/:boxid", DeactivateBaseBox, middleware.AuthorizedHandler)
	router.Put("/basebox/delete", DeletePoll, middleware.AuthorizedHandler)
}

// GetBaseBoxesByAccountId gets all boxes for an owner
func GetBaseBoxesByAccountId(w http.ResponseWriter, r *http.Request) {
	tokenClaims := middleware.GetTokenClaims(r)

	boxes, result := baseboxdomain.GetBoxesByAccountId(tokenClaims.Account.Id, tokenClaims)
	if result.IsSuccess() {
		responseString, err := json.Marshal(boxes)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write([]byte(responseString))
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

}

// GetBaseBoxByCode should get the box with the provided code
func GetBaseBoxByCode(w http.ResponseWriter, r *http.Request) {
	context := flight.Context(w, r)
	code := context.Param("code")
	box, result := baseboxdomain.GetBaseBoxByCode(code)
	if result.IsSuccess() {
		responseString, err := json.Marshal(box)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write([]byte(responseString))
		}
	} else if result.Message == "record not found" {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// BoxCodeExists gets all boxes for an owner
func BoxCodeExists(w http.ResponseWriter, r *http.Request) {
	context := flight.Context(w, r)
	code := context.Param("code")

	exists, result := baseboxdomain.BoxCodeExists(code)
	if result.IsSuccess() {
		responseString, err := json.Marshal(exists)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write([]byte(responseString))
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// ActivateBaseBox activates a base box
func ActivateBaseBox(w http.ResponseWriter, r *http.Request) {
	context := flight.Context(w, r)
	boxid, err := uuid.FromString(context.Param("boxid"))
	tokenClaims := middleware.GetTokenClaims(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		phoneNumber, result := baseboxdomain.ActivateBaseBox(boxid, tokenClaims.ID, tokenClaims.MembershipDetails)
		if result.IsSuccess() {
			w.WriteHeader(http.StatusOK)
			responseString, err := json.Marshal(phoneNumber)
			if err == nil {
				w.Write([]byte(responseString))
			}
		} else if result.Status == operationresult.Conflict {
			w.WriteHeader(http.StatusConflict)
			responseString, err := json.Marshal(result)
			if err == nil {
				w.Write([]byte(responseString))
			}
		} else if result.Status == operationresult.Forbidden {
			w.WriteHeader(http.StatusForbidden)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

// DeactivateBaseBox activates a base box
func DeactivateBaseBox(w http.ResponseWriter, r *http.Request) {
	context := flight.Context(w, r)
	boxid, err := uuid.FromString(context.Param("boxid"))
	tokenClaims := middleware.GetTokenClaims(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		result := baseboxdomain.DeactivateBaseBox(boxid, tokenClaims.ID)
		if result.IsSuccess() {
			w.WriteHeader(http.StatusOK)
			responseString, err := json.Marshal(nil)
			if err == nil {
				w.Write([]byte(responseString))
			}
		} else if result.Status == operationresult.Conflict {
			w.WriteHeader(http.StatusConflict)
			responseString, err := json.Marshal(result)
			if err == nil {
				w.Write([]byte(responseString))
			}
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

// GetRandomAccessCode returns a random access code
func GetRandomAccessCode(w http.ResponseWriter, r *http.Request) {

	var code = utility.Generate()

	codeString, err := json.Marshal(code)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Write([]byte(codeString))
	}
}

//GetBaseBoxByCodeForDetailsPage gets a base box for the box details page
func GetBaseBoxByCodeAuthorized(w http.ResponseWriter, r *http.Request) {
	context := flight.Context(w, r)
	code := context.Param("code")

	tokenClaims := middleware.GetTokenClaims(r)

	box, result := baseboxdomain.GetBaseBoxByCodeAuthorized(code, tokenClaims)
	if result.IsSuccess() {
		responseString, err := json.Marshal(box)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write([]byte(responseString))
		}
	} else if result.Status == operationresult.Forbidden {
		w.WriteHeader(http.StatusForbidden)
	} else if result.Message == "record not found" {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

//DeletePoll deletes a poll from an account
func DeletePoll(w http.ResponseWriter, r *http.Request) {
	var baseBox baseboxmodel.BaseBox
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&baseBox)
	tokenClaims := middleware.GetTokenClaims(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		result := baseboxdomain.DeleteBaseBox(baseBox.ID, tokenClaims)
		if result.IsSuccess() {
			w.WriteHeader(http.StatusOK)
		} else if result.Status == operationresult.Forbidden {
			w.WriteHeader(http.StatusForbidden)
		} else if result.Status == operationresult.NotFound {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
