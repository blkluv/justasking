package tokencontroller

import (
	"justasking/GO/common/operationresult"
	"justasking/GO/core/domain/token"
	"net/http"

	"github.com/blue-jay/core/router"

	"encoding/json"
	"justasking/GO/api/startup/middleware"
	"justasking/GO/common/authcontainer"
	"justasking/GO/common/token"
	"justasking/GO/core/domain/authentication"
)

// Load the routes.
func Load() {
	router.Post("/token", Token)
	router.Get("/token", NewToken, middleware.AuthorizedHandler)
}

// Token returns a JWT
func Token(w http.ResponseWriter, r *http.Request) {

	var authContainer authcontainer.AuthContainer
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&authContainer)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		tokenString, userWasCreated, result := authenticationdomain.GetToken(authContainer)
		if result.IsSuccess() {
			token := &token.Token{Value: tokenString}
			responseString, err := json.Marshal(token)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				if userWasCreated {
					w.WriteHeader(http.StatusCreated)
				}
				w.Write([]byte(responseString))
			}
		} else if result.Status == operationresult.Conflict {
			w.WriteHeader(http.StatusConflict)
		} else if result.Status == operationresult.Unauthorized {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

// NewToken returns a JWT for a user who is already logged in. Used when changing subscriptions.
func NewToken(w http.ResponseWriter, r *http.Request) {
	tokenClaims := middleware.GetTokenClaims(r)

	tokenString, tokenResult := tokendomain.GetNewToken(tokenClaims.ID)
	if tokenResult.IsSuccess() {
		token := &token.Token{Value: tokenString}
		responseString, err := json.Marshal(token)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write([]byte(responseString))
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
