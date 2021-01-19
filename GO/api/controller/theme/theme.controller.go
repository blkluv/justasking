package themecontroller

import (
	"encoding/json"
	"justasking/GO/api/startup/middleware"
	"justasking/GO/core/domain/theme"
	"justasking/GO/core/startup/flight"
	"net/http"

	"strconv"

	"github.com/blue-jay/core/router"
)

func Load() {
	router.Get("/theme/:themeid", GetTheme, middleware.AuthorizedHandler)
	router.Get("/themes", GetThemes, middleware.AuthorizedHandler)
}

// GetTheme gets data for a specific theme
func GetTheme(w http.ResponseWriter, r *http.Request) {
	context := flight.Context(w, r)
	themeId, err := strconv.Atoi(context.Param("themeid"))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		theme, result := themedomain.GetTheme(themeId)
		if result.IsSuccess() {
			responseString, err := json.Marshal(theme)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				w.Write([]byte(responseString))
			}
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

// GetThemes gets data for a specific theme
func GetThemes(w http.ResponseWriter, r *http.Request) {

	themes, result := themedomain.GetThemes()
	if result.IsSuccess() {
		responseString, err := json.Marshal(themes)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write([]byte(responseString))
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

}
