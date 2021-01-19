package usercontroller

import (
	"encoding/json"
	"justasking/GO/common/operationresult"
	"justasking/GO/common/token"
	"justasking/GO/core/domain/accountuser"
	"justasking/GO/core/domain/token"
	"justasking/GO/core/model/passwordresetrequest"
	"justasking/GO/core/model/user"
	"net/http"

	"github.com/blue-jay/core/router"

	"justasking/GO/api/startup/middleware"
	"justasking/GO/core/domain/user"
)

// Load the routes.
func Load() {
	router.Get("/user", GetCurrentUser, middleware.AuthorizedHandler)
	router.Post("/user/resetpassword", RequestPasswordReset)
	router.Post("/user/logout", LogOut, middleware.AuthorizedHandler)
	router.Put("/user/password", PasswordReset)
}

// GetCurrentUser returns the current user based on the token provided in the headers
func GetCurrentUser(w http.ResponseWriter, r *http.Request) {

	tokenClaims := middleware.GetTokenClaims(r)
	user, userResult := userdomain.GetUser(tokenClaims.ID)

	if userResult.IsSuccess() == true {
		responseString, err := json.Marshal(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write([]byte(responseString))
		}
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

// RequestPasswordReset creates a reset password request
func RequestPasswordReset(w http.ResponseWriter, r *http.Request) {
	var user usermodel.User

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		requestResult := userdomain.RequestPasswordReset(user.Email)
		if requestResult.IsSuccess() {
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

// PasswordReset resets a password
func PasswordReset(w http.ResponseWriter, r *http.Request) {
	var resetRequest passwordresetrequestmodel.PasswordResetRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&resetRequest)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		user, updateResult := userdomain.UpdatePassword(resetRequest)
		if updateResult.IsSuccess() {
			tokenString, tokenStringResult := tokendomain.GetNewToken(user.ID)
			if tokenStringResult.IsSuccess() {
				token := &token.Token{Value: tokenString}
				responseString, err := json.Marshal(token)
				if err != nil {
					// purposely not returning error in case of token generation failure. they will get redirected to the login page if the token is invalid.
					w.WriteHeader(http.StatusOK)
				} else {
					w.Write([]byte(responseString))
				}
			}
			w.WriteHeader(http.StatusOK)
		} else if updateResult.Status == operationresult.Gone {
			w.WriteHeader(http.StatusGone)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

// LogOut invalidates a user's token by updating its version
func LogOut(w http.ResponseWriter, r *http.Request) {
	tokenClaims := middleware.GetTokenClaims(r)

	tokenVersionResult := accountuserdomain.UpdateAccountUserToken(tokenClaims.ID, tokenClaims.Account.Id)
	if tokenVersionResult.IsSuccess() {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
