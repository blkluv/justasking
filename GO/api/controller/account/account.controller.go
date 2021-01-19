package accountcontroller

import (
	"encoding/json"
	"justasking/GO/common/operationresult"
	"justasking/GO/core/model/accountuser"
	"justasking/GO/core/startup/flight"
	"net/http"

	"justasking/GO/api/startup/middleware"
	"justasking/GO/core/domain/account"
	"justasking/GO/core/domain/accountuser"
	"justasking/GO/core/domain/token"
	"justasking/GO/core/model/account"
	"justasking/GO/core/model/accountinvitation"

	"github.com/blue-jay/core/router"
)

// Load the routes.
func Load() {
	router.Post("/account/invite", InviteAccountUser, middleware.AuthorizedHandler)
	router.Post("/account/invite/resend", ResendAccountUserInvitation, middleware.AuthorizedHandler)
	router.Post("/account/join", JoinAccount, middleware.AuthorizedHandler)
	router.Post("/account/transfer", TransferOwnership, middleware.AuthorizedHandler)
	router.Put("/account/current", UpdateCurrentAccount, middleware.AuthorizedHandler)
	router.Put("/account", UpdateAccount, middleware.AuthorizedHandler)
	router.Put("/account/invite/cancel", CancelInvitation, middleware.AuthorizedHandler)
	router.Put("/account/user", UpdateAccountUser, middleware.AuthorizedHandler)
	router.Put("/account/user/delete", RemoveUserFromAccount, middleware.AuthorizedHandler)
	router.Get("/account/users", GetAllActiveAndPendingUsers, middleware.AuthorizedHandler)
	router.Get("/account/invite/:invitationCode", GetInvitationForJoin, middleware.AuthorizedHandler)
	router.Get("/accounts", GetAccounts, middleware.AuthorizedHandler)
}

// UpdateAccount updates an account
func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	tokenClaims := middleware.GetTokenClaims(r)

	if tokenClaims.RolePermissions.EditAccountName {
		var account accountmodel.Account

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&account)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			account, result := accountdomain.UpdateAccount(account, tokenClaims.ID)
			if result.IsSuccess() {
				responseString, err := json.Marshal(account)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
				} else {
					w.Write([]byte(responseString))
				}
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	} else {
		w.WriteHeader(http.StatusForbidden)
	}

}

// InviteAccountUser creates an account invitation
func InviteAccountUser(w http.ResponseWriter, r *http.Request) {
	tokenClaims := middleware.GetTokenClaims(r)

	var accountInvitation accountinvitationmodel.AccountInvitation

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&accountInvitation)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		invitationResult := accountdomain.InviteAccountUser(tokenClaims, accountInvitation)
		if invitationResult.IsSuccess() {
			w.WriteHeader(http.StatusOK)
		} else if invitationResult.Status == operationresult.UnprocessableEntity {
			//this means the model failed validation
			w.WriteHeader(http.StatusUnprocessableEntity)
		} else if invitationResult.Status == operationresult.Conflict {
			//this means that there is already an invitation for this user
			w.WriteHeader(http.StatusConflict)
		} else if invitationResult.Status == operationresult.PaymentRequired {
			//this means they're on BASIC
			w.WriteHeader(http.StatusPaymentRequired)
		} else if invitationResult.Status == operationresult.Forbidden {
			//this means they've reached the limit
			w.WriteHeader(http.StatusForbidden)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

// ResendAccountUserInvitation creates an account invitation
func ResendAccountUserInvitation(w http.ResponseWriter, r *http.Request) {
	tokenClaims := middleware.GetTokenClaims(r)

	var accountInvitation accountinvitationmodel.AccountInvitation

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&accountInvitation)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		invitationResult := accountdomain.ResendAccountUserInvitation(tokenClaims, accountInvitation)
		if invitationResult.IsSuccess() {
			w.WriteHeader(http.StatusOK)
		} else if invitationResult.Status == operationresult.Conflict {
			//this means that there is already an invitation for this user
			w.WriteHeader(http.StatusConflict)
		} else if invitationResult.Status == operationresult.PaymentRequired {
			//this means they're on BASIC
			w.WriteHeader(http.StatusPaymentRequired)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

// GetAllActiveAndPendingUsers gets all confirmed and pending users for an account
func GetAllActiveAndPendingUsers(w http.ResponseWriter, r *http.Request) {

	tokenClaims := middleware.GetTokenClaims(r)

	users, usersResult := accountdomain.GetActiveAndPendingAccountUsers(tokenClaims.Account.Id)
	if usersResult.IsSuccess() {
		responseString, err := json.Marshal(users)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write([]byte(responseString))
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

}

// UpdateCurrentAccount updates the user's current account and returns a new token
func UpdateCurrentAccount(w http.ResponseWriter, r *http.Request) {
	tokenClaims := middleware.GetTokenClaims(r)

	var accountUser accountusermodel.AccountUser

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&accountUser)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		updateResult := accountuserdomain.UpdateCurrentAccount(tokenClaims.ID, accountUser.AccountId)
		if updateResult.IsSuccess() {
			token, tokenResult := tokendomain.GetNewToken(tokenClaims.ID)
			if tokenResult.IsSuccess() {
				responseString, err := json.Marshal(token)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
				} else {
					w.Write([]byte(responseString))
				}
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
		} else if updateResult.Status == operationresult.Forbidden {
			w.WriteHeader(http.StatusForbidden)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

// JoinAccount addso a user to an account
func JoinAccount(w http.ResponseWriter, r *http.Request) {
	tokenClaims := middleware.GetTokenClaims(r)

	var accountInvitation accountinvitationmodel.AccountInvitation

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&accountInvitation)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		joinAccountResult := accountdomain.JoinAccount(tokenClaims, accountInvitation.InvitationCode)
		if joinAccountResult.IsSuccess() {
			token, tokenResult := tokendomain.GetNewToken(tokenClaims.ID)
			if tokenResult.IsSuccess() {
				responseString, err := json.Marshal(token)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
				} else {
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte(responseString))
				}
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
		} else if joinAccountResult.Status == operationresult.Forbidden {
			w.WriteHeader(http.StatusForbidden)
		} else if joinAccountResult.Status == operationresult.PaymentRequired {
			w.WriteHeader(http.StatusPaymentRequired)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

// GetInvitationForJoin retrieves a pending invitation's data
func GetInvitationForJoin(w http.ResponseWriter, r *http.Request) {
	context := flight.Context(w, r)
	accountInvitation := context.Param("invitationCode")

	tokenClaims := middleware.GetTokenClaims(r)

	invitation, invitationResult := accountdomain.GetInvitationForJoin(tokenClaims, accountInvitation)
	if invitationResult.IsSuccess() {
		responseString, err := json.Marshal(invitation)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write([]byte(responseString))
		}
	} else if invitationResult.Status == operationresult.Gone {
		w.WriteHeader(http.StatusGone)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// GetAccounts gets all accounts for a user
func GetAccounts(w http.ResponseWriter, r *http.Request) {
	tokenClaims := middleware.GetTokenClaims(r)

	accounts, accountsResult := accountdomain.GetAccounts(tokenClaims.ID)
	if accountsResult.IsSuccess() {
		responseString, err := json.Marshal(accounts)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write([]byte(responseString))
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// CancelInvitation cancels an invitation
func CancelInvitation(w http.ResponseWriter, r *http.Request) {
	tokenClaims := middleware.GetTokenClaims(r)

	var accountInvitation accountinvitationmodel.AccountInvitation

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&accountInvitation)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		cancelResult := accountdomain.CancelInvitation(tokenClaims, accountInvitation)
		if cancelResult.IsSuccess() {
			w.WriteHeader(http.StatusOK)
		} else if cancelResult.Status == operationresult.Forbidden {
			w.WriteHeader(http.StatusForbidden)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

// UpdateAccountUser updates an account user's role
func UpdateAccountUser(w http.ResponseWriter, r *http.Request) {
	tokenClaims := middleware.GetTokenClaims(r)

	var accountUser accountusermodel.AccountUser

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&accountUser)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		updateResult := accountuserdomain.UpdateAccountUserRole(tokenClaims, accountUser)
		if updateResult.IsSuccess() {
			responseString, err := json.Marshal(accountUser)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				w.Write([]byte(responseString))
			}
		} else if updateResult.Status == operationresult.Forbidden {
			w.WriteHeader(http.StatusForbidden)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

// TransferOwnership transfers ownership from one user to another
func TransferOwnership(w http.ResponseWriter, r *http.Request) {
	tokenClaims := middleware.GetTokenClaims(r)

	var accountUser accountusermodel.AccountUser

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&accountUser)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		transferResult := accountuserdomain.TransferOwnership(tokenClaims, accountUser)
		if transferResult.IsSuccess() {
			token, tokenResult := tokendomain.GetNewToken(tokenClaims.ID)
			if tokenResult.IsSuccess() {
				responseString, err := json.Marshal(token)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
				} else {
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte(responseString))
				}
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
		} else if transferResult.Status == operationresult.Forbidden {
			w.WriteHeader(http.StatusForbidden)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

// RemoveUserFromAccount removes a user from an account
func RemoveUserFromAccount(w http.ResponseWriter, r *http.Request) {
	tokenClaims := middleware.GetTokenClaims(r)

	var accountUser accountusermodel.AccountUser

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&accountUser)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		removeResult := accountuserdomain.RemoveUserFromAccount(tokenClaims, accountUser)
		if removeResult.IsSuccess() {
			w.WriteHeader(http.StatusOK)
		} else if removeResult.Status == operationresult.Forbidden {
			w.WriteHeader(http.StatusForbidden)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
