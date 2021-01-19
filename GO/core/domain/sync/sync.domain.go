package syncdomain

import (
	"fmt"
	"justasking/GO/common/constants/priceplan"
	"justasking/GO/common/operationresult"
	"justasking/GO/core/domain/account"
	"justasking/GO/core/domain/accountuser"
	"justasking/GO/core/domain/applogs"
	"justasking/GO/core/domain/boxes/basebox"
	"justasking/GO/core/domain/email"
	"justasking/GO/core/domain/phonenumbers"
	"justasking/GO/core/domain/priceplan"
	"justasking/GO/core/domain/user"
	"justasking/GO/core/repo/emailtemplate"
	"time"

	uuid "github.com/satori/go.uuid"
)

var domainName = "SyncDomain"

// CancelExpiredPlans is the main function for canceling plans which have expired
func CancelExpiredPlans() *operationresult.OperationResult {
	functionName := "CancelExpiredPlans"
	result := operationresult.New()

	applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("Starting [%v] function", functionName))

	//1. Get all active accounts which are not on the BASIC plan
	accounts, accountsResponse := accountdomain.GetActiveNonBasicAccounts()
	if accountsResponse.IsSuccess() {
		if len(accounts) > 0 {
			for _, account := range accounts {

				//2. Get stripe data for each account, check expiration date
				accountData, accountDataResponse := priceplandomain.GetPricePlanDetailsByAccountId(account.Id)
				if accountDataResponse.IsSuccess() {

					if time.Now().After(accountData.PeriodEnd) {

						//3. Close boxes
						deactivationResult := baseboxdomain.DeactivateAllBaseBoxesByAccountId(account.Id)
						if deactivationResult.IsSuccess() {

							//4. Update users
							accountUsers, accountUsersResult := accountdomain.GetActiveAccountUsers(account.Id)
							if accountUsersResult.IsSuccess() {

								for _, user := range accountUsers {

									//if it's the owner, just update the token
									if account.OwnerId == user.UserId {
										updateResult := accountuserdomain.UpdateAccountUserToken(user.UserId, account.Id)
										if !updateResult.IsSuccess() {
											msg := fmt.Sprintf("Error updating token for user [%v] during plan expiration for account [%v]", user.UserId, account.Id)
											applogsdomain.LogError(domainName, functionName, msg, true)
										}
									} else {

										//if it's not the owner, switch current_account off of this account and update the token
										updateResult := accountuserdomain.UpdateUserForPlanExpiration(user.UserId, account.Id)
										if !updateResult.IsSuccess() {
											msg := fmt.Sprintf("Error updating user record for user [%v] during plan expiration for account [%v]", user.UserId, account.Id)
											applogsdomain.LogError(domainName, functionName, msg, true)
										}
									}
								}

								//5. Revert to BASIC
								planId, _ := uuid.FromString(priceplanconstants.BASIC)
								planResult := accountdomain.UpdateAccountPricePlan(account.Id, planId)
								if planResult.IsSuccess() {

									//6. Send Email
									user, userResponse := userdomain.GetUser(account.OwnerId)
									if userResponse.IsSuccess() {
										expirationTemplate, err := emailtemplaterepo.GetEmailTemplateByName("plan_expiration")
										if err != nil {
											applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Unable to get email template for syncing expired accounts. Error: [%v]", err), true)
											result.Status = operationresult.Error
										} else {
											expirationTemplate.To = user.Email
											emailResult := emaildomain.SendEmail(expirationTemplate)
											if !emailResult.IsSuccess() {
												applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Unable to sync expired accounts. Error: [%v]", emailResult.Message), true)
												result.Status = operationresult.Error
											}
										}
									} else {
										applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Unable to sync expired accounts. Error: [%v]", userResponse.Message), true)
										result.Status = operationresult.Error
									}
								} else {
									applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Unable to sync expired accounts. Error: [%v]", planResult.Message), true)
									result.Status = operationresult.Error
								}
							} else {
								applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Unable to sync expired accounts. Error: [%v]", accountUsersResult.Message), true)
								result.Status = operationresult.Error
							}
						} else {
							applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Unable to sync expired accounts. Error: [%v]", deactivationResult.Message), true)
							result.Status = operationresult.Error
						}
					} else {
						applogsdomain.LogInfo(domainName, functionName, "No expired accounts to sync.")
					}
				} else {
					applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Unable to sync expired accounts. Error: [%v]", accountDataResponse.Message), true)
					result.Status = operationresult.Error
				}
			}
		} else {
			applogsdomain.LogInfo(domainName, functionName, "No paid accounts to sync.")
		}
	} else {
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Unable to sync expired accounts. Error: [%v]", accountsResponse.Message), true)
		result.Status = operationresult.Error
	}

	applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("Ending [%v] function", functionName))

	return result
}

// SyncPhoneNumbers buys or releases phone numbers based on what we need
func SyncPhoneNumbers(threshold int) *operationresult.OperationResult {
	functionName := "SyncPhoneNumbers"
	result := operationresult.New()

	applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("Starting [%v] function", functionName))

	//1. Calculate how many phone numbers we need

	//1a. Get all active accounts which are not on the BASIC plan
	usablePhoneNumbersNeeded := 0
	totalPhoneNumbersNeeded := 0
	accounts, accountsResponse := accountdomain.GetActiveNonBasicAccounts()
	if accountsResponse.IsSuccess() {
		for _, account := range accounts {

			//1b. Get number of maximum active boxes for each account (one box = one phone number)
			pricePlan, pricePlanResponse := priceplandomain.GetPricePlanDetailsByAccountId(account.Id)
			if pricePlanResponse.IsSuccess() {
				usablePhoneNumbersNeeded += pricePlan.ActiveBoxesLimit
			} else {
				applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Unable to sync phone numbers. Error: [%v]", pricePlanResponse.Message), true)
				result.Status = operationresult.Error
			}
		}
		totalPhoneNumbersNeeded = usablePhoneNumbersNeeded + threshold

		//2. Buy or sell phone numbers based on step 1

		//2a. Count how many phone numbers we have
		ourPhoneNumbers, ourPhoneNumbersResponse := phonenumbersdomain.GetAllActiveJustAskingPhoneNumbers()
		if ourPhoneNumbersResponse.IsSuccess() {
			internalPhoneNumbersCount := len(ourPhoneNumbers)

			applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("Have [%v] phone numbers, need [%v]", internalPhoneNumbersCount, totalPhoneNumbersNeeded))

			// internalPhoneNumbersCount should be equal to totalPhoneNumbersNeeded in the end
			if internalPhoneNumbersCount != totalPhoneNumbersNeeded {
				if internalPhoneNumbersCount < totalPhoneNumbersNeeded {
					//we need to buy numbers
					amountNeeded := totalPhoneNumbersNeeded - internalPhoneNumbersCount

					applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("Need to buy [%v] numbers", amountNeeded))

					//Get available numbers
					availablePhoneNumbers, availablePhoneNumbersResponse := phonenumbersdomain.GetPhoneNumbersForPurchase()
					if availablePhoneNumbersResponse.IsSuccess() {

						availablePhoneNumbersCount := len(availablePhoneNumbers.PhoneNumbers)
						if availablePhoneNumbersCount >= amountNeeded {
							for i := 0; i < amountNeeded; i++ {
								fmt.Println(fmt.Sprintf("Buying number [%v], Phone ID is [%v]", i, availablePhoneNumbers.PhoneNumbers[0].Sid))
								currentPhoneNumber := availablePhoneNumbers.PhoneNumbers[i]
								_, purchasedNumberResponse := phonenumbersdomain.PurchasePhoneNumber(currentPhoneNumber)
								if purchasedNumberResponse.IsSuccess() {
									applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("Successfully purchased phone number [%v]", currentPhoneNumber.PhoneNumber))
								} else {
									applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Unable to purchase phone number [%v] during sync. Error: [%v]", currentPhoneNumber.PhoneNumber, purchasedNumberResponse.Message), true)
								}
							}
						} else {
							applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Not enough phone numbers available for purchase. Numbers needed: [%v]. Numbers available: [%v]", amountNeeded, availablePhoneNumbersCount), true)
						}
					} else {
						applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Unable to buy numbers in sync. Message: [%v]", availablePhoneNumbersResponse.Message), true)
					}
				} else {
					//we need to release numbers
					amountToRelease := internalPhoneNumbersCount - totalPhoneNumbersNeeded
					fmt.Println(fmt.Sprintf("Need to release [%v] numbers", amountToRelease))
					for i := 0; i < amountToRelease; i++ {
						fmt.Println(fmt.Sprintf("Releasing number [%v] of [%v]", i, amountToRelease))
						releasedNumber, releasedNumberResponse := phonenumbersdomain.ReleasePhoneNumber()
						if releasedNumberResponse.IsSuccess() {
							applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("Successfully released phone number [%v] with Sid [%v]", releasedNumber.PhoneNumber, releasedNumber.Sid))
						} else {
							applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Unable to release phone number during sync. Error: [%v]", releasedNumberResponse.Message), true)
						}
					}
				}
			} else {
				//we have the right amount of phone numbers. no need to do anything.
				applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("Correct phone number counts. No need to sync phone numbers."))
			}
		} else {
			applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Unable to sync phone numbers. Error: [%v]", ourPhoneNumbersResponse.Message), true)
			result.Status = operationresult.Error
		}

		//3. Import phone numbers
		importedPhoneNumbers, importedPhoneNumbersResult := phonenumbersdomain.ImportTwilioNumbers()
		if importedPhoneNumbersResult.IsSuccess() {
			applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("Successfully imported phone numbers. Currently have [%v] phone numbers to service [%v] accounts.", len(importedPhoneNumbers.PhoneNumbers), len(accounts)))
		} else {
			applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error importing phone numbers during sync. Error: [%v]", importedPhoneNumbersResult.Message), true)
		}

	} else {
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Unable to sync phone numbers. Error: [%v]", accountsResponse.Message), true)
		result.Status = operationresult.Error
	}

	applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("Ending [%v] function", functionName))

	return result
}
