package userdomain

import (
	"fmt"
	"justasking/GO/common/operationresult"
	"justasking/GO/common/utility"
	"justasking/GO/core/domain/appconfigs"
	"justasking/GO/core/domain/applogs"
	"justasking/GO/core/domain/email"
	"justasking/GO/core/domain/priceplan"
	"justasking/GO/core/domain/role"
	"justasking/GO/core/model/passwordresetrequest"
	"justasking/GO/core/model/user"
	"justasking/GO/core/repo/emailtemplate"
	"justasking/GO/core/repo/user"
	"strconv"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

var domainName = "UserDomain"

// GetUser returns a User for the id provided
func GetUser(userId uuid.UUID) (usermodel.User, *operationresult.OperationResult) {
	functionName := "GetUser"
	result := operationresult.New()

	user, err := userrepo.GetUserById(userId)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting user with id [%v]. Error: [%v]", userId, msg), false)
	} else {
		pricePlan, pricePlanResult := priceplandomain.GetPricePlanDetailsByAccountId(user.Account.Id)
		if pricePlanResult.IsSuccess() {
			user.MembershipDetails = pricePlan

			roles, rolesResult := roledomain.GetRolePermissionsByUserId(user.ID)
			if rolesResult.IsSuccess() {
				user.RolePermissions = roles
			} else {
				msg := fmt.Sprintf("Error retrieving user, unable to get permissions user [%v]. Error: [%v]", user.ID, pricePlanResult.Message)
				result.Message = pricePlanResult.Message
				result.Status = pricePlanResult.Status
				applogsdomain.LogError(domainName, functionName, msg, false)
			}
		} else {
			msg := fmt.Sprintf("Error retrieving user, unable to get price plan details for user [%v]. Error: [%v]", user.ID, pricePlanResult.Message)
			result.Message = pricePlanResult.Message
			result.Status = pricePlanResult.Status
			applogsdomain.LogError(domainName, functionName, msg, false)
		}
	}

	return user, result
}

// GetSimpleUserRecord gets the actual user record, without getting plan, role, or account data
func GetSimpleUserRecord(userId uuid.UUID) (usermodel.User, *operationresult.OperationResult) {
	functionName := "GetSimpleUserRecord"
	result := operationresult.New()

	user, err := userrepo.GetSimpleUserRecord(userId)
	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting user with id [%v]. Error: [%v]", userId, msg), false)
	}

	return user, result
}

// creates a reset password request
func RequestPasswordReset(email string) *operationresult.OperationResult {
	functionName := "RequestPasswordReset"
	result := operationresult.New()

	user, err := userrepo.GetUserByEmail(email)
	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Could not create password reset request. Unable to find user by email. Error: [%v]", msg), false)
	} else {
		var resetRequest passwordresetrequestmodel.PasswordResetRequest
		resetRequest.Id, _ = uuid.NewV4()
		resetRequest.UserId = user.ID
		resetRequest.ResetCode = utility.RandSeq(256)
		resetRequest.IsActive = true
		resetRequest.ExpiresAt = time.Now().AddDate(0, 0, 1)
		err := userrepo.InvalidateOldPasswordResetRequests(resetRequest)
		if err != nil {
			msg := err.Error()
			result = operationresult.CreateErrorResult(msg, err)
			applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Could not invalidate previous password reset requests for user [%v]. Error: [%v]", resetRequest.UserId, msg), false)
		} else {
			err := userrepo.CreatePasswordResetRequest(resetRequest)
			if err != nil {
				msg := err.Error()
				result = operationresult.CreateErrorResult(msg, err)
				applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Could not create password reset request. Error: [%v]", msg), false)
			} else {
				applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("Password reset request created for user [%v]", email))

				//send password reset email
				passwordResetTemplate, err := emailtemplaterepo.GetEmailTemplateByName("password_reset_request")
				if err != nil {
					applogsdomain.LogError(domainName, functionName, "Unable to retrieve password_reset_request email template.", false)
				} else {
					passwordResetTemplate.To = user.Email
					passwordResetTemplate.Body = strings.Replace(passwordResetTemplate.Body, "{resetCode}", resetRequest.ResetCode, -1)

					emailSendResult := emaildomain.SendEmail(passwordResetTemplate)
					if emailSendResult.IsSuccess() {
						applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("password_reset_request email sent to user [%v].", email))
					} else {
						applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Unable to send password_reset_request email to user [%v]. Error: [%v]", email, emailSendResult.Message), true)
					}
				}
			}
		}
	}

	return result
}

// UpdatePassword updates a password for a reset request
func UpdatePassword(resetRequest passwordresetrequestmodel.PasswordResetRequest) (usermodel.User, *operationresult.OperationResult) {
	functionName := "UpdatePasswordResetRequest"
	result := operationresult.New()
	var err error
	var user usermodel.User

	resetRequest, err = userrepo.GetResetRequest(resetRequest)
	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Could not update password for resetid [%v], unable to retrieve request from database. Error: [%v]", resetRequest.ResetCode, msg), false)
	} else {
		user, err = userrepo.GetUserByResetCode(resetRequest)

		if err != nil {
			msg := err.Error()
			result = operationresult.CreateErrorResult(msg, err)
			applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Unable to retrieve user for reset code [%v]. Error: [%v]", resetRequest.ResetCode, msg), false)
		} else {
			if time.Now().After(resetRequest.ExpiresAt) || resetRequest.IsActive == false {
				msg := fmt.Sprintf("Could not update password for resetid [%v] and userid [%v]. Link is no longer valid.", resetRequest.ResetCode, resetRequest.UserId)
				result.Status = operationresult.Gone
				result.Message = msg
				applogsdomain.LogError(domainName, functionName, msg, false)
			} else {
				justAskingConfigs, configsResult := appconfigsdomain.GetAppConfigs("justasking")
				if configsResult.IsSuccess() {
					minimumPasswordLength, _ := strconv.Atoi(justAskingConfigs["MinimumPasswordLength"])
					maximumPasswordLength, _ := strconv.Atoi(justAskingConfigs["MaximumPasswordLength"])
					if len(resetRequest.Password) >= minimumPasswordLength && len(resetRequest.Password) <= maximumPasswordLength {
						passwordHash, err := bcrypt.GenerateFromPassword([]byte(resetRequest.Password), bcrypt.DefaultCost)
						if err != nil {
							msg := err.Error()
							result = operationresult.CreateErrorResult(msg, err)
							applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error hashing password [%v]. Error: [%v]", resetRequest.Password, msg), false)
						} else {

							//update password
							err = userrepo.UpdatePassword(resetRequest, string(passwordHash))
							if err != nil {
								msg := err.Error()
								result = operationresult.CreateErrorResult(msg, err)
								applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error updating password for user [%v]. Error: [%v]", resetRequest.UserId, msg), false)
							} else {

								//send password reset confirmation
								passwordResetConfirmTemplate, err := emailtemplaterepo.GetEmailTemplateByName("password_reset_confirmation")
								if err != nil {
									applogsdomain.LogError(domainName, functionName, "Unable to retrieve password_reset_confirmation email template.", false)
								} else {
									passwordResetConfirmTemplate.To = user.Email

									emailSendResult := emaildomain.SendEmail(passwordResetConfirmTemplate)
									if emailSendResult.IsSuccess() {
										applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("password_reset_confirmation email sent to user [%v].", user.Email))
									} else {
										applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Unable to send password_reset_confirmation email to user [%v]. Error: [%v]", user.Email, emailSendResult.Message), true)
									}
								}
							}
						}
					} else {
						msg := "Password does not meet length requirements."
						result.Message = msg
						result.Status = operationresult.UnprocessableEntity
						applogsdomain.LogError(domainName, functionName, msg, false)
					}
				} else {
					msg := configsResult.Error.Error()
					result = operationresult.CreateErrorResult(msg, configsResult.Error)
					applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting app configs. Error: [%v]", msg), false)
				}
			}
		}
	}

	return user, result
}

// GetUserByEmail gets a user by email address
func GetUserByEmail(email string) (usermodel.User, *operationresult.OperationResult) {
	functionName := "GetUserByEmail"
	result := operationresult.New()

	user, err := userrepo.GetUserByEmail(email)
	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting user with email [%v]. Error: [%v]", email, msg), false)
	}

	return user, result
}
