// Package flight provides access to the application settings safely.
package flight

import (
	"net/http"
	"sync"

	"github.com/blue-jay/core/router"
	"github.com/jinzhu/gorm"

	"justasking/GO/core/startup/env"
)

var (
	configInfo env.Info
	dbInfo     *gorm.DB
	mutex      sync.RWMutex
)

// StoreConfig stores the application settings so controller functions can
//access them safely.
func StoreConfig(ci env.Info) {
	mutex.Lock()
	configInfo = ci
	mutex.Unlock()
}

// StoreDB stores the database connection settings so controller functions can
// access them safely.
func StoreDB(db *gorm.DB) {
	mutex.Lock()
	dbInfo = db
	mutex.Unlock()
}

// Info structures the application settings.
type Info struct {
	Config env.Info
	W      http.ResponseWriter
	R      *http.Request
	DB     *gorm.DB
}

// Context returns the application settings.
func Context(w http.ResponseWriter, r *http.Request) Info {

	mutex.RLock()
	i := Info{
		W:      w,
		R:      r,
		DB:     dbInfo,
		Config: configInfo,
	}
	mutex.RUnlock()

	return i
}

// Reset will delete all package globals
func Reset() {
	mutex.Lock()
	configInfo = env.Info{}
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
