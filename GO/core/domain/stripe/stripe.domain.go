package stripedomain

import (
	"fmt"
	"justasking/GO/common/authenticationclaim"
	"justasking/GO/common/constants/priceplan"
	"justasking/GO/common/operationresult"
	"justasking/GO/core/domain/appconfigs"
	"justasking/GO/core/domain/applogs"
	"justasking/GO/core/domain/email"
	"justasking/GO/core/domain/priceplan"
	"justasking/GO/core/model/emailtemplate"
	"justasking/GO/core/model/userstripe"
	"justasking/GO/core/repo/emailtemplate"
	"justasking/GO/core/repo/stripe"
	"justasking/GO/core/repo/userstripe"
	"time"

	"github.com/stripe/stripe-go"

	uuid "github.com/satori/go.uuid"
)

var domainName = "StripeDomain"

// UpdateSubscription updates a subscription in Stripe
func UpdateSubscription(plan stripe.Plan, tokenClaims *authenticationclaim.AuthenticationClaim) *operationresult.OperationResult {
	functionName := "UpdateSubscription"
	result := operationresult.New()
	var err error

	if tokenClaims.RolePermissions.AccessBilling {
		stripeData, stripeDataResult := GetUserStripeData(tokenClaims.ID, tokenClaims.Account.Id)
		if stripeDataResult.IsSuccess() {
			stripeKey, configsResult := appconfigsdomain.GetAppConfig("stripe", "StripeSecretKey")
			if configsResult.IsSuccess() {
				planId, _ := uuid.FromString(plan.ID)
				planData, planDataResponse := priceplandomain.GetPricePlanDetailsByPlanId(planId)
				var endDatePtr *time.Time
				if planDataResponse.IsSuccess() {

					endDate := time.Now().AddDate(0, 0, planData.ExpiresInDays)
					endDatePtr = &endDate

					err = striperepo.UpdateSubscription(planData, stripeKey.ConfigValue, stripeData, endDatePtr, false)
					if err != nil {
						//send us an email in this case. We might have charged the customer without updating plan data
						msg := fmt.Sprintf("Error updating subscription for user [%v] with stripe customer ID [%v]. Error: [%v]. Check that the Stripe data matches with our database, and that the payment was successful in stripe.",
							tokenClaims.ID, stripeData.StripeUserId, err.Error())
						result = operationresult.CreateErrorResult(msg, err)
						applogsdomain.LogError(domainName, functionName, msg, true)
					} else {
						applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("Updated subscription for user [%v] to [%v] in stripe.", tokenClaims.ID, planData.Name))

						//send confirmation/thanks email
						var confirmationTemplate emailtemplatemodel.EmailTemplate
						if plan.ID == priceplanconstants.PREMIUM_ONEMONTH {
							confirmationTemplate, err = emailtemplaterepo.GetEmailTemplateByName("upgrade_confirmation_month")
						} else if plan.ID == priceplanconstants.PREMIUM_ONEYEAR {
							confirmationTemplate, err = emailtemplaterepo.GetEmailTemplateByName("upgrade_confirmation_year")
						} else if plan.ID == priceplanconstants.PREMIUM_ONEWEEK {
							confirmationTemplate, err = emailtemplaterepo.GetEmailTemplateByName("upgrade_confirmation_week")
						}

						if err != nil {
							applogsdomain.LogError(domainName, functionName, "Unable to retrieve upgrade email template.", false)
						} else {
							confirmationTemplate.To = tokenClaims.Email
							emailSendResult := emaildomain.SendEmail(confirmationTemplate)
							if emailSendResult.IsSuccess() {
								applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("Upgrade email sent to user [%v].", tokenClaims.ID))
							} else {
								applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Unable to send upgrade email to user [%v]. Error: [%v]", tokenClaims.ID, emailSendResult.Message), false)
							}
						}
					}
				} else {
					result.Status = planDataResponse.Status
					result.Message = planDataResponse.Message
					applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting plan data by plan id [%v]. Error: [%v]", plan.ID, planDataResponse.Message), false)
				}
			} else {
				result.Status = configsResult.Status
				result.Message = configsResult.Message
				applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting app configs. Error: [%v]", configsResult.Message), false)
			}
		} else {
			result.Status = stripeDataResult.Status
			result.Message = stripeDataResult.Message
			applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting stripe data for user [%v].", tokenClaims.ID), false)
		}
	} else {
		msg := fmt.Sprintf("Unable to update plan. User [%v] does not have permission.", tokenClaims.ID)
		result.Status = operationresult.Forbidden
		result.Message = msg
		applogsdomain.LogError(domainName, functionName, msg, false)
	}

	return result
}

// UpdateCustomSubscription updates a subscription in Stripe
func UpdateCustomSubscription(licensCode string, tokenClaims *authenticationclaim.AuthenticationClaim) *operationresult.OperationResult {
	functionName := "UpdateCustomSubscription"
	result := operationresult.New()

	if tokenClaims.RolePermissions.AccessBilling {
		planLicense, planLicenseResult := priceplandomain.GetCustomPlanLicense(licensCode)
		if planLicenseResult.IsSuccess() {
			if tokenClaims.ID == planLicense.UserId {
				stripeData, stripeDataResult := GetUserStripeData(tokenClaims.ID, planLicense.AccountId)
				if stripeDataResult.IsSuccess() {
					stripeKey, configsResult := appconfigsdomain.GetAppConfig("stripe", "StripeSecretKey")
					if configsResult.IsSuccess() {
						planData, planDataResponse := priceplandomain.GetPricePlanDetailsByPlanId(planLicense.PlanId)
						var endDatePtr *time.Time
						if planDataResponse.IsSuccess() {
							endDate := time.Now().AddDate(0, 0, planData.ExpiresInDays)
							endDatePtr = &endDate

							stripeData.AccountId = planLicense.AccountId
							err := striperepo.UpdateSubscription(planData, stripeKey.ConfigValue, stripeData, endDatePtr, true)
							if err != nil {
								//send us an email in this case. We might have charged the customer without updating plan data
								msg := fmt.Sprintf("Error updating subscription for user [%v] with stripe customer ID [%v]. Error: [%v]. Check that the Stripe data matches with our database, and that the payment was successful in stripe.",
									tokenClaims.ID, stripeData.StripeUserId, err.Error())
								result = operationresult.CreateErrorResult(msg, err)
								applogsdomain.LogError(domainName, functionName, msg, true)
							} else {
								applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("Updated subscription for user [%v] to [%v] in stripe.", tokenClaims.ID, planData.Name))

								//send confirmation/thanks email
								confirmationTemplate, err := emailtemplaterepo.GetEmailTemplateByName("upgrade_confirmation_custom")

								if err != nil {
									applogsdomain.LogError(domainName, functionName, "Unable to retrieve upgrade email template.", false)
								} else {
									confirmationTemplate.To = tokenClaims.Email
									emailSendResult := emaildomain.SendEmail(confirmationTemplate)
									if emailSendResult.IsSuccess() {
										applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("Upgrade email sent to user [%v].", tokenClaims.ID))
									} else {
										applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Unable to send upgrade email to user [%v]. Error: [%v]", tokenClaims.ID, emailSendResult.Message), false)
									}
								}
							}
						} else {
							result.Status = planDataResponse.Status
							result.Message = planDataResponse.Message
							applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting plan data by plan id [%v]. Error: [%v]", planLicense.PlanId, planDataResponse.Message), false)
						}
					} else {
						result.Status = configsResult.Status
						result.Message = configsResult.Message
						applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting app configs. Error: [%v]", configsResult.Message), false)
					}
				} else {
					result.Status = stripeDataResult.Status
					result.Message = stripeDataResult.Message
					applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting stripe data for user [%v].", tokenClaims.ID), false)
				}
			} else {
				msg := fmt.Sprintf("Unable to update custom subscription. User [%v] does not have access to license [%v]", tokenClaims.ID, licensCode)
				result.Message = msg
				result.Status = operationresult.Forbidden
				applogsdomain.LogError(domainName, functionName, msg, false)
			}
		} else {
			msg := fmt.Sprintf("Unable to update custom subscription. Message: [%v].", planLicenseResult.Message)
			result.Status = operationresult.Error
			result.Message = planLicenseResult.Message
			applogsdomain.LogError(domainName, functionName, msg, false)
		}
	} else {
		msg := fmt.Sprintf("Unable to update plan. User [%v] does not have permission.", tokenClaims.ID)
		result.Status = operationresult.Forbidden
		result.Message = msg
		applogsdomain.LogError(domainName, functionName, msg, false)
	}

	return result
}

// GetUserStripeData gets the mapping between a justasking user and stripe customer
func GetUserStripeData(userId uuid.UUID, accountId uuid.UUID) (userstripemodel.UserStripe, *operationresult.OperationResult) {
	functionName := "GetUserStripeData"
	result := operationresult.New()

	userStripe, err := userstriperepo.GetUserStripeData(userId, accountId)
	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting user stripe mapping for user [%v]. Error: [%v]", userId, msg), false)
	}

	return userStripe, result
}

// GetUserStripeDataByAccountId gets the mapping between a justasking user and stripe customer
func GetUserStripeDataByAccountId(accountId uuid.UUID) (userstripemodel.UserStripe, *operationresult.OperationResult) {
	functionName := "GetUserStripeDataByAccountId"
	var user userstripemodel.UserStripe
	result := operationresult.New()

	userStripe, err := userstriperepo.GetUserStripeDataByAccountId(accountId)
	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting user stripe mapping for account Id  [%v]. Error: [%v]", accountId, msg), false)
	} else {
		if len(userStripe) == 0 {
			result.Status = operationresult.NotFound
			result.Message = fmt.Sprintf("No user stripe mapping found for account [%v]", accountId)
		} else {
			user = userStripe[0]
		}
	}

	return user, result
}

// UpdateCreditCard adds or updates a default credit card for a stripe customer
func UpdateCreditCard(stripeToken stripe.Token, userId uuid.UUID, accountId uuid.UUID) (*stripe.Customer, *operationresult.OperationResult) {
	functionName := "UpdateCreditCard"
	result := operationresult.New()
	var updatedStripeCustomer *stripe.Customer
	var err error

	stripeCustomer, stripeCustomerResult := GetUserStripeData(userId, accountId)
	if stripeCustomerResult.IsSuccess() {
		stripeKey, configsResult := appconfigsdomain.GetAppConfig("stripe", "StripeSecretKey")
		if configsResult.IsSuccess() {
			updatedStripeCustomer, err = striperepo.UpdateCreditCard(stripeCustomer, stripeToken.ID, stripeToken.Card.LastFour, stripeKey.ConfigValue)
			if err != nil {
				msg := fmt.Sprintf("Error updating card for user [%v] with stripe customer ID [%v]. Error: [%v]. Check that the Stripe card data matches with our database.",
					userId, stripeCustomer.StripeUserId, err.Error())
				//send us an email. this could be bad, the card could have been updated in stripe but not in our system.
				// Could cause a problem when we go to create the subscription.
				result = operationresult.CreateErrorResult(msg, err)
				applogsdomain.LogError(domainName, functionName, msg, true)
			} else {
				applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("Successfully updated stripe card and mapping data for user [%v].", userId))
			}
		}
	} else {
		result.Status = stripeCustomerResult.Status
		result.Message = stripeCustomerResult.Message
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting stripe data for user [%v].", userId), false)
	}

	return updatedStripeCustomer, result
}
