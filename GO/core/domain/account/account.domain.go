package accountdomain

import (
	"bytes"
	"fmt"
	authenticationclaim "justasking/GO/common/authenticationclaim"
	"justasking/GO/common/constants/role"
	"justasking/GO/common/operationresult"
	"justasking/GO/common/utility"
	"justasking/GO/core/domain/applogs"
	"justasking/GO/core/domain/email"
	"justasking/GO/core/domain/priceplan"
	"justasking/GO/core/domain/user"
	"justasking/GO/core/model/account"
	"justasking/GO/core/model/accountinvitation"
	"justasking/GO/core/model/accountuser"
	"justasking/GO/core/repo/account"
	"justasking/GO/core/repo/accountuser"
	"justasking/GO/core/repo/emailtemplate"
	"regexp"
	"strings"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

var domainName = "AccountDomain"

// GetAccount gets an account by account Id
func GetAccount(accountId uuid.UUID) (accountmodel.Account, *operationresult.OperationResult) {
	functionName := "GetAccount"
	result := operationresult.New()

	account, err := accountrepo.GetAccount(accountId)
	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting account [%v]. Error: [%v]", accountId, msg), false)
	}

	return account, result
}

// UpdateAccount is for updating an account
func UpdateAccount(account accountmodel.Account, userId uuid.UUID) (accountmodel.Account, *operationresult.OperationResult) {
	functionName := "UpdateAccount"
	result := operationresult.New()
	var updatedAccount accountmodel.Account
	var err error

	userBelongsToAccount := UserBelongsToAccount(userId, account.Id)

	if userBelongsToAccount {
		updatedAccount, err = accountrepo.UpdateAccount(account)

		if err != nil {
			msg := err.Error()
			result = operationresult.CreateErrorResult(msg, err)
			applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error updating account id [%v]. Error: [%v]", account.Id, msg), false)
		} else {
			applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("Account [%v] updated by user [%v].", account.Id, userId))
		}
	} else {
		msg := fmt.Sprintf("Unable to update account. User [%v] does not belong to account [%v]", userId, account.Id)
		result.Message = msg
		result.Status = operationresult.Error
		applogsdomain.LogError(domainName, functionName, msg, false)
	}

	return updatedAccount, result
}

// GetAccountByAnswerBoxEntry returns the account associated with the box associated with the question associated with the entry
func GetAccountByAnswerBoxEntry(entryId uuid.UUID) (accountmodel.Account, *operationresult.OperationResult) {
	functionName := "GetAccountByAnswerBoxEntry"

	result := operationresult.New()

	account, err := accountrepo.GetAccountByAnswerBoxEntry(entryId)
	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting account for answer box entry [%v]. Error: [%v]", entryId, msg), false)
	}

	return account, result
}

// GetAccountByAnswerBoxQuestion returns the account associated with the box associated with the question
func GetAccountByAnswerBoxQuestion(questionId uuid.UUID) (accountmodel.Account, *operationresult.OperationResult) {
	functionName := "GetAccountByAnswerBoxQuestion"

	result := operationresult.New()

	account, err := accountrepo.GetAccountByAnswerBoxQuestion(questionId)
	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting account for answer box question [%v]. Error: [%v]", questionId, msg), false)
	}

	return account, result
}

// GetAccountByQuestionBoxEntry returns the account associated with the box associated with the question associated with the entry
func GetAccountByQuestionBoxEntry(entryId uuid.UUID) (accountmodel.Account, *operationresult.OperationResult) {
	functionName := "GetAccountByQuestionBoxEntry"

	result := operationresult.New()

	account, err := accountrepo.GetAccountByQuestionBoxEntry(entryId)
	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting account for question box entry [%v]. Error: [%v]", entryId, msg), false)
	}

	return account, result
}

// GetAccountByBoxId returns the account associated with the box
func GetAccountByBoxId(boxId uuid.UUID) (accountmodel.Account, *operationresult.OperationResult) {
	functionName := "GetAccountByBoxId"

	result := operationresult.New()

	account, err := accountrepo.GetAccountByBoxId(boxId)
	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting account for box [%v]. Error: [%v]", boxId, msg), false)
	}

	return account, result
}

// GetAccountByCode returns the account associated with the box
func GetAccountByCode(code string) (accountmodel.Account, *operationresult.OperationResult) {
	functionName := "GetAccountByCode"

	result := operationresult.New()

	account, err := accountrepo.GetAccountByCode(code)
	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting account for box [%v]. Error: [%v]", code, msg), false)
	}

	return account, result
}

// UpdateAccountPricePlan updates an account/priceplan mapping
func UpdateAccountPricePlan(accountId uuid.UUID, priceplanId uuid.UUID) *operationresult.OperationResult {
	functionName := "UpdateAccountPricePlan"
	result := operationresult.New()

	err := accountrepo.UpdateAccountPricePlan(accountId, priceplanId)
	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error updating price plan for account [%v]. Error: [%v]", accountId, msg), true)
	} else {
		applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("Updated price plan to [%v] for account [%v].", priceplanId, accountId))
	}

	return result
}

// UserBelongsToAccount returns true if the user belongs to the account, false otherwise
func UserBelongsToAccount(userId uuid.UUID, accountId uuid.UUID) bool {
	functionName := "UserBelongsToAccount"
	found := false

	userAccount, err := accountuserrepo.GetAccountUser(userId, accountId)

	if err != nil {
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting accounts by userid [%v]. Error: [%v]", accountId, err.Error()), false)
	} else {
		if userAccount.IsActive {
			found = true
		}
	}

	return found
}

//GetActiveNonBasicAccounts gets all active accounts which are not on the BASIC plan
func GetActiveNonBasicAccounts() ([]accountmodel.Account, *operationresult.OperationResult) {
	functionName := "GetActiveNonBasicAccounts"
	result := operationresult.New()

	accounts, err := accountrepo.GetActiveNonBasicAccounts()
	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting active paid accounts. Error: [%v]", err.Error()), false)
	}

	return accounts, result
}

// InviteAccountUser creates an account invitation
func InviteAccountUser(tokenClaims *authenticationclaim.AuthenticationClaim, accountInvitation accountinvitationmodel.AccountInvitation) *operationresult.OperationResult {
	functionName := "InviteAccountUser"
	result := operationresult.New()

	//get plan details for account
	planDetails, planDetailsResult := priceplandomain.GetPricePlanDetailsByAccountId(tokenClaims.Account.Id)
	if planDetailsResult.IsSuccess() {

		//check whether account can add users
		if planDetails.Delegates <= 0 {
			msg := fmt.Sprintf("Unable to invite users on the [%v] plan with id [%v].", planDetails.Name, planDetails.Id)
			result.Status = operationresult.PaymentRequired
			applogsdomain.LogError(domainName, functionName, msg, false)
		} else {

			//get account users
			accountUsers, err := accountuserrepo.GetActiveAccountUsers(tokenClaims.Account.Id)
			if err != nil {
				msg := err.Error()
				result = operationresult.CreateErrorResult(msg, err)
				applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error creating account invite. Error: [%v]", err.Error()), false)
			} else {

				//get account invites
				accountInvites, err := accountrepo.GetActiveAccountInvites(tokenClaims.Account.Id)

				if err != nil {
					msg := err.Error()
					result = operationresult.CreateErrorResult(msg, err)
					applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error creating account invite. Error: [%v]", err.Error()), false)
				} else {

					//check if user has an active invitation for this account
					var alreadyInvited bool
					for _, invite := range accountInvites {
						if invite.Email == accountInvitation.Email {
							alreadyInvited = true
						}
					}

					//check if user already exists on this account
					var alreadyExists bool
					for _, accountUser := range accountUsers {
						user, userResult := userdomain.GetUser(accountUser.UserId)
						if userResult.IsSuccess() {
							if user.Email == accountInvitation.Email {
								alreadyExists = true
							}
						} else {
							continue
						}
					}

					if alreadyInvited {
						msg := fmt.Sprintf("User [%v] has already been invited to account [%v].", accountInvitation.Email, accountInvitation.AccountId)
						result.Status = operationresult.Conflict
						applogsdomain.LogError(domainName, functionName, msg, false)
					} else if alreadyExists {
						msg := fmt.Sprintf("User [%v] is already on account [%v].", accountInvitation.Email, accountInvitation.AccountId)
						result.Status = operationresult.Conflict
						applogsdomain.LogError(domainName, functionName, msg, false)
					} else {

						//check whether account can still add users
						// subtracting one to account for account owner
						if (len(accountUsers)+len(accountInvites))-1 < planDetails.Delegates {

							inviteIsValid, inviteIsValidMessage := validateAccountInvite(accountInvitation)
							if inviteIsValid {
								invitationCode := utility.RandSeq(256)
								err := accountrepo.CreateAccountInvite(accountInvitation, invitationCode, tokenClaims.Account.Id, tokenClaims.ID)
								if err != nil {
									msg := err.Error()
									result = operationresult.CreateErrorResult(msg, err)
									applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error creating account invite. Error: [%v]", err.Error()), false)
								} else {
									//send invitation email
									accountInvitationTemplate, err := emailtemplaterepo.GetEmailTemplateByName("account_invitation")
									if err != nil {
										applogsdomain.LogError(domainName, functionName, "Unable to retrieve account_invitation email template.", false)
									} else {
										invitationSender := tokenClaims.FirstName + " " + tokenClaims.LastName
										accountInvitationTemplate.To = accountInvitation.Email
										accountInvitationTemplate.Body = strings.Replace(accountInvitationTemplate.Body, "{resetCode}", invitationCode, -1)
										accountInvitationTemplate.Body = strings.Replace(accountInvitationTemplate.Body, "{InvitationSender}", invitationSender, -1)
										accountInvitationTemplate.Body = strings.Replace(accountInvitationTemplate.Body, "{InvitationSenderEmail}", tokenClaims.Email, -1)

										emailSendResult := emaildomain.SendEmail(accountInvitationTemplate)
										if emailSendResult.IsSuccess() {
											applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("account_invitation email sent to [%v].", accountInvitation.Email))
										} else {
											applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Unable to send account_invitation email to [%v]. Error: [%v]", accountInvitation.Email, emailSendResult.Message), true)
										}
									}
								}
							} else {
								msg := fmt.Sprintf("Could not create invitation. Error: [%v]", inviteIsValidMessage)
								result.Message = msg
								result.Status = operationresult.UnprocessableEntity
								applogsdomain.LogInfo(domainName, functionName, msg)
							}
						} else {
							msg := fmt.Sprintf("Account has reached maximum number of users.")
							result.Status = operationresult.Forbidden
							applogsdomain.LogError(domainName, functionName, msg, false)
						}
					}
				}
			}
		}
	} else {
		msg := fmt.Sprintf("Unable to invite user. Error: [%v]", planDetailsResult.Message)
		result.Status = operationresult.Error
		applogsdomain.LogError(domainName, functionName, msg, false)
	}

	return result
}

// ResendAccountUserInvitation creates an account invitation
func ResendAccountUserInvitation(tokenClaims *authenticationclaim.AuthenticationClaim, accountInvitation accountinvitationmodel.AccountInvitation) *operationresult.OperationResult {
	functionName := "ResendAccountUserInvitation"
	result := operationresult.New()

	//get plan details for account
	planDetails, planDetailsResult := priceplandomain.GetPricePlanDetailsByAccountId(tokenClaims.Account.Id)
	if planDetailsResult.IsSuccess() {

		//check whether account can add users
		if planDetails.Delegates <= 0 {
			msg := fmt.Sprintf("Unable to invite users on the [%v] plan with id [%v].", planDetails.Name, planDetails.Id)
			result.Status = operationresult.PaymentRequired
			applogsdomain.LogError(domainName, functionName, msg, false)
		} else {

			//get account users
			accountUsers, err := accountuserrepo.GetActiveAccountUsers(tokenClaims.Account.Id)
			if err != nil {
				msg := err.Error()
				result = operationresult.CreateErrorResult(msg, err)
				applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error creating account invite. Error: [%v]", err.Error()), false)
			} else {

				//get account invites
				accountInvites, err := accountrepo.GetActiveAccountInvites(tokenClaims.Account.Id)

				if err != nil {
					msg := err.Error()
					result = operationresult.CreateErrorResult(msg, err)
					applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error creating account invite. Error: [%v]", err.Error()), false)
				} else {

					//check if user has an active invitation for this account
					var alreadyInvited bool
					for _, invite := range accountInvites {
						if invite.Email == accountInvitation.Email {
							alreadyInvited = true
						}
					}

					//check if user already exists on this account
					var alreadyExists bool
					for _, accountUser := range accountUsers {
						user, userResult := userdomain.GetUser(accountUser.UserId)
						if userResult.IsSuccess() {
							if user.Email == accountInvitation.Email {
								alreadyExists = true
							}
						} else {
							continue
						}
					}

					if alreadyInvited && !alreadyExists {

						existingAccountInvitation, err := accountrepo.GetInvitationByAccountAndEmail(tokenClaims.Account.Id, accountInvitation.Email)
						if err != nil {
							msg := err.Error()
							result = operationresult.CreateErrorResult(msg, err)
							applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error creating account invite. Error: [%v]", err.Error()), false)
						} else {
							invitationCode := existingAccountInvitation.InvitationCode

							if len(invitationCode) > 0 {
								//send invitation email
								accountInvitationTemplate, err := emailtemplaterepo.GetEmailTemplateByName("account_invitation")
								if err != nil {
									applogsdomain.LogError(domainName, functionName, "Unable to retrieve account_invitation email template.", false)
								} else {
									invitationSender := tokenClaims.FirstName + " " + tokenClaims.LastName
									accountInvitationTemplate.To = accountInvitation.Email
									accountInvitationTemplate.Body = strings.Replace(accountInvitationTemplate.Body, "{resetCode}", invitationCode, -1)
									accountInvitationTemplate.Body = strings.Replace(accountInvitationTemplate.Body, "{InvitationSender}", invitationSender, -1)
									accountInvitationTemplate.Body = strings.Replace(accountInvitationTemplate.Body, "{InvitationSenderEmail}", tokenClaims.Email, -1)

									emailSendResult := emaildomain.SendEmail(accountInvitationTemplate)
									if emailSendResult.IsSuccess() {
										applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("account_invitation email sent to [%v].", accountInvitation.Email))
									} else {
										applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Unable to send account_invitation email to [%v]. Error: [%v]", accountInvitation.Email, emailSendResult.Message), true)
									}
								}
							} else {
								msg := fmt.Sprintf("Could not resend invitation to account. User [%v] did not have an active invitation to account [%v].", accountInvitation.Email, accountInvitation.AccountId)
								result.Status = operationresult.Conflict
								applogsdomain.LogError(domainName, functionName, msg, false)
							}
						}

					} else if !alreadyInvited {
						msg := fmt.Sprintf("User [%v] has not been invited to account before resending invitation to account [%v].", accountInvitation.Email, accountInvitation.AccountId)
						result.Status = operationresult.Conflict
						applogsdomain.LogError(domainName, functionName, msg, false)
					} else if alreadyExists {
						msg := fmt.Sprintf("User [%v] is already on account [%v].", accountInvitation.Email, accountInvitation.AccountId)
						result.Status = operationresult.Conflict
						applogsdomain.LogError(domainName, functionName, msg, false)
					} else {
						msg := fmt.Sprintf("Account invitation could not be sent to user [%v] for account [%v] because invitation does not exist.", accountInvitation.Email, accountInvitation.AccountId)
						result.Status = operationresult.Conflict
						applogsdomain.LogError(domainName, functionName, msg, false)
					}
				}
			}
		}
	} else {
		msg := fmt.Sprintf("Unable to resend invitation to user user. Error: [%v]", planDetailsResult.Message)
		result.Status = operationresult.Error
		applogsdomain.LogError(domainName, functionName, msg, false)
	}

	return result
}

// GetInvitationForJoin retrieves a pending invitation's data
func GetInvitationForJoin(tokenClaims *authenticationclaim.AuthenticationClaim, invitationCode string) (accountinvitationmodel.AccountInvitation, *operationresult.OperationResult) {
	functionName := "GetInvitationForJoin"
	result := operationresult.New()
	var err error
	var fullInvitation accountinvitationmodel.AccountInvitation

	fullInvitation, err = accountrepo.GetInvitation(invitationCode)
	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting account invite. Error: [%v]", err.Error()), false)
	} else if fullInvitation.IsActive == false {
		msg := fmt.Sprintf("Invitation with code [%v] is no longer active.", invitationCode)
		result.Status = operationresult.Gone
		result.Message = msg
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting account invite. Error: [%v]", msg), false)
	} else if strings.ToLower(fullInvitation.Email) != strings.ToLower(tokenClaims.Email) {
		msg := "Invalid email."
		result.Status = operationresult.Forbidden
		result.Message = msg
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting account invite. Error: [%v]", msg), false)
	}

	return fullInvitation, result
}

// JoinAccount adds a user to an account
func JoinAccount(tokenClaims *authenticationclaim.AuthenticationClaim, invitationCode string) *operationresult.OperationResult {
	functionName := "JoinAccount"
	result := operationresult.New()

	invitation, invitationResult := GetInvitationForJoin(tokenClaims, invitationCode)
	if invitationResult.IsSuccess() {
		if strings.ToLower(tokenClaims.Email) != strings.ToLower(invitation.Email) {
			msg := "Unable to add user to account. User email does not match invitation"
			result.Status = operationresult.Forbidden
			result.Message = msg
			applogsdomain.LogError(domainName, functionName, msg, false)
		} else {

			//get plan details for account
			planDetails, planDetailsResult := priceplandomain.GetPricePlanDetailsByAccountId(invitation.AccountId)
			if planDetailsResult.IsSuccess() {

				//check whether account can add users
				if planDetails.Delegates <= 0 {
					msg := fmt.Sprintf("Unable to invite users on the [%v] plan with id [%v].", planDetails.Name, planDetails.Id)
					result.Status = operationresult.PaymentRequired
					result.Message = msg
					applogsdomain.LogError(domainName, functionName, msg, false)
				} else {
					//get account users
					accountUsers, err := accountuserrepo.GetActiveAccountUsers(invitation.AccountId)
					if err != nil {
						msg := err.Error()
						result = operationresult.CreateErrorResult(msg, err)
						applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error adding user to account. Error: [%v]", err.Error()), false)
					} else {

						//subtracting one to account for owner
						if len(accountUsers)-1 == planDetails.ActiveBoxesLimit {
							msg := fmt.Sprintf("Account [%v] has reached the limit and canot add users.", invitation.AccountId)
							result.Status = operationresult.Forbidden
							result.Message = msg
							applogsdomain.LogError(domainName, functionName, msg, false)
						} else {

							//check if user is already active on this account
							var alreadyActive bool
							for _, accountUser := range accountUsers {
								user, userResult := userdomain.GetUser(accountUser.UserId)
								if userResult.IsSuccess() {
									if user.Email == invitation.Email {
										alreadyActive = true
									}
								} else {
									continue
								}
							}

							if alreadyActive {
								msg := fmt.Sprintf("User [%v] is already active on account [%v] and cannot be added.", invitation.Email, invitation.AccountId)
								result.Status = operationresult.Forbidden
								result.Message = msg
								applogsdomain.LogError(domainName, functionName, msg, false)
							} else {

								_, err := accountuserrepo.GetAccountUser(tokenClaims.ID, invitation.AccountId)
								if err != nil && err != gorm.ErrRecordNotFound {
									msg := err.Error()
									result = operationresult.CreateErrorResult(msg, err)
									applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error adding user to account. Error: [%v]", err.Error()), false)
								} else {

									var userAccountExists bool
									if err != nil && err == gorm.ErrRecordNotFound {
										//the user actually does not exist on this account
										userAccountExists = false
									} else {
										//the user exists but is inactive on this
										userAccountExists = true
									}

									//deactivate invitation, add user to account
									err := accountuserrepo.RedeemAccountInvitation(invitation, tokenClaims.ID, userAccountExists)
									if err != nil {
										msg := err.Error()
										result = operationresult.CreateErrorResult(msg, err)
										applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error adding user to account. Error: [%v]", err.Error()), false)
									} else {
										msg := fmt.Sprintf("User [%v] successfully added to account [%v]", tokenClaims.ID, invitation.AccountId)
										applogsdomain.LogInfo(domainName, functionName, msg)
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		msg := fmt.Sprintf("Unable to add user to account. Error: [%v]", invitationResult.Message)
		result.Status = invitationResult.Status
		result.Message = msg
		applogsdomain.LogError(domainName, functionName, msg, false)
	}

	return result
}

// GetActiveAndPendingAccountUsers gets all active and pending users for an account
func GetActiveAndPendingAccountUsers(accountId uuid.UUID) ([]accountusermodel.AccountUser, *operationresult.OperationResult) {
	functionName := "GetActiveAndPendingAccountUsers"
	result := operationresult.New()

	var accountUsers []accountusermodel.AccountUser

	activeUsers, err := accountuserrepo.GetActiveAccountUsers(accountId)
	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting all account users. Error: [%v]", err.Error()), false)
	} else {
		for _, activeUser := range activeUsers {
			user, _ := accountuserrepo.GetAccountUser(activeUser.UserId, accountId)
			activeUser.Email = user.Email
			activeUser.IsPending = false
			activeUser.RoleName = user.RoleName
			accountUsers = append(accountUsers, activeUser)
		}

		pendingUsers, err := accountrepo.GetActiveAccountInvites(accountId)
		if err != nil {
			msg := err.Error()
			result = operationresult.CreateErrorResult(msg, err)
			applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting pending account users. Error: [%v]", err.Error()), false)
		} else {
			for _, pendingUser := range pendingUsers {

				var invite accountusermodel.AccountUser
				invite.AccountId = pendingUser.AccountId
				invite.RoleId = pendingUser.RoleId
				invite.IsPending = true
				invite.Email = pendingUser.Email
				invite.RoleName = pendingUser.RoleName

				accountUsers = append(accountUsers, invite)
			}
		}
	}

	return accountUsers, result
}

// GetActiveAccountUsers gets active account users
func GetActiveAccountUsers(accountId uuid.UUID) ([]accountusermodel.AccountUser, *operationresult.OperationResult) {
	functionName := "GetActiveAccountUsers"
	result := operationresult.New()

	accountUsers, err := accountuserrepo.GetActiveAccountUsers(accountId)
	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting all active account users. Error: [%v]", err.Error()), false)
	}

	return accountUsers, result
}

// GetAccounts gets all accounts that a user belongs to
func GetAccounts(userId uuid.UUID) ([]accountusermodel.AccountUser, *operationresult.OperationResult) {
	functionName := "GetAccounts"
	result := operationresult.New()
	var accounts []accountusermodel.AccountUser

	allAccounts, err := accountuserrepo.GetUserAccounts(userId)
	if err != nil {
		msg := fmt.Sprintf("Unable to get accounts for user [%v]. Error: [%v]", userId, err.Error())
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, msg, false)
	} else {
		for _, account := range allAccounts {
			accountData, accountDataResult := GetAccount(account.AccountId)
			if accountDataResult.IsSuccess() {
				if accountData.OwnerId == userId {
					accounts = append(accounts, account)
				} else {
					planData, planDataResult := priceplandomain.GetPricePlanDetailsByAccountId(account.AccountId)
					if planDataResult.IsSuccess() {
						if planData.Delegates > 0 {
							accounts = append(accounts, account)
						}
					}
				}
			}
		}
	}

	return accounts, result
}

// CancelInvitation cancels an account invitation
func CancelInvitation(tokenClaims *authenticationclaim.AuthenticationClaim, accountInvitation accountinvitationmodel.AccountInvitation) *operationresult.OperationResult {
	functionName := "CancelInvitation"
	result := operationresult.New()

	if tokenClaims.RolePermissions.ManageUsers {
		fullInvitation, err := accountrepo.GetInvitationForUpdate(accountInvitation)
		if err != nil {
			msg := err.Error()
			result = operationresult.CreateErrorResult(msg, err)
			applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting account invite. Error: [%v]", err.Error()), false)
		} else {
			if fullInvitation.AccountId != tokenClaims.Account.Id {
				msg := fmt.Sprintf("Unable to cancel invitation. User [%v] on account [%v] tried to cancel invitation for account [%v]", tokenClaims.ID, tokenClaims.Account.Id, fullInvitation.AccountId)
				result.Status = operationresult.Forbidden
				result.Message = msg
				applogsdomain.LogError(domainName, functionName, msg, false)
			} else {
				err := accountrepo.CancelInvitation(fullInvitation.InvitationCode, tokenClaims.ID)
				if err != nil {
					msg := fmt.Sprintf("Unable to cancel invitation. Error: [%v]", err.Error())
					result.Status = operationresult.Error
					result.Message = msg
					applogsdomain.LogError(domainName, functionName, msg, false)
				} else {
					msg := fmt.Sprintf("Account invitation [%v] has been canceled by user [%v]", fullInvitation.InvitationCode, tokenClaims.ID)
					applogsdomain.LogInfo(domainName, functionName, msg)
				}
			}
		}
	} else {
		msg := fmt.Sprintf("Unable to cancel invitation. User [%v] with role [%v] cannot cancel invitations.", tokenClaims.ID, tokenClaims.RolePermissions.RoleName)
		result.Status = operationresult.Forbidden
		result.Message = msg
		applogsdomain.LogError(domainName, functionName, msg, false)
	}

	return result
}

//PRIVATE FUNCTIONS
/************************************************************************************/
func validateAccountInvite(accountInvitation accountinvitationmodel.AccountInvitation) (bool, string) {
	var buffer bytes.Buffer
	isValid := true

	if len(accountInvitation.RoleId) <= 0 {
		isValid = false
		buffer.WriteString(" [Empty role Id] ")
	}

	if len(accountInvitation.Email) <= 0 {
		isValid = false
		buffer.WriteString(" [Empty email] ")
	}

	emailPattern := regexp.MustCompile(`^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`)
	validEmail := emailPattern.MatchString(accountInvitation.Email)
	if !validEmail {
		isValid = false
		buffer.WriteString(" [Invalid email] ")
	}

	if accountInvitation.RoleId.String() != roleconstants.ADMIN && accountInvitation.RoleId.String() != roleconstants.PRESENTER {
		isValid = false
		buffer.WriteString(" [Invalid role] ")
	}

	return isValid, buffer.String()
}
