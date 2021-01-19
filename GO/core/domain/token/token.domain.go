package tokendomain

import (
	"fmt"
	"io/ioutil"
	"justasking/GO/common/authenticationclaim"
	"justasking/GO/common/operationresult"
	"justasking/GO/core/domain/accountuser"
	"justasking/GO/core/domain/applogs"
	"justasking/GO/core/domain/priceplan"
	"justasking/GO/core/domain/role"
	"justasking/GO/core/domain/user"
	"justasking/GO/core/model/user"

	jwt "github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

var domainName = "TokenDomain"
var hmacSecret = GetSecret()

// GetNewToken generates a new token for a user who is already logged in. used when changing subscriptions.
func GetNewToken(userId uuid.UUID) (string, *operationresult.OperationResult) {
	functionName := "GetNewToken"
	result := operationresult.New()
	var tokenString string
	var err error

	user, userResult := userdomain.GetUser(userId)
	if userResult.IsSuccess() {
		pricePlan, pricePlanResult := priceplandomain.GetPricePlanDetailsByAccountId(user.Account.Id)
		if pricePlanResult.IsSuccess() {
			user.MembershipDetails = pricePlan

			rolePermissions, rolePermissionsResult := roledomain.GetRolePermissionsByUserId(user.ID)
			if rolePermissionsResult.IsSuccess() {
				user.RolePermissions = rolePermissions

				accountUser, accountUserResult := accountuserdomain.GetAccountUser(user.ID, user.Account.Id)
				if accountUserResult.IsSuccess() {
					tokenString, err = SerializeToken(user, accountUser.TokenVersion)
					if err != nil {
						msg := fmt.Sprintf("Error serializing token for user: [%v]. Error: [%v]", user.ID, err.Error())
						result = operationresult.CreateErrorResult(msg, err)
						applogsdomain.LogError(domainName, functionName, msg, false)
					}
				} else {
					msg := fmt.Sprintf("Error account details for user [%v] on account [%v].", user.ID, user.Account.Id)
					result.Message = msg
					result.Status = pricePlanResult.Status
					applogsdomain.LogError(domainName, functionName, msg, false)
				}
			} else {
				msg := fmt.Sprintf("Error retrieving permissions for user [%v].", user.ID)
				result.Message = msg
				result.Status = rolePermissionsResult.Status
				applogsdomain.LogError(domainName, functionName, msg, false)
			}
		} else {
			msg := fmt.Sprintf("Error retrieving price plan details for account [%v].", user.Account.Id)
			result.Message = msg
			result.Status = pricePlanResult.Error.Error()
			applogsdomain.LogError(domainName, functionName, msg, false)
		}
	} else {
		msg := fmt.Sprintf("Error retrieving user details for user [%v].", userId)
		result.Message = msg
		result.Status = operationresult.Error
		applogsdomain.LogError(domainName, functionName, msg, false)
	}

	return tokenString, result
}

// SerializeToken takes a User and returns a token string with a claim on User.ID.
func SerializeToken(user usermodel.User, tokenVersion uuid.UUID) (string, error) {
	funcName := "SerializeToken"

	// Create the Claims
	authenticationclaim := authenticationclaim.New()
	authenticationclaim.ID = user.ID
	authenticationclaim.FirstName = user.FirstName
	authenticationclaim.LastName = user.LastName
	authenticationclaim.Email = user.Email
	authenticationclaim.TokenVersion = tokenVersion
	authenticationclaim.MembershipDetails = user.MembershipDetails
	authenticationclaim.Account = user.Account
	authenticationclaim.RolePermissions = user.RolePermissions

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, authenticationclaim)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSecret)
	if err != nil {
		msg := err.Error()
		applogsdomain.LogError(domainName, funcName, msg, false)
	}

	return tokenString, err
}

// DeserializeToken takes a token string and parses it into a token object with claims.
func DeserializeToken(tokenString string) *authenticationclaim.AuthenticationClaim {
	token, err := jwt.ParseWithClaims(tokenString, &authenticationclaim.AuthenticationClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return GetSecret(), nil
	})

	if err != nil {
		fmt.Println(err.Error())
	}
	parsedToken := authenticationclaim.NewFromToken(token)

	return parsedToken
}

func GetSecret() []byte {
	funcName := "GetSecret"
	var secret []byte
	if keyData, e := ioutil.ReadFile("keys/hmac"); e == nil {
		secret = keyData
	} else {
		applogsdomain.LogError(domainName, funcName, e.Error(), false)
	}
	return secret
}
