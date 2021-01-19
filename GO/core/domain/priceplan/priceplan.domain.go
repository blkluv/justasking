package priceplandomain

import (
	"fmt"
	"justasking/GO/common/operationresult"
	"justasking/GO/core/domain/applogs"
	"justasking/GO/core/model/customplanlicense"
	"justasking/GO/core/model/priceplan"
	"justasking/GO/core/repo/priceplan"
	"strconv"

	"github.com/jinzhu/gorm"

	uuid "github.com/satori/go.uuid"
)

var domainName = "PricePlanDomain"

// GetPublicPricePlans gets all price plans features
func GetPublicPricePlans() ([]priceplanmodel.PricePlan, *operationresult.OperationResult) {
	functionName := "GetPublicPricePlans"
	result := operationresult.New()

	plans, err := priceplanrepo.GetPublicPricePlans()

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting all price plans. Error: [%v]", msg), false)
	}

	return plans, result
}

// GetPricePlanDetailsByPlanName gets price plan feature details
func GetPricePlanDetailsByPlanName(planName string) (priceplanmodel.PricePlan, *operationresult.OperationResult) {
	functionName := "GetPricePlanDetailsByPlanName"
	result := operationresult.New()
	pricePlan := priceplanmodel.PricePlan{}

	planDetails, err := priceplanrepo.GetPricePlanDetailsByPlanName(planName)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting price plan details for price plan [%v]. Error: [%v]", planName, msg), false)
	} else if len(planDetails) <= 0 {
		msg := fmt.Sprintf("No plan details found for plan name [%v]", planName)
		result.Status = operationresult.Error
		result.Message = msg
		applogsdomain.LogError(domainName, functionName, msg, false)
	} else {
		pricePlan = mapPlanFeatures(planDetails)
	}

	return pricePlan, result
}

// GetPricePlanDetailsByPlanId gets price plan feature details
func GetPricePlanDetailsByPlanId(planId uuid.UUID) (priceplanmodel.PricePlan, *operationresult.OperationResult) {
	functionName := "GetPricePlanDetailsByPlanId"
	result := operationresult.New()
	pricePlan := priceplanmodel.PricePlan{}

	planDetails, err := priceplanrepo.GetPricePlanDetailsByPlanId(planId)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting price plan details for price plan [%v]. Error: [%v]", planId, msg), false)
	} else if len(planDetails) <= 0 {
		msg := fmt.Sprintf("No plan details found for planId [%v]", planId)
		result.Status = operationresult.Error
		result.Message = msg
		applogsdomain.LogError(domainName, functionName, msg, false)
	} else {
		pricePlan = mapPlanFeatures(planDetails)
	}

	return pricePlan, result
}

// GetPricePlanDetailsByAccountId gets price plan feature details for an account
func GetPricePlanDetailsByAccountId(accountId uuid.UUID) (priceplanmodel.PricePlan, *operationresult.OperationResult) {
	functionName := "GetPricePlanDetailsByAccountId"
	result := operationresult.New()
	pricePlan := priceplanmodel.PricePlan{}

	planDetails, err := priceplanrepo.GetPricePlanDetailsByAccountId(accountId)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting price plan details for account [%v]. Error: [%v]", accountId, msg), false)
	} else if len(planDetails) <= 0 {
		msg := fmt.Sprintf("Unable to retrieve priceplandetails for accountId [%v].", accountId)
		result.Status = operationresult.Error
		result.Message = msg
		applogsdomain.LogError(domainName, functionName, msg, false)
	} else {
		pricePlan = mapPlanFeatures(planDetails)
	}

	return pricePlan, result
}

// GetPricePlanDetailsByBoxId gets price plan feature details for a box
func GetPricePlanDetailsByBoxId(boxId uuid.UUID) (priceplanmodel.PricePlan, *operationresult.OperationResult) {
	functionName := "GetPricePlanDetailsByAccountId"
	result := operationresult.New()
	pricePlan := priceplanmodel.PricePlan{}

	planDetails, err := priceplanrepo.GetPricePlanDetailsByBoxId(boxId)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting price plan details for box [%v]. Error: [%v]", boxId, msg), false)
	} else if len(planDetails) <= 0 {
		msg := fmt.Sprintf("No plan details found for boxId [%v]", boxId)
		result.Status = operationresult.Error
		result.Message = msg
		applogsdomain.LogError(domainName, functionName, msg, false)
	} else {
		pricePlan = mapPlanFeatures(planDetails)
	}

	return pricePlan, result
}

// GetPricePlanDetailsByUserId gets price plan feature details for an account
func GetPricePlanDetailsByUserId(userId uuid.UUID) (priceplanmodel.PricePlan, *operationresult.OperationResult) {
	functionName := "GetPricePlanDetailsByUserId"
	result := operationresult.New()
	pricePlan := priceplanmodel.PricePlan{}

	planDetails, err := priceplanrepo.GetPricePlanDetailsByUserId(userId)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting price plan details for user [%v]. Error: [%v]", userId, msg), false)
	} else if len(planDetails) <= 0 {
		msg := fmt.Sprintf("No plan details found for userId [%v]", userId)
		result.Status = operationresult.Error
		result.Message = msg
		applogsdomain.LogError(domainName, functionName, msg, false)
	} else {
		pricePlan = mapPlanFeatures(planDetails)
	}

	return pricePlan, result
}

// GetPricePlanByLicenseCode gets custom price plan details by license code
func GetPricePlanByLicenseCode(licenseCode string, userId uuid.UUID) (priceplanmodel.PricePlan, *operationresult.OperationResult) {
	functionName := "GetPricePlanByLicenseCode"
	result := operationresult.New()
	pricePlan := priceplanmodel.PricePlan{}

	planLicense, planLicenseResult := GetCustomPlanLicense(licenseCode)
	if planLicenseResult.IsSuccess() {

		if planLicense.UserId == userId {
			planDetails, err := priceplanrepo.GetPricePlanByLicenseCode(licenseCode)

			if err != nil {
				msg := err.Error()
				result = operationresult.CreateErrorResult(msg, err)
				applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting price plan details for license code [%v]. Error: [%v]", licenseCode, msg), false)
			} else if len(planDetails) <= 0 {
				msg := fmt.Sprintf("No plan details found for license code [%v]", licenseCode)
				result.Status = operationresult.Error
				result.Message = msg
				applogsdomain.LogError(domainName, functionName, msg, false)
			} else {
				pricePlan = mapPlanFeatures(planDetails)
			}
		} else {
			msg := fmt.Sprintf("Unable to get license details for user [%v]. License plan is for user [%v]", licenseCode, planLicense.UserId)
			result.Status = operationresult.Forbidden
			result.Message = msg
			applogsdomain.LogError(domainName, functionName, msg, false)
		}
	} else if planLicenseResult.Status == operationresult.NotFound {
		msg := fmt.Sprintf("Unable to get license details for license code [%v]. Message: [%v]", licenseCode, planLicenseResult.Message)
		result.Status = operationresult.NotFound
		result.Message = msg
		applogsdomain.LogError(domainName, functionName, msg, false)
	} else {
		msg := fmt.Sprintf("Unable to get license details for license code [%v]. Message: [%v]", licenseCode, planLicenseResult.Message)
		result.Status = operationresult.Error
		result.Message = msg
		applogsdomain.LogError(domainName, functionName, msg, false)
	}

	return pricePlan, result
}

// GetCustomPlanLicense gets custom plan license details
func GetCustomPlanLicense(licenseCode string) (customplanlicensemodel.CustomPlanLicense, *operationresult.OperationResult) {
	functionName := "GetCustomPlanLicense"
	result := operationresult.New()

	planLicense, err := priceplanrepo.GetCustomPlanLicense(licenseCode)
	if err != nil {
		msg := err.Error()

		if err == gorm.ErrRecordNotFound {
			result.Status = operationresult.NotFound
			result.Message = fmt.Sprintf("Custom plan with license code [%v] not found.", licenseCode)
		} else {
			result = operationresult.CreateErrorResult(msg, err)
		}

		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting license details for license code [%v]. Error: [%v]", licenseCode, msg), false)
	}

	return planLicense, result
}

func mapPlanFeatures(planDetails []priceplanmodel.PricePlan) priceplanmodel.PricePlan {

	var features map[string]string
	features = make(map[string]string)

	for _, row := range planDetails {
		features[row.FeatureName] = row.FeatureValue
	}

	pricePlan := priceplanmodel.PricePlan{}

	pricePlan.Id = planDetails[0].Id
	pricePlan.Name = planDetails[0].Name
	pricePlan.DisplayName = planDetails[0].DisplayName
	pricePlan.Price = planDetails[0].Price
	pricePlan.PriceDescription = planDetails[0].PriceDescription
	pricePlan.ImagePath = planDetails[0].ImagePath
	pricePlan.PeriodEnd = planDetails[0].PeriodEnd
	pricePlan.ExpiresInDays = planDetails[0].ExpiresInDays
	pricePlan.Responses, _ = strconv.Atoi(features["Responses"])
	pricePlan.ActiveBoxesLimit, _ = strconv.Atoi(features["Active Boxes"])
	pricePlan.WordCloud, _ = strconv.ParseBool(features["Wordcloud"])
	pricePlan.QuestionBox, _ = strconv.ParseBool(features["Question Box"])
	pricePlan.AnswerBox, _ = strconv.ParseBool(features["Answer Box"])
	pricePlan.VotesBox, _ = strconv.ParseBool(features["Votes Box"])
	pricePlan.ToggleResponses, _ = strconv.ParseBool(features["Toggle Responses"])
	pricePlan.Sms, _ = strconv.ParseBool(features["SMS"])
	pricePlan.CustomCode, _ = strconv.ParseBool(features["Custom Code"])
	pricePlan.Delegates, _ = strconv.Atoi(features["Delegates"])
	pricePlan.Support, _ = strconv.ParseBool(features["Support"])

	return pricePlan
}
