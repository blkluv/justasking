package idpjustaskingcontroller

import (
	"encoding/json"
	"justasking/GO/common/operationresult"
	"justasking/GO/core/domain/authentication"
	"justasking/GO/core/model/idpjustasking"
	"net/http"

	"github.com/blue-jay/core/router"
)

// Load the routes.
func Load() {
	router.Post("/justaskinguser", CreateJustAskingUser)
}

// CreateJustAskingUser creates a record in the idpjustasking table
func CreateJustAskingUser(w http.ResponseWriter, r *http.Request) {
	var idpJustAskingUser idpjustaskingmodel.IdpJustAsking
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&idpJustAskingUser)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		createdUser, createdUserResult := authenticationdomain.CreateJustAskingUser(idpJustAskingUser)
		if createdUserResult.IsSuccess() {
			responseString, err := json.Marshal(createdUser)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				w.WriteHeader(http.StatusCreated)
				w.Write([]byte(responseString))
			}
		} else {
			if createdUserResult.Status == operationresult.UnprocessableEntity {
				w.WriteHeader(http.StatusUnprocessableEntity)
			} else if createdUserResult.Status == operationresult.Conflict {
				w.WriteHeader(http.StatusConflict)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	}
}
