// Package middleware - middleware handlers are supposed to be applied here if you wish every route to have a specific middleware
package middleware

import (
	authenticationclaim "justasking/GO/common/authenticationclaim"
	"net/http"

	"github.com/blue-jay/core/router"
	"github.com/gorilla/context"
)

// SetUpMiddleware contains the middleware that applies to every request.
func SetUpMiddleware(h http.Handler) http.Handler {
	return router.ChainHandler( // Chain middleware, top middlware runs first
		h,                    // Handler to wrap
		LogRequestHandler,    // Log every request
		context.ClearHandler, // Prevent memory leak with gorilla.sessions
	)
}

// GetTokenClaims returns the parsed object for the current request
func GetTokenClaims(r *http.Request) *authenticationclaim.AuthenticationClaim {
	return context.Get(r, "Claims").(*authenticationclaim.AuthenticationClaim)
}
