package accountuserdomain

import (
	"fmt"
	"justasking/GO/common/authenticationclaim"
	"justasking/GO/common/constants/role"
	"justasking/GO/common/operationresult"
	"justasking/GO/core/domain/account"
	"justasking/GO/core/domain/applogs"
	"justasking/GO/core/domain/priceplan"
	"justasking/GO/core/model/accountuser"
	"justasking/GO/core/repo/accountuser"
	"justasking/GO/core/repo/userstripe"

	uuid "github.com/satori/go.uuid"
)

var domainName = "AccountUserDomain"

// UpdateAccountUserRole updates an account user's role
func UpdateAccountUserRole(tokenClaims *authenticationclaim.AuthenticationClaim, accountUser accountusermodel.AccountUser) *operationresult.OperationResult {
	functionName := "UpdateAccountUserRole"
	result := operationresult.New()

	//check that they're trying to update a user for the account they're actually on
	if tokenClaims.Account.Id == accountUser.AccountId {

		//check that they're trying to change someone else's role
		if accountUser.UserId != tokenClaims.ID {

			//check that the user can manage users
			if tokenClaims.RolePermissions.ManageUsers {

				currentAccountUserData, err := accountuserrepo.GetAccountUser(accountUser.UserId, accountUser.AccountId)
				if err != nil {
					msg := fmt.Sprintf("Unable to update accountuser with accountId [%v] and userId [%v]. Error: [%v]", accountUser.AccountId, accountUser.UserId, err.Error())
					result.Status = operationresult.Error
					result.Message = msg
					applogsdomain.LogError(domainName, functionName, msg, false)
				} else {

					//check status of current user
					if currentAccountUserData.IsActive {
						ownerId, _ := uuid.FromString(roleconstants.OWNER)

						//check that they aren't trying to modify Owner status
						if accountUser.RoleId == ownerId || currentAccountUserData.RoleId == ownerId {

							//if we're here, they're trying to change someone's Owner status
							msg := fmt.Sprintf("Unable to update owner. User [%v] does not have permission.", tokenClaims.ID)
							result.Status = operationresult.Forbidden
							result.Message = msg
							applogsdomain.LogError(domainName, functionName, msg, false)
						} else {
							err := accountuserrepo.UpdateAccountUserRole(accountUser, tokenClaims.ID)
							if err != nil {
								msg := fmt.Sprintf("Unable to update accountuser with accountId [%v] and userId [%v]. Error: [%v]", accountUser.AccountId, accountUser.UserId, err.Error())
								result.Status = operationresult.Error
								result.Message = msg
								applogsdomain.LogError(domainName, functionName, msg, false)
							} else {
								msg := fmt.Sprintf("Successfully updated AccountUser. UserId: [%v] AccountId: [%v]", accountUser.UserId, accountUser.AccountId)
								applogsdomain.LogInfo(domainName, functionName, msg)
							}
						}
					} else {
						msg := fmt.Sprintf("Unable to update accountuser with accountId [%v] and userId [%v]. User is not active on this account", accountUser.AccountId, accountUser.UserId)
						result.Status = operationresult.Error
						result.Message = msg
						applogsdomain.LogError(domainName, functionName, msg, false)
					}
				}
			} else {
				msg := fmt.Sprintf("Unable to update account user. User [%v] does not have permission.", tokenClaims.ID)
				result.Status = operationresult.Forbidden
				result.Message = msg
				applogsdomain.LogError(domainName, functionName, msg, false)
			}
		} else {
			msg := fmt.Sprintf("Unable to update account user. Self-update attempted by user [%v].", tokenClaims.ID)
			result.Status = operationresult.Forbidden
			result.Message = msg
			applogsdomain.LogError(domainName, functionName, msg, false)
		}
	} else {
		msg := fmt.Sprintf("Unable to update account user. User on account [%v] can't update user on account [%v].", tokenClaims.Account.Id, accountUser.AccountId)
		result.Status = operationresult.Forbidden
		result.Message = msg
		applogsdomain.LogError(domainName, functionName, msg, false)
	}

	return result
}

// UpdateAccountUserToken updates the token_version for an account user
func UpdateAccountUserToken(userId uuid.UUID, accountId uuid.UUID) *operationresult.OperationResult {
	functionName := "UpdateAccountUserToken"
	result := operationresult.New()

	err := accountuserrepo.UpdateAccountUserToken(userId, accountId)
	if err != nil {
		msg := fmt.Sprintf("Unable to update accountuser token for user with accountId [%v] and userId [%v]. Error: [%v]", accountId, userId, err.Error())
		result.Status = operationresult.Error
		result.Message = msg
		applogsdomain.LogError(domainName, functionName, msg, false)
	}

	return result
}

// TransferOwnership transfers ownership from one user to another
func TransferOwnership(tokenClaims *authenticationclaim.AuthenticationClaim, accountUser accountusermodel.AccountUser) *operationresult.OperationResult {
	functionName := "TransferOwnership"
	result := operationresult.New()

	//check that they're trying to change someone else's role
	if accountUser.UserId != tokenClaims.ID {

		//check that the user is actually on this account. this throws an error if no record is found
		currentAccountUserData, err := accountuserrepo.GetAccountUser(accountUser.UserId, tokenClaims.Account.Id)
		if err != nil {
			msg := fmt.Sprintf("Unable to transfer ownership of account [%v] from user [%v] to user [%v]. Error: [%v]", tokenClaims.Account.Id, tokenClaims.ID, accountUser.UserId, err.Error())
			result.Status = operationresult.Error
			result.Message = msg
			applogsdomain.LogError(domainName, functionName, msg, false)
		} else {

			//check status of current user
			if currentAccountUserData.IsActive {
				accountUserRecords, err := accountuserrepo.GetUserAccounts(tokenClaims.ID)
				if err != nil {
					msg := fmt.Sprintf("Unable to transfer ownership of account [%v] from user [%v] to user [%v]. Error: [%v]", tokenClaims.Account.Id, tokenClaims.ID, accountUser.UserId, err.Error())
					result.Status = operationresult.Error
					result.Message = msg
					applogsdomain.LogError(domainName, functionName, msg, false)
				} else {

					needsNewAccount := true
					//check whether old owner has owner status of a different account
					for _, account := range accountUserRecords {
						if account.RoleName == roleconstants.OWNER_STRING && account.AccountId != accountUser.AccountId {
							needsNewAccount = false
						}
					}

					//check that they can update owners
					if tokenClaims.RolePermissions.ManageOwners {
						oldOwnerStripeData, err := userstriperepo.GetUserStripeDataByUserId(tokenClaims.ID)
						if err != nil {
							msg := fmt.Sprintf("Unable to transfer ownership of account [%v] from user [%v] to user [%v]. Could not get old owner's stripe data. Error: [%v]", tokenClaims.Account.Id, tokenClaims.ID, accountUser.UserId, err.Error())
							result.Status = operationresult.Error
							result.Message = msg
							applogsdomain.LogError(domainName, functionName, msg, false)
						} else {

							newOwnerStripeData, err := userstriperepo.GetUserStripeDataByUserId(accountUser.UserId)
							if err != nil {
								msg := fmt.Sprintf("Unable to transfer ownership of account [%v] from user [%v] to user [%v]. Could not get new owner's stripe data. Error: [%v]", tokenClaims.Account.Id, tokenClaims.ID, accountUser.UserId, err.Error())
								result.Status = operationresult.Error
								result.Message = msg
								applogsdomain.LogError(domainName, functionName, msg, false)
							} else {
								newAccountName := tokenClaims.FirstName + " " + tokenClaims.LastName
								err := accountuserrepo.TransferOwnership(tokenClaims.ID, accountUser.UserId, tokenClaims.Account.Id, newOwnerStripeData.StripeUserId, newAccountName, needsNewAccount, oldOwnerStripeData.StripeUserId)
								if err != nil {
									msg := fmt.Sprintf("Unable to transfer ownership of account [%v] from user [%v] to user [%v]. Error: [%v]", tokenClaims.Account.Id, tokenClaims.ID, accountUser.UserId, err.Error())
									result.Status = operationresult.Error
									result.Message = msg
									applogsdomain.LogError(domainName, functionName, msg, false)
								} else {
									msg := fmt.Sprintf("User [%v] successfully transferred ownership of account [%v] to user [%v]", tokenClaims.ID, tokenClaims.Account.Id, accountUser.UserId)
									applogsdomain.LogInfo(domainName, functionName, msg)
								}
							}
						}
					} else {
						msg := fmt.Sprintf("Unable to transfer account ownership. User [%v] does not have permission.", tokenClaims.ID)
						result.Status = operationresult.Forbidden
						result.Message = msg
						applogsdomain.LogError(domainName, functionName, msg, false)
					}
				}
			} else {
				msg := fmt.Sprintf("Unable to transfer account ownership. User [%v] is not active on account [%v].", accountUser.UserId, tokenClaims.Account.Id)
				result.Status = operationresult.Forbidden
				result.Message = msg
				applogsdomain.LogError(domainName, functionName, msg, false)
			}
		}
	} else {
		msg := fmt.Sprintf("Unable to transfer account ownership. Self-update attempted by user [%v].", tokenClaims.ID)
		result.Status = operationresult.Forbidden
		result.Message = msg
		applogsdomain.LogError(domainName, functionName, msg, false)
	}

	return result
}

// RemoveUserFromAccount removes a user from an account
func RemoveUserFromAccount(tokenClaims *authenticationclaim.AuthenticationClaim, accountUser accountusermodel.AccountUser) *operationresult.OperationResult {
	functionName := "RemoveUserFromAccount"
	result := operationresult.New()

	//check that they're trying to remove a different person
	if accountUser.UserId != tokenClaims.ID {

		//check that the user is actually on this account. this throws an error if no record is found
		currentUserData, err := accountuserrepo.GetAccountUser(accountUser.UserId, tokenClaims.Account.Id)
		if err != nil {
			msg := fmt.Sprintf("Unable to remove user [%v] from account [%v]. Error: [%v]", accountUser.UserId, tokenClaims.Account.Id, err.Error())
			result.Status = operationresult.Error
			result.Message = msg
			applogsdomain.LogError(domainName, functionName, msg, false)
		} else {

			//check status of current user
			if currentUserData.IsActive {

				//check that the person doing the removing has the right permissions
				if tokenClaims.RolePermissions.ManageUsers {

					//check that the user being removed is not the owner
					currentUserRoleId, _ := uuid.FromString(roleconstants.OWNER)
					if currentUserData.RoleId != currentUserRoleId {

						//check whether the user needs current_account updated
						currentAccount, err := accountuserrepo.GetCurrentAccount(accountUser.UserId)
						if err != nil {
							msg := fmt.Sprintf("Unable to remove user [%v] from account [%v]. Error: [%v]", accountUser.UserId, tokenClaims.Account.Id, err.Error())
							result.Status = operationresult.Error
							result.Message = msg
							applogsdomain.LogError(domainName, functionName, msg, false)
						} else {
							needsCurrentAccountUpdated := currentAccount.AccountId == accountUser.AccountId
							err := accountuserrepo.RemoveUserFromAccount(accountUser.UserId, accountUser.AccountId, tokenClaims.ID, needsCurrentAccountUpdated)
							if err != nil {
								msg := fmt.Sprintf("Unable to remove user [%v] from account [%v]. Error: [%v]", accountUser.UserId, tokenClaims.Account.Id, err.Error())
								result.Status = operationresult.Error
								result.Message = msg
								applogsdomain.LogError(domainName, functionName, msg, false)
							} else {
								msg := fmt.Sprintf("User [%v] successfully removed user [%v] from account [%v].", tokenClaims.ID, accountUser.UserId, accountUser.AccountId)
								applogsdomain.LogInfo(domainName, functionName, msg)
							}
						}
					} else {
						msg := fmt.Sprintf("Unable to remove owner with userid [%v] from account [%v].", accountUser.UserId, tokenClaims.Account.Id)
						result.Status = operationresult.Forbidden
						result.Message = msg
						applogsdomain.LogError(domainName, functionName, msg, false)
					}
				} else {
					msg := fmt.Sprintf("Unable to remove user [%v] from account [%v]. User [%v] does not have permission.", accountUser.UserId, tokenClaims.Account.Id, tokenClaims.ID)
					result.Status = operationresult.Forbidden
					result.Message = msg
					applogsdomain.LogError(domainName, functionName, msg, false)
				}
			} else {
				msg := fmt.Sprintf("Unable to remove user [%v] from account [%v]. User is not active on the account.", accountUser.UserId, tokenClaims.Account.Id)
				result.Status = operationresult.Forbidden
				result.Message = msg
				applogsdomain.LogError(domainName, functionName, msg, false)
			}
		}
	} else {
		msg := fmt.Sprintf("Unable to remove user from account. Self-removal attempted by user [%v].", tokenClaims.ID)
		result.Status = operationresult.Forbidden
		result.Message = msg
		applogsdomain.LogError(domainName, functionName, msg, false)
	}

	return result
}

// UpdateUserForPlanExpiration updates a user's current account and token_version for an account that recently expired
func UpdateUserForPlanExpiration(userId uuid.UUID, accountId uuid.UUID) *operationresult.OperationResult {
	functionName := "UpdateUserForPlanExpiration"
	result := operationresult.New()

	//check whether the user needs current_account updated
	currentAccount, err := accountuserrepo.GetCurrentAccount(userId)
	if err != nil {
		msg := fmt.Sprintf("Unable to get current account for updating user [%v] for expiration of account [%v]. Error: [%v]", userId, accountId, err.Error())
		result.Status = operationresult.Error
		result.Message = msg
		applogsdomain.LogError(domainName, functionName, msg, false)
	} else {
		needsCurrentAccountUpdated := currentAccount.AccountId == accountId
		err := accountuserrepo.UpdateUserForPlanExpiration(userId, accountId, needsCurrentAccountUpdated)
		if err != nil {
			msg := fmt.Sprintf("Unable to update user [%v] for expiration of account [%v]. Error: [%v]", userId, accountId, err.Error())
			result.Status = operationresult.Error
			result.Message = msg
			applogsdomain.LogError(domainName, functionName, msg, false)
		}
	}

	return result
}

// UpdateCurrentAccount updates the user's current account
func UpdateCurrentAccount(userId uuid.UUID, accountId uuid.UUID) *operationresult.OperationResult {
	functionName := "UpdateCurrentAccount"
	result := operationresult.New()

	canUpdateAccount := false

	userBelongsToAccount := accountdomain.UserBelongsToAccount(userId, accountId)
	if userBelongsToAccount {

		//check if this is the owner, and whether the account has delegates if not
		accountDetails, accountDetailsResult := accountdomain.GetAccount(accountId)
		if accountDetailsResult.IsSuccess() {

			if accountDetails.OwnerId != userId {

				// account needs delegates if this person isn't the owner
				planDetails, planDetailsResult := priceplandomain.GetPricePlanDetailsByAccountId(accountId)
				if planDetailsResult.IsSuccess() {
					if planDetails.Delegates > 0 {
						canUpdateAccount = true
					}
				} else {
					msg := fmt.Sprintf("Unable to update current account to [%v] for user [%v]. Error getting account details: [%v].", accountId, userId, planDetailsResult.Message)
					result.Status = operationresult.Error
					result.Message = msg
					applogsdomain.LogError(domainName, functionName, msg, false)
				}
			} else {
				canUpdateAccount = true
			}

			if canUpdateAccount {
				err := accountuserrepo.UpdateCurrentAccount(userId, accountId)
				if err != nil {
					msg := err.Error()
					result = operationresult.CreateErrorResult(msg, err)
					applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Unable to update current account to [%v] for user [%v]. Error: [%v]", accountId, userId, err.Error()), false)
				} else {
					msg := fmt.Sprintf("Current account for user [%v] updated to [%v]", userId, accountId)
					applogsdomain.LogInfo(domainName, functionName, msg)
				}
			} else {
				msg := fmt.Sprintf("Unable to update current account to [%v] for user [%v]. User does not have access.", accountId, userId)
				result.Status = operationresult.Forbidden
				result.Message = msg
				applogsdomain.LogError(domainName, functionName, msg, false)
			}
		} else {
			msg := fmt.Sprintf("Unable to update current account to [%v] for user [%v]. Error getting account details: [%v].", accountId, userId, accountDetailsResult.Message)
			result.Status = operationresult.Forbidden
			result.Message = msg
			applogsdomain.LogError(domainName, functionName, msg, false)
		}
	} else {
		msg := fmt.Sprintf("Unable to update current account to [%v] for user [%v]. User does not belong to account.", accountId, userId)
		result.Status = operationresult.Forbidden
		result.Message = msg
		applogsdomain.LogError(domainName, functionName, msg, false)
	}

	return result
}

// GetAccountUser gets an accountUser record
func GetAccountUser(userId uuid.UUID, accountId uuid.UUID) (accountusermodel.AccountUser, *operationresult.OperationResult) {
	functionName := "GetAccountUser"
	result := operationresult.New()

	accountUser, err := accountuserrepo.GetAccountUser(userId, accountId)
	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Unable to retrieve account user with userId [%v] and accountId [%v]. Error: [%v]", userId, accountId, err.Error()), false)
	}

	return accountUser, result
}
