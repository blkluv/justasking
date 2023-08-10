// Package boot handles the initialization of the web components.
package boot

import (
	"github.com/chande/justasking/realtimehub/controller"
	"github.com/chande/justasking/realtimehub/startup/flight"
)

// RegisterServices sets up all the web components.
func RegisterServices() {

	// Load the controller routes
	controller.LoadRoutes()

	//Create hubs for boxes
	flight.StoreHubs()
}
