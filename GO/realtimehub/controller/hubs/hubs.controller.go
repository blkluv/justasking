package hubscontroller

import (
	"io/ioutil"
	"justasking/GO/realtimehub/startup/flight"
	"net/http"

	"github.com/blue-jay/core/router"

	"justasking/GO/realtimehub/domain/hub"
)

// Load the routes.
func Load() {

	//Box Hub connection for regular user
	router.Post("/broadcast/:code", func(w http.ResponseWriter, r *http.Request) {
		context := flight.Context(w, r)
		activeHubs := flight.Context(w, r).HubMap
		hubName := context.Param("code")
		hubdomain.ReceiveAndBroadcastMessage(hubName, activeHubs, w, r)
	})

	//Box Hub connection for regular user
	router.Get("/hubs/box/:code", func(w http.ResponseWriter, r *http.Request) {
		context := flight.Context(w, r)
		activeHubs := flight.Context(w, r).HubMap
		hubName := context.Param("code")
		hubdomain.ServeWs(hubName, activeHubs, "user", w, r)
	})

	//Box Hub connection for dashboard
	router.Get("/hubs/box/dashboard/:code", func(w http.ResponseWriter, r *http.Request) {
		context := flight.Context(w, r)
		activeHubs := flight.Context(w, r).HubMap
		hubName := context.Param("code")
		hubdomain.ServeWs(hubName, activeHubs, "dashboard", w, r)
	})

	//Hub connection for account with multiple users
	router.Get("/hubs/account/:accountId", func(w http.ResponseWriter, r *http.Request) {
		context := flight.Context(w, r)
		activeHubs := flight.Context(w, r).HubMap
		hubName := context.Param("accountId")
		hubdomain.ServeWs(hubName, activeHubs, "account", w, r)
	})

	//Box Hub connection for dashboard
	router.Get("/.well-known/acme-challenge/:challenge", func(w http.ResponseWriter, r *http.Request) {
		response, _ := http.Get("http://storage.googleapis.com/justasking-web.appspot.com/certs/certificate.txt")
		responseData, _ := ioutil.ReadAll(response.Body)
		responseString := string(responseData)
		w.Write([]byte(responseString))
	})

}
