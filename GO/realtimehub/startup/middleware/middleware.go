// Package middleware - middleware handlers are supposed to be applied here if you wish every route to have a specific middleware
package middleware

import (
	"net/http"

	"github.com/blue-jay/core/router"
	"github.com/gorilla/context"
)

// SetUpMiddleware contains the middleware that applies to every request.
func SetUpMiddleware(h http.Handler) http.Handler {
	return router.ChainHandler( // Chain middleware, top middlware runs first
		h,                    // Handler to wrap
		context.ClearHandler, // Prevent memory leak with gorilla.sessions
	)
}
