package authenticationclaim

import (
	"justasking/GO/core/model/role"
	"strconv"
	"time"

	"justasking/GO/core/model/account"
	"justasking/GO/core/model/priceplan"
	"justasking/GO/core/startup/flight"

	jwt "github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

// AuthenticationClaim represents a claim that justasking uses to authorize requests
type AuthenticationClaim struct {
	ID                uuid.UUID                `json:"id"`
	FirstName         string                   `json:"firstname"`
	LastName          string                   `json:"lastname"`
	Email             string                   `json:"email"`
	TokenVersion      uuid.UUID                `json:"tokenversion"`
	MembershipDetails priceplanmodel.PricePlan `json:"membershipdetails"`
	Account           accountmodel.Account     `json:"account"`
	StandardClaims    jwt.StandardClaims       `json:"standardclaims"`
	RolePermissions   rolemodel.Role           `json:"rolepermissions"`
}

// New returns an initialized AuthenticationClaim
func New() *AuthenticationClaim {
	authenticationClaim := new(AuthenticationClaim)
	config := flight.Context(nil, nil).Config
	exp, err := strconv.Atoi(config.Settings["TokenExpirationInDays"])
	if err != nil {
		//setting default
		exp = 30
	}
	authenticationClaim.StandardClaims = jwt.StandardClaims{
		Issuer:    "justasking",
		ExpiresAt: time.Now().AddDate(0, 0, exp).Unix(),
	}
	return authenticationClaim
}

// NewFromToken returns a parsed AuthenticationClaim object from jwt claims passed in
func NewFromToken(token *jwt.Token) *AuthenticationClaim {
	authenticationClaim := new(AuthenticationClaim)

	if claims, ok := token.Claims.(*AuthenticationClaim); ok && token.Valid {
		authenticationClaim.ID = claims.ID
		authenticationClaim.FirstName = claims.FirstName
		authenticationClaim.LastName = claims.LastName
		authenticationClaim.Email = claims.Email
		authenticationClaim.MembershipDetails = claims.MembershipDetails
		authenticationClaim.Account = claims.Account
		authenticationClaim.RolePermissions = claims.RolePermissions
		authenticationClaim.TokenVersion = claims.TokenVersion
	}

	return authenticationClaim
}

// Valid returns a bool with the validity of the claims
func (a *AuthenticationClaim) Valid() error {
	return a.StandardClaims.Valid()
}
