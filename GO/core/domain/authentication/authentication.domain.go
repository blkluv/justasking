package authenticationdomain

import (
	"fmt"
	"justasking/GO/common/authcontainer"
	"justasking/GO/common/clients/recaptcha"
	"justasking/GO/common/constants/role"
	"justasking/GO/common/operationresult"
	"justasking/GO/core/domain/accountuser"
	"justasking/GO/core/domain/appconfigs"
	"justasking/GO/core/domain/applogs"
	"justasking/GO/core/domain/email"
	"justasking/GO/core/domain/priceplan"
	"justasking/GO/core/domain/role"
	"justasking/GO/core/domain/token"
	"justasking/GO/core/model/account"
	"justasking/GO/core/model/accountuser"
	"justasking/GO/core/model/authentication"
	"justasking/GO/core/model/idpjustasking"
	"justasking/GO/core/model/idpmapping"
	"justasking/GO/core/model/user"
	"justasking/GO/core/repo/authentication"
	"justasking/GO/core/repo/emailtemplate"
	"justasking/GO/core/repo/user"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/api/oauth2/v2"
)

var domainName = "AuthenticationDomain"

// GetToken returns a JWT based on a one time code and an Identity Provider
func GetToken(container authcontainer.AuthContainer) (string, bool, *operationresult.OperationResult) {
	funcName := "GetToken"
	var tokenString string
	var userCreated bool

	result := operationresult.New()

	// check data depending on IDP used
	if strings.ToLower(container.IdpName) == authcontainer.GOOGLE {
		tokenString, userCreated, result = loginWithGoogle(container)
	} else if strings.ToLower(container.IdpName) == authcontainer.JUSTASKING {
		userCreated = false
		tokenString, result = loginWithJustAsking(container)
	} else {
		//user is using a different IDP
		applogsdomain.LogInfo(domainName, funcName, fmt.Sprintf("Unknown IDP supplied: [%v]", container.IdpName))
	}

	fmt.Println(tokenString)
	return tokenString, userCreated, result
}

// Calls Google token endpoint to validate id_token sent by client
func getGoogleTokenInfo(container authcontainer.AuthContainer) (*oauth2.Tokeninfo, error) {
	funcName := "getGoogleTokenInfo"
	var httpClient = &http.Client{}
	oauth2Service, _ := oauth2.New(httpClient)
	tokenInfoCall := oauth2Service.Tokeninfo()
	tokenInfoCall.IdToken(container.IdpData["id_token"])
	tokenInfo, err := tokenInfoCall.Do()
	if err != nil {
		applogsdomain.LogError(domainName, funcName, err.Error(), false)
		return nil, err
	}
	return tokenInfo, nil
}

func getUserByGoogleSub(googleUserID string) (usermodel.User, error) {
	result, err := userrepo.GetJustAskingUserByGoogleSub(googleUserID)

	return result, err
}

func createGoogleUser(container authcontainer.AuthContainer, googleSub string) (usermodel.User, error) {
	var err error
	var user *usermodel.User
	functionName := "createGoogleUser"
	userIdguid, _ := uuid.NewV4()

	stripeKey, configsResult := appconfigsdomain.GetAppConfig("stripe", "StripeSecretKey")
	if configsResult.IsSuccess() {
		name := container.IdpData["name"]
		email := container.IdpData["email"]
		imageUrl := container.IdpData["imageUrl"]
		firstName := container.IdpData["givenName"]
		lastName := container.IdpData["familyName"]

		var googleUser authenticationmodel.IdpGoogle
		googleUser.Sub = googleSub
		googleUser.Name = &name
		googleUser.Email = &email
		googleUser.ImageUrl = &imageUrl
		googleUser.GivenName = &firstName
		googleUser.FamilyName = &lastName

		var justAskingUser usermodel.User
		justAskingUser.FirstName = firstName
		justAskingUser.LastName = lastName
		justAskingUser.Email = email
		justAskingUser.ImageUrl = imageUrl
		justAskingUser.ID = userIdguid
		justAskingUser.IsActive = true

		var idpMapping idpmappingmodel.IdpMapping
		idpMapping.IdpId = authcontainer.GOOGLEID
		idpMapping.Sub = googleSub

		var account accountmodel.Account
		account.Id, _ = uuid.NewV4()
		account.OwnerId = justAskingUser.ID
		account.Name = fmt.Sprintf("%v %v", firstName, lastName)
		account.CreatedBy = justAskingUser.ID.String()
		account.IsActive = true

		var accountUser accountusermodel.AccountUser
		accountUser.AccountId = account.Id
		accountUser.UserId = justAskingUser.ID
		accountUser.IsActive = true
		accountUser.RoleId, _ = uuid.FromString(roleconstants.OWNER)
		accountUser.CreatedBy = justAskingUser.ID.String()
		accountUser.CurrentAccount = true
		accountUser.TokenVersion, _ = uuid.NewV4()

		user, err = authenticationrepo.CreateGoogleUser(googleUser, justAskingUser, idpMapping, account, accountUser, stripeKey.ConfigValue)

		if err != nil {
			msg := fmt.Sprintf("Unable to create user [%v %v]. Error: [%v]", justAskingUser.FirstName, justAskingUser.LastName, err.Error())
			applogsdomain.LogError(domainName, functionName, msg, false)

			//assigning value to user to avoid panic
			user = new(usermodel.User)
		} else {
			applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("Successfully created user [%v]", justAskingUser.ID))

			user.Account = account

			//welcome user email
			welcomeEmailTemplate, err := emailtemplaterepo.GetEmailTemplateByName("service_welcome")
			if err != nil {
				applogsdomain.LogError(domainName, functionName, "Unable to retrieve service_welcome email template.", false)
			} else {
				welcomeEmailTemplate.To = user.Email
				emailSendResult := emaildomain.SendEmail(welcomeEmailTemplate)
				if emailSendResult.IsSuccess() {
					applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("service_welcome email sent to user [%v].", justAskingUser.ID))
				} else {
					applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Unable to send service_welcome email to user [%v]. Error: [%v]", justAskingUser.ID, emailSendResult.Message), false)
				}
			}
		}
	} else {
		err = configsResult.Error
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting app configs. Error: [%v]", configsResult.Message), false)
	}

	return *user, err
}

func loginWithGoogle(container authcontainer.AuthContainer) (string, bool, *operationresult.OperationResult) {
	functionName := "loginWithGoogle"
	result := operationresult.New()
	var tokenString string
	userWasCreated := false

	//validate id_token by sending to google
	data, err := getGoogleTokenInfo(container)
	if err == nil {
		tokenInfo := data
		//try to get the user record, using the google sub
		user, err := getUserByGoogleSub(tokenInfo.UserId)
		if err == gorm.ErrRecordNotFound {
			_, err := userrepo.GetUserByEmail(container.IdpData["email"])
			if err != nil && err != gorm.ErrRecordNotFound {
				//there was a real error
				msg := fmt.Sprintf("Error getting user by email while creating google user with email [%v]. Error: [%v]", container.IdpData["email"], err.Error())
				result = operationresult.CreateErrorResult(msg, err)
				applogsdomain.LogError(domainName, functionName, msg, false)
			} else if err != nil && err == gorm.ErrRecordNotFound {
				//user doesn't exist, we need to create it
				user, err := createGoogleUser(container, tokenInfo.UserId)
				if err == nil {
					// after user is created, get price plan details, return user and serialize a token with it
					pricePlan, result := priceplandomain.GetPricePlanDetailsByAccountId(user.Account.Id)
					if result.IsSuccess() {
						user.MembershipDetails = pricePlan

						rolePermissions, rolePermissionsResult := roledomain.GetRolePermissionsByUserId(user.ID)
						if rolePermissionsResult.IsSuccess() {
							user.RolePermissions = rolePermissions

							accountUser, accountUserResult := accountuserdomain.GetAccountUser(user.ID, user.Account.Id)
							if accountUserResult.IsSuccess() {

								if accountUser.IsActive {
									tokenString, err = tokendomain.SerializeToken(user, accountUser.TokenVersion)

									userWasCreated = true
									if err != nil {
										msg := fmt.Sprintf("Error serializing token for user: [%v]. Error: [%v]", user.ID, err.Error())
										result = operationresult.CreateErrorResult(msg, err)
										applogsdomain.LogError(domainName, functionName, msg, false)
									}
								} else {
									msg := fmt.Sprintf("Error retrieving accountUser details for user [%v] with accountId [%v]. User is not active on the account.", user.ID, user.Account.Id)
									result.Status = operationresult.Error
									result.Message = msg
									applogsdomain.LogError(domainName, functionName, msg, false)
								}
							} else {
								msg := fmt.Sprintf("Error retrieving accountUser details for user [%v] with accountId [%v]. Error: [%v]", user.ID, user.Account.Id, accountUserResult.Message)
								result.Status = operationresult.Error
								result.Message = msg
								applogsdomain.LogError(domainName, functionName, msg, false)
							}
						} else {
							msg := fmt.Sprintf("Error getting rolepermissions for user: [%v]. Error: [%v]", user.ID, rolePermissionsResult.Message)
							result.Status = operationresult.Error
							result.Message = msg
							applogsdomain.LogError(domainName, functionName, msg, false)
						}
					} else {
						msg := fmt.Sprintf("Error getting priceplan for user: [%v]. Error: [%v]", user.ID, result.Message)
						result.Status = operationresult.Error
						result.Message = msg
						applogsdomain.LogError(domainName, functionName, msg, false)
					}

				} else {
					msg := fmt.Sprintf("Error creating Google user. Error: [%v]", err.Error())
					result = operationresult.CreateErrorResult(msg, err)
					applogsdomain.LogError(domainName, functionName, msg, false)
				}
			} else {
				//the user exists, return conflict status
				msg := fmt.Sprintf("Justasking user already exists. Cannot create user with email [%v].", container.IdpData["email"])
				result.Message = msg
				result.Status = operationresult.Conflict
				applogsdomain.LogError(domainName, functionName, msg, false)
			}
		} else if err == nil {
			user, err = userrepo.GetUserById(user.ID)
			if err != nil {
				msg := fmt.Sprintf("Error getting existing user [%v]. Error: [%v]", user.ID, err.Error())
				result = operationresult.CreateErrorResult(msg, err)
				applogsdomain.LogError(domainName, functionName, msg, false)
			} else {
				//user exists, get price plan details, return user and serialize a token with it
				pricePlan, result := priceplandomain.GetPricePlanDetailsByAccountId(user.Account.Id)
				if result.IsSuccess() {
					user.MembershipDetails = pricePlan

					rolePermissions, rolePermissionsResult := roledomain.GetRolePermissionsByUserId(user.ID)
					if rolePermissionsResult.IsSuccess() {
						user.RolePermissions = rolePermissions

						accountUser, accountUserResult := accountuserdomain.GetAccountUser(user.ID, user.Account.Id)
						if accountUserResult.IsSuccess() {

							tokenString, err = tokendomain.SerializeToken(user, accountUser.TokenVersion)
							if err != nil {
								msg := fmt.Sprintf("Error serializing token for user: [%v]. Error: [%v]", user.ID, err.Error())
								result = operationresult.CreateErrorResult(msg, err)
								applogsdomain.LogError(domainName, functionName, msg, false)
							}

							userrepo.UpdateUserLastLogin(user.ID)

						} else {
							msg := fmt.Sprintf("Error getting accountUser data for user [%v] with accountId [%v]. Message: [%v]", user.ID, user.Account.Id, accountUserResult.Message)
							result.Status = operationresult.Error
							result.Message = msg
							applogsdomain.LogError(domainName, functionName, msg, false)
						}
					} else {
						msg := fmt.Sprintf("Error permissions for user: [%v].", user.ID)
						result.Status = operationresult.Error
						result.Message = msg
						applogsdomain.LogError(domainName, functionName, msg, false)
					}
				} else {
					msg := fmt.Sprintf("Error getting priceplan for user: [%v].", user.ID)
					result.Status = operationresult.Error
					result.Message = msg
					applogsdomain.LogError(domainName, functionName, msg, false)
				}
			}
		} else {
			msg := fmt.Sprintf("Error getting google user from database. Error: [%v]", err.Error())
			result = operationresult.CreateErrorResult(msg, err)
			applogsdomain.LogError(domainName, functionName, msg, false)
		}
	} else {
		msg := fmt.Sprintf("Error vaidating google token against google api. Error: [%v]", err.Error())
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, msg, false)
	}

	return tokenString, userWasCreated, result
}

// CreateJustAskingUser creates a record in the idp_justasking table
func CreateJustAskingUser(idpJustAskingUser idpjustaskingmodel.IdpJustAsking) (string, *operationresult.OperationResult) {
	functionName := "createGoogleUser"
	var user *usermodel.User
	var tokenString string
	result := operationresult.New()

	captchaIsValid := recaptchaclient.ValidateReCaptchaToken(idpJustAskingUser.CaptchaToken)
	if captchaIsValid {
		_, err := userrepo.GetUserByEmail(idpJustAskingUser.Email)
		if err != nil && err != gorm.ErrRecordNotFound {
			//there was a real error
			msg := fmt.Sprintf("Error creating idpJustAsking user. Error: [%v]", err.Error())
			result = operationresult.CreateErrorResult(msg, err)
			applogsdomain.LogError(domainName, functionName, msg, false)
		} else if err != nil && err == gorm.ErrRecordNotFound {
			// the user does not exist
			justAskingConfigs, configsResult := appconfigsdomain.GetAppConfigs("justasking")
			if configsResult.IsSuccess() {
				isValid, validationMessage := validateJustAskingUser(idpJustAskingUser, justAskingConfigs)
				if isValid {
					passwordHash, err := bcrypt.GenerateFromPassword([]byte(idpJustAskingUser.Password), bcrypt.DefaultCost)
					if err != nil {
						msg := configsResult.Error.Error()
						result = operationresult.CreateErrorResult(msg, configsResult.Error)
						applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error hashing password [%v]. Error: [%v]", idpJustAskingUser.Password, msg), false)
					} else {
						userId, _ := uuid.NewV4()
						idpJustAskingUser.Sub = userId
						idpJustAskingUser.Password = string(passwordHash)

						stripeKey, configsResult := appconfigsdomain.GetAppConfig("stripe", "StripeSecretKey")
						if configsResult.IsSuccess() {

							var justAskingUser usermodel.User
							justAskingUser.FirstName = idpJustAskingUser.GivenName
							justAskingUser.LastName = idpJustAskingUser.FamilyName
							justAskingUser.Email = idpJustAskingUser.Email
							justAskingUser.ImageUrl = idpJustAskingUser.ImageUrl
							justAskingUser.ID = userId
							justAskingUser.IsActive = true

							var idpMapping idpmappingmodel.IdpMapping
							idpMapping.IdpId = authcontainer.JUSTASKINGID
							idpMapping.Sub = userId.String()

							var account accountmodel.Account
							account.Id, _ = uuid.NewV4()
							account.OwnerId = justAskingUser.ID
							account.Name = fmt.Sprintf("%v %v", idpJustAskingUser.GivenName, idpJustAskingUser.FamilyName)
							account.CreatedBy = justAskingUser.ID.String()
							account.IsActive = true

							var accountUser accountusermodel.AccountUser
							accountUser.AccountId = account.Id
							accountUser.UserId = justAskingUser.ID
							accountUser.IsActive = true
							accountUser.RoleId, _ = uuid.FromString(roleconstants.OWNER)
							accountUser.CreatedBy = justAskingUser.ID.String()
							accountUser.CurrentAccount = true
							accountUser.TokenVersion, _ = uuid.NewV4()

							user, err = authenticationrepo.CreateIdpJustAskingUser(idpJustAskingUser, justAskingUser, idpMapping, account, accountUser, stripeKey.ConfigValue)
							if err != nil {
								msg := fmt.Sprintf("Error creating idpJustAsking user. Error: [%v]", err.Error())
								result = operationresult.CreateErrorResult(msg, err)
								applogsdomain.LogError(domainName, functionName, msg, false)
							} else {
								user.Account = account

								applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("Successfully created user [%v]", justAskingUser.ID))

								//welcome user email
								welcomeEmailTemplate, err := emailtemplaterepo.GetEmailTemplateByName("service_welcome")
								if err != nil {
									applogsdomain.LogError(domainName, functionName, "Unable to retrieve service_welcome email template.", false)
								} else {
									welcomeEmailTemplate.To = user.Email
									emailSendResult := emaildomain.SendEmail(welcomeEmailTemplate)
									if emailSendResult.IsSuccess() {
										applogsdomain.LogInfo(domainName, functionName, fmt.Sprintf("service_welcome email sent to user [%v].", justAskingUser.ID))
									} else {
										applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Unable to send service_welcome email to user [%v]. Error: [%v]", justAskingUser.ID, emailSendResult.Message), false)
									}
								}

								//create token
								pricePlan, result := priceplandomain.GetPricePlanDetailsByAccountId(user.Account.Id)
								if result.IsSuccess() {
									user.MembershipDetails = pricePlan

									rolePermissions, rolePermissionsResult := roledomain.GetRolePermissionsByUserId(user.ID)
									if rolePermissionsResult.IsSuccess() {
										user.RolePermissions = rolePermissions

										tokenString, err = tokendomain.SerializeToken(*user, accountUser.TokenVersion)
										if err != nil {
											msg := fmt.Sprintf("Error serializing token for user: [%v]. Error: [%v]", user.ID, err.Error())
											result = operationresult.CreateErrorResult(msg, err)
											applogsdomain.LogError(domainName, functionName, msg, false)
										}
									} else {
										msg := fmt.Sprintf("Error getting permissions for user: [%v]. Error: [%v]", user.ID, rolePermissionsResult.Message)
										result.Status = operationresult.Error
										result.Message = msg
										applogsdomain.LogError(domainName, functionName, msg, false)
									}
								} else {
									msg := fmt.Sprintf("Error getting priceplan for user: [%v]. Error: [%v]", user.ID, result.Message)
									result.Status = operationresult.Error
									result.Message = msg
									applogsdomain.LogError(domainName, functionName, msg, false)
								}
							}
						} else {
							err = configsResult.Error
							applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting app configs. Error: [%v]", configsResult.Message), false)
						}
					}
				} else {
					result.Message = validationMessage
					result.Status = operationresult.UnprocessableEntity
					applogsdomain.LogError(domainName, functionName, validationMessage, false)
				}
			} else {
				msg := configsResult.Error.Error()
				result = operationresult.CreateErrorResult(msg, configsResult.Error)
				applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting app configs. Error: [%v]", msg), false)
			}
		} else {
			//the user exists, return conflict status
			msg := fmt.Sprintf("Justasking user already exists. Cannot create user with email [%v].", idpJustAskingUser.Email)
			result.Message = msg
			result.Status = operationresult.Conflict
			applogsdomain.LogError(domainName, functionName, msg, false)
		}
	} else {
		msg := fmt.Sprintf("Email [%v] did not validate successfully with recaptcha.", idpJustAskingUser.Email)
		result.Status = operationresult.Forbidden
		result.Message = msg
		applogsdomain.LogError(domainName, functionName, msg, false)
	}

	return tokenString, result
}

func loginWithJustAsking(container authcontainer.AuthContainer) (string, *operationresult.OperationResult) {
	functionName := "loginWithJustAsking"
	result := operationresult.New()
	var tokenString string

	email := container.IdpData["email"]
	password := container.IdpData["password"]

	idpJustAskingUser, err := userrepo.GetJustAskingIdpUser(email)
	if err != nil {
		msg := "Incorrect username or password."
		result.Message = msg
		result.Status = operationresult.Unauthorized
		applogsdomain.LogError(domainName, functionName, msg, false)
	} else {
		err := bcrypt.CompareHashAndPassword([]byte(idpJustAskingUser.Password), []byte(password))
		if err != nil {
			msg := "Incorrect username or password."
			result.Message = msg
			result.Status = operationresult.Unauthorized
			applogsdomain.LogError(domainName, functionName, msg, false)
		} else {
			token, tokenResult := tokendomain.GetNewToken(idpJustAskingUser.Sub)
			if tokenResult.IsSuccess() {
				userrepo.UpdateUserLastLogin(idpJustAskingUser.Sub)
				tokenString = token
			} else {
				msg := tokenResult.Error.Error()
				result = operationresult.CreateErrorResult(msg, tokenResult.Error)
				applogsdomain.LogError(domainName, functionName, msg, false)
			}
		}
	}

	return tokenString, result
}

func validateJustAskingUser(idpJustAskingUser idpjustaskingmodel.IdpJustAsking, justAskingConfigs map[string]string) (bool, string) {
	var validationMessage string
	isValid := true

	minimumPasswordLength, _ := strconv.Atoi(justAskingConfigs["MinimumPasswordLength"])
	maximumPasswordLength, _ := strconv.Atoi(justAskingConfigs["MaximumPasswordLength"])
	if len(idpJustAskingUser.Password) < minimumPasswordLength || len(idpJustAskingUser.Password) > maximumPasswordLength {
		isValid = false
		validationMessage = "Invalid password length."
	}

	if len(idpJustAskingUser.GivenName) <= 0 {
		isValid = false
		validationMessage = " Empty firstname."
	}

	if len(idpJustAskingUser.FamilyName) <= 0 {
		isValid = false
		validationMessage = "Empty lastname."
	}

	emailPattern := regexp.MustCompile(`^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`)
	validEmail := emailPattern.MatchString(idpJustAskingUser.Email)
	if !validEmail {
		isValid = false
		validationMessage = "Invalid email."
	}

	return isValid, validationMessage
}
