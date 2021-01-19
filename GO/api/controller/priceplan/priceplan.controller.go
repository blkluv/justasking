package priceplancontroller

import (
	"encoding/json"
	"justasking/GO/api/startup/middleware"
	"justasking/GO/common/operationresult"
	"justasking/GO/core/domain/priceplan"
	"justasking/GO/core/startup/flight"
	"net/http"

	"github.com/blue-jay/core/router"
	uuid "github.com/satori/go.uuid"
)

// Load the routes.
func Load() {
	router.Get("/priceplan", GetPublicPricePlans)
	router.Get("/priceplan/name/:priceplanname", GetPricePlanByName)
	router.Get("/priceplan/id/:priceplanid", GetPricePlanById)
	router.Get("/priceplan/custom/:licensecode", GetPricePlanByLicenseCode, middleware.AuthorizedHandler)
}

// GetPublicPricePlans should gets all boxes for an owner
func GetPublicPricePlans(w http.ResponseWriter, r *http.Request) {
	plans, result := priceplandomain.GetPublicPricePlans()
	if result.IsSuccess() {
		responseString, err := json.Marshal(plans)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write([]byte(responseString))
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

}

// GetPricePlanByName should gets all boxes for an owner
func GetPricePlanByName(w http.ResponseWriter, r *http.Request) {
	context := flight.Context(w, r)
	pricePlanName := context.Param("priceplanname")

	planDetails, result := priceplandomain.GetPricePlanDetailsByPlanName(pricePlanName)
	if result.IsSuccess() {
		responseString, err := json.Marshal(planDetails)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write([]byte(responseString))
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

}

// GetPricePlanById should gets all boxes for an owner
func GetPricePlanById(w http.ResponseWriter, r *http.Request) {
	context := flight.Context(w, r)
	pricePlanId, err := uuid.FromString(context.Param("priceplanid"))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		planDetails, result := priceplandomain.GetPricePlanDetailsByPlanId(pricePlanId)
		if result.IsSuccess() {
			responseString, err := json.Marshal(planDetails)
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

// GetPricePlanByLicenseCode gets custom price plan details by license code
func GetPricePlanByLicenseCode(w http.ResponseWriter, r *http.Request) {
	context := flight.Context(w, r)
	licenseCode := context.Param("licensecode")

	tokenClaims := middleware.GetTokenClaims(r)

	planDetails, result := priceplandomain.GetPricePlanByLicenseCode(licenseCode, tokenClaims.ID)
	if result.IsSuccess() {
		responseString, err := json.Marshal(planDetails)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write([]byte(responseString))
		}
	} else if result.Status == operationresult.Forbidden {
		w.WriteHeader(http.StatusForbidden)
	} else if result.Status == operationresult.NotFound {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

}
