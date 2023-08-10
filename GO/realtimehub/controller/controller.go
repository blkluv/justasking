package controller

import (
	hubscontroller "github.com/chande/justasking/realtimehub/controller/hubs"
)

// LoadRoutes loads the routes for the controllers
func LoadRoutes() {
	hubscontroller.Load()
}
