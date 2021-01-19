package middleware

import (
	"fmt"
	"justasking/GO/common/authenticationclaim"
	"justasking/GO/common/constants/priceplan"
	"justasking/GO/core/domain/accountuser"
	"justasking/GO/core/domain/token"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gorilla/context"
	uuid "github.com/satori/go.uuid"
)

// LogRequestHandler will log the HTTP requests.
func LogRequestHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(time.Now().Format("2006-01-02 03:04:05 PM"), r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

// AuthorizedHandler will autorize given request based on JWT that was passed in.
func AuthorizedHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//verify that the token provided in the call is legit
		token, err := request.ParseFromRequestWithClaims(r, request.AuthorizationHeaderExtractor, &authenticationclaim.AuthenticationClaim{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return tokendomain.GetSecret(), nil
		})

		if err != nil || !token.Valid {
			//deny request if it is not legit
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			//proceed with request if it is indeed legit
			parsedToken := authenticationclaim.NewFromToken(token)

			user, userResult := accountuserdomain.GetAccountUser(parsedToken.ID, parsedToken.Account.Id)
			if userResult.IsSuccess() && user.IsActive {

				if parsedToken.TokenVersion != user.TokenVersion {
					w.WriteHeader(http.StatusUnauthorized)
				} else {
					basicPlanID, _ := uuid.FromString(priceplanconstants.BASIC)
					if parsedToken.MembershipDetails.Id != basicPlanID {

						if time.Now().After(parsedToken.MembershipDetails.PeriodEnd) {
							w.WriteHeader(http.StatusUnauthorized)
						} else {
							context.Set(r, "Claims", parsedToken)
							next.ServeHTTP(w, r)
						}
					} else {
						context.Set(r, "Claims", parsedToken)
						next.ServeHTTP(w, r)
					}
				}
			} else {
				w.WriteHeader(http.StatusUnauthorized)
			}
		}
	})
}
