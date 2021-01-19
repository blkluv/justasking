package baseboxdomain

import (
	"fmt"
	authenticationclaim "justasking/GO/common/authenticationclaim"
	"justasking/GO/common/operationresult"
	"justasking/GO/core/domain/account"
	"justasking/GO/core/domain/applogs"
	"justasking/GO/core/model/boxes/basebox"
	"justasking/GO/core/model/phonenumber"
	"justasking/GO/core/model/priceplan"
	"justasking/GO/core/repo/boxes/basebox"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

var domainName = "BaseBoxDomain"

// GetBoxesByAccountId gets all boxes for a specific account
func GetBoxesByAccountId(accountId uuid.UUID, tokenClaims *authenticationclaim.AuthenticationClaim) ([]baseboxmodel.BaseBox, *operationresult.OperationResult) {
	functionName := "GetBoxesByAccountId"
	result := operationresult.New()
	var boxes []baseboxmodel.BaseBox
	var err error

	if tokenClaims.RolePermissions.SeeAllBoxes {
		boxes, err = baseboxrepo.GetBoxesByAccountId(accountId)
		if err != nil {
			msg := err.Error()
			result = operationresult.CreateErrorResult(msg, err)
			applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting boxes for accoundId  [%v]. Error: [%v]", accountId, msg), false)
		}
	} else {
		boxes, err = baseboxrepo.GetBoxesByUserId(tokenClaims.ID, accountId)
		if err != nil {
			msg := err.Error()
			result = operationresult.CreateErrorResult(msg, err)
			applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting boxes for accoundId  [%v]. Error: [%v]", accountId, msg), false)
		}
	}

	return boxes, result
}

// GetActiveBoxesByAccountId gets all active boxes for a specific account
func GetActiveBoxesByAccountId(accountId uuid.UUID) ([]baseboxmodel.BaseBox, *operationresult.OperationResult) {
	functionName := "GetActiveBoxesByAccountId"
	result := operationresult.New()

	boxes, err := baseboxrepo.GetActiveBoxesByAccountId(accountId)
	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting active boxes for accoundId  [%v]. Error: [%v]", accountId, msg), false)
	}

	return boxes, result
}

// GetBaseBoxByCode gets the base box by code
func GetBaseBoxByCode(code string) (baseboxmodel.BaseBox, *operationresult.OperationResult) {
	functionName := "GetBaseBoxByCode"
	result := operationresult.New()

	box, err := baseboxrepo.GetBaseBoxByCode(code)
	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting base box by code  [%v]. Error: [%v]", code, msg), false)
	}

	return box, result
}

// GetBaseBoxByCodeAuthorized gets the base box by code
func GetBaseBoxByCodeAuthorized(code string, tokenClaims *authenticationclaim.AuthenticationClaim) (baseboxmodel.BaseBox, *operationresult.OperationResult) {
	functionName := "GetBaseBoxByCodeAuthorized"
	result := operationresult.New()

	box, err := baseboxrepo.GetBaseBoxByCode(code)
	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting base box by code  [%v]. Error: [%v]", code, msg), false)
	} else {
		//check if box belongs to this user's accounts
		if box.AccountId != tokenClaims.Account.Id {
			msg := fmt.Sprintf("Unable to retrieve basebox for user [%v] on account [%v]. Box belongs to account [%v].", tokenClaims.ID, tokenClaims.Account.Id, box.AccountId)
			result.Status = operationresult.Forbidden
			result.Message = msg
			applogsdomain.LogError(domainName, functionName, msg, false)
		} else {

			// check if user can see all boxes
			if !tokenClaims.RolePermissions.SeeAllBoxes {

				//if the user can't see all boxes, need to check if the box belongs to this user
				userId := tokenClaims.ID.String()
				if box.CreatedBy != userId {
					//box doesn't belong to user, need to return forbidden
					msg := fmt.Sprintf("Unable to retrieve basebox for user [%v]. Insufficient permissions and box doesn't belong to this user.", tokenClaims.ID)
					result.Status = operationresult.Forbidden
					result.Message = msg
					applogsdomain.LogError(domainName, functionName, msg, false)
				}
			}
		}
	}

	return box, result
}

// GetBaseBoxByPhoneNumber gets the base box by code
func GetBaseBoxByPhoneNumber(phoneNumber string) (baseboxmodel.BaseBox, *operationresult.OperationResult) {
	functionName := "GetBaseBoxByPhoneNumber"
	result := operationresult.New()

	box, err := baseboxrepo.GetBaseBoxByPhoneNumber(phoneNumber)
	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting base box by phone number  [%v]. Error: [%v]", phoneNumber, msg), false)
	}

	return box, result
}

// BoxCodeExists checks wether a code exists or not
func BoxCodeExists(code string) (struct{ Exists bool }, *operationresult.OperationResult) {
	result := operationresult.New()
	var data struct{ Exists bool }

	codeExists, _ := baseboxrepo.BoxCodeExists(code)

	data.Exists = codeExists

	return data, result
}

// ActivateBaseBox sets is_live flag to TRUE
func ActivateBaseBox(boxId uuid.UUID, updatedBy uuid.UUID, membershipDetails priceplanmodel.PricePlan) (phonenumbermodel.PhoneNumber, *operationresult.OperationResult) {
	functionName := "ActivateBaseBox"
	result := operationresult.New()

	returnPhoneNumber := phonenumbermodel.PhoneNumber{}

	box, err := baseboxrepo.GetBaseBoxById(boxId)
	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting base box for activation. Error: [%v]", msg), false)
	} else {
		userBelongsToAccount := accountdomain.UserBelongsToAccount(updatedBy, box.AccountId)
		if !userBelongsToAccount {
			msg := "Unable to activate box for different user."
			result = operationresult.CreateErrorResult(msg, err)
			applogsdomain.LogError(domainName, functionName, fmt.Sprintf("User [%v] attempted to open box owned by user [%v].", updatedBy, box.AccountId), false)
		} else if box.IsLive {
			msg := "This box is already open."
			result.Message = msg
			result.Status = operationresult.Conflict
			applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("User [%v] tried to open box [%v], which was already open.", updatedBy, boxId))
		} else {
			// TODO: check that the user who made this call has access to this box or other people's boxes

			// check if account hasn't hit limit on active boxes
			boxCanBeActivated := BoxCanBeActivated(box.AccountId, membershipDetails.ActiveBoxesLimit)
			if boxCanBeActivated.IsSuccess() {
				if membershipDetails.Sms {
					phoneNumber, err := baseboxrepo.ActivateBaseBoxAndAssignNumber(boxId, updatedBy)
					returnPhoneNumber = phoneNumber
					if err != nil {
						msg := err.Error()
						result = operationresult.CreateErrorResult(msg, err)
						applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error activating and assigning number to basebox [%v]. Error: [%v]", boxId, msg), false)

						//if there was an error opening/assigning the number, try to open without a number
						err := baseboxrepo.ActivateBaseBox(boxId, updatedBy)
						if err != nil {
							msg := err.Error()
							result = operationresult.CreateErrorResult(msg, err)
							applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error activating basebox [%v] without phone number after error assigning number. Error: [%v]", boxId, msg), false)
						} else {
							result.Status = operationresult.Success
							applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("User [%v] activated basebox [%v] without phone number after error assigning number.", updatedBy, boxId))
						}
					} else {
						applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("User [%v] activated and assigned number to basebox [%v]", updatedBy, boxId))
					}
				} else {
					err := baseboxrepo.ActivateBaseBox(boxId, updatedBy)
					if err != nil {
						msg := err.Error()
						result = operationresult.CreateErrorResult(msg, err)
						applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error activating basebox [%v] without phone number. Error: [%v]", boxId, msg), false)
					} else {
						applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("User [%v] activated basebox [%v] without phone number", updatedBy, boxId))
					}
				}
			} else {
				result.Status = boxCanBeActivated.Status
				result.Message = boxCanBeActivated.Message
				applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error updating IsLive flag to TRUE for box [%v]. Error: [%v]", boxId, boxCanBeActivated.Message), false)
			}

		}
	}

	return returnPhoneNumber, result
}

// DeactivateBaseBox sets is_live flag to FALSE
func DeactivateBaseBox(guid uuid.UUID, updatedBy uuid.UUID) *operationresult.OperationResult {
	functionName := "DeactivateBaseBox"
	result := operationresult.New()

	box, err := baseboxrepo.GetBaseBoxById(guid)
	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting base box for deactivation. Error: [%v]", msg), false)
	} else {
		userBelongsToAccount := accountdomain.UserBelongsToAccount(updatedBy, box.AccountId)
		if !userBelongsToAccount {
			msg := "Unable to deactivate box for different account."
			result = operationresult.CreateErrorResult(msg, err)
			applogsdomain.LogError(domainName, functionName, fmt.Sprintf("User [%v] attempted to close box owned by account [%v].", updatedBy, box.AccountId), false)
		} else if !box.IsLive {
			msg := "This box is already closed."
			result.Message = msg
			result.Status = operationresult.Conflict
			applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("User [%v] tried to close box [%v], which was already closed.", updatedBy, guid))
		} else {
			// TODO: check that the user who made this call has access to this box or other people's boxes

			err := baseboxrepo.DeactivateBaseBox(guid, updatedBy)
			if err != nil {
				msg := err.Error()
				result = operationresult.CreateErrorResult(msg, err)
				applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error updating IsLive flag to FALSE for box [%v]. Error: [%v]", guid, msg), false)
			} else {
				applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("User [%v] updated IsLive flag to FALSE for box [%v]", updatedBy, guid))
			}
		}
	}

	return result
}

// DeactivateAllBaseBoxesByAccountId deactivates all boxes for an account
func DeactivateAllBaseBoxesByAccountId(accountId uuid.UUID) *operationresult.OperationResult {
	functionName := "DeactivateAllBaseBoxesByAccountId"
	result := operationresult.New()

	err := baseboxrepo.DeactivateAllBaseBoxesByAccountId(accountId)
	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error deactivating all baseboxes for account [%v] after plan expiration. Error: [%v]", accountId, msg), true)
	} else {
		applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("Deactivated all boxes for account [%v] after plan expiration", accountId))
	}

	return result
}

// BoxCanBeActivated checks whether an account has hit the open box limit
func BoxCanBeActivated(accountId uuid.UUID, boxesLimit int) *operationresult.OperationResult {
	functionName := "BoxCanBeActivated"
	result := operationresult.New()

	activeBoxes, activeBoxesResult := GetActiveBoxesByAccountId(accountId)
	if activeBoxesResult.IsSuccess() {
		if len(activeBoxes) >= boxesLimit {
			msg := fmt.Sprintf("Account [%v] attempted to activate box after reaching monthly limit of [%v].", accountId, boxesLimit)
			result.Status = operationresult.Forbidden
			result.Message = msg
			applogsdomain.LogError(domainName, functionName, msg, false)
		}
	} else {
		result.Error = activeBoxesResult.Error
		result.Status = activeBoxesResult.Status
		result.Message = activeBoxesResult.Message
	}

	return result
}

// UserCanUpdateBox returns true if the user is on the account associated with the box
func UserCanUpdateBox(userId uuid.UUID, boxId uuid.UUID) bool {

	userCanUpdateEntry := false
	account, accountResponse := accountdomain.GetAccountByBoxId(boxId)
	if accountResponse.IsSuccess() {
		userBelongsToAccount := accountdomain.UserBelongsToAccount(userId, account.Id)
		if userBelongsToAccount {
			userCanUpdateEntry = true
		}
	}

	return userCanUpdateEntry
}

// GetBaseBoxByAnswerBoxQuestionId gets a base box by an answer box question id
func GetBaseBoxByAnswerBoxQuestionId(questionId uuid.UUID) (baseboxmodel.BaseBox, *operationresult.OperationResult) {
	functionName := "GetBaseBoxByAnswerBoxQuestionId"
	result := operationresult.New()

	box, err := baseboxrepo.GetBaseBoxByAnswerBoxQuestionId(questionId)
	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting base box by question Id [%v]. Error: [%v]", questionId, msg), false)
	}

	return box, result
}

// DeleteBaseBox marks a basebox as deleted
func DeleteBaseBox(boxId uuid.UUID, tokenClaims *authenticationclaim.AuthenticationClaim) *operationresult.OperationResult {
	functionName := "DeleteBaseBox"
	result := operationresult.New()

	box, err := baseboxrepo.GetBaseBoxById(boxId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			msg := fmt.Sprintf("Unable to delete box [%v]. Message: [%v]", boxId, err.Error())
			result.Status = operationresult.NotFound
			result.Message = msg
			applogsdomain.LogError(domainName, functionName, msg, false)
		} else {
			msg := err.Error()
			result = operationresult.CreateErrorResult(msg, err)
			applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting base box for deletion. Error: [%v]", msg), false)
		}
	} else {
		//check if the user who made this call is on the account that the box belongs to
		if box.AccountId != tokenClaims.Account.Id {
			msg := "Unable to delete box for different account."
			result.Status = operationresult.Forbidden
			result.Message = msg
			applogsdomain.LogError(domainName, functionName, fmt.Sprintf("User [%v] attempted to delete box owned by account [%v].", tokenClaims.ID, box.AccountId), false)
		} else {

			// check that the user who made this call has access to this box or other people's boxes
			canUpdateBox := false
			if tokenClaims.RolePermissions.SeeAllBoxes {
				canUpdateBox = true
			} else {
				createdBy, _ := uuid.FromString(box.CreatedBy)
				if createdBy == tokenClaims.ID {
					canUpdateBox = true
				}
			}

			if canUpdateBox {
				err := baseboxrepo.DeleteBaseBox(boxId, tokenClaims.ID)
				if err != nil {
					msg := err.Error()
					result = operationresult.CreateErrorResult(msg, err)
					applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error deleting basebox with boxId [%v] for user [%v]. Error: [%v]", boxId, tokenClaims.ID, msg), false)
				} else {
					applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("Basebox [%v] deleted by user [%v]", boxId, tokenClaims.ID))
				}
			} else {
				msg := fmt.Sprintf("Unable to delete box [%v] by user [%v]. User does not have permission.", box.ID, tokenClaims.ID)
				result.Message = msg
				result.Status = operationresult.Forbidden
				applogsdomain.LogError(domainName, functionName, msg, false)
			}
		}
	}

	return result
}
