// Package flight provides access to the application settings safely.
package flight

import (
	"net/http"
	"sync"

	"github.com/blue-jay/core/router"
	"github.com/jinzhu/gorm"

	"justasking/GO/realtimehub/model/hubs"
)

var (
	dbInfo     *gorm.DB
	mutex      sync.RWMutex
	activeHubs *hubs.HubMap
)

// StoreHubs stores the hubs that will be used for boxes
func StoreHubs() {
	mutex.Lock()
	activeHubs = &hubs.HubMap{Hubs: make(map[string]*hubs.Hub)}
	mutex.Unlock()
}

// Info structures the application settings.
type Info struct {
	W      http.ResponseWriter
	R      *http.Request
	DB     *gorm.DB
	HubMap *hubs.HubMap
}

// Context returns the application settings.
func Context(w http.ResponseWriter, r *http.Request) Info {

	mutex.RLock()
	i := Info{
		W:      w,
		R:      r,
		DB:     dbInfo,
		HubMap: activeHubs,
	}
	mutex.RUnlock()

	return i
}

// Reset will delete all package globals
func Reset() {
	mutex.Lock()
	dbInfo = &gorm.DB{}
	mutex.Unlock()
}

// Param gets the URL parameter.
func (c *Info) Param(name string) string {
	return router.Param(c.R, name)
}

// Redirect sends a temporary redirect.
func (c *Info) Redirect(urlStr string) {
	http.Redirect(c.W, c.R, urlStr, http.StatusFound)
}
