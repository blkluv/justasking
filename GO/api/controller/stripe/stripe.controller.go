package stripecontroller

import (
	"encoding/json"
	"justasking/GO/api/startup/middleware"
	"justasking/GO/core/domain/stripe"
	"justasking/GO/core/domain/token"
	"justasking/GO/core/model/customplanlicense"
	"net/http"

	"github.com/stripe/stripe-go"

	"github.com/blue-jay/core/router"
)

// Load the routes.
func Load() {
	router.Post("/card", UpdateCard, middleware.AuthorizedHandler)
	router.Post("/plan", UpdateSubscription, middleware.AuthorizedHandler)
	router.Post("/plan/custom", UpdateCustomSubscription, middleware.AuthorizedHandler)
	router.Get("/stripedata", GetStripeData, middleware.AuthorizedHandler)
}

// UpdateSubscription creates a subscription in Stripe
func UpdateSubscription(w http.ResponseWriter, r *http.Request) {
	tokenClaims := middleware.GetTokenClaims(r)
	var stripePlan stripe.Plan

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&stripePlan)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		planResult := stripedomain.UpdateSubscription(stripePlan, tokenClaims)
		if planResult.IsSuccess() {
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
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

// UpdateCustomSubscription creates a subscription in Stripe
func UpdateCustomSubscription(w http.ResponseWriter, r *http.Request) {
	tokenClaims := middleware.GetTokenClaims(r)
	var pricePlan customplanlicensemodel.CustomPlanLicense

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&pricePlan)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		planResult := stripedomain.UpdateCustomSubscription(pricePlan.LicenseCode, tokenClaims)
		if planResult.IsSuccess() {
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
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

// UpdateCard updates a user's default card in stripe
func UpdateCard(w http.ResponseWriter, r *http.Request) {
	tokenClaims := middleware.GetTokenClaims(r)
	var stripeToken stripe.Token

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&stripeToken)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		customer, result := stripedomain.UpdateCreditCard(stripeToken, tokenClaims.ID, tokenClaims.Account.Id)
		if result.IsSuccess() {
			responseString, err := json.Marshal(customer)
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

// GetStripeData gets the stripe user data from our database
func GetStripeData(w http.ResponseWriter, r *http.Request) {
	tokenClaims := middleware.GetTokenClaims(r)

	stripeData, stripeDataResponse := stripedomain.GetUserStripeDataByAccountId(tokenClaims.Account.Id)
	if stripeDataResponse.IsSuccess() {
		responseString, err := json.Marshal(stripeData)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write([]byte(responseString))
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
