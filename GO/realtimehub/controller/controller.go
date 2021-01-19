package controller

import (
	"justasking/GO/realtimehub/controller/hubs"
)

// LoadRoutes loads the routes for the controllers
func LoadRoutes() {
	hubscontroller.Load()
}
