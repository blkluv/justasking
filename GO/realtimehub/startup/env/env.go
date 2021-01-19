// Package env reads the application settings.
package env

import (
	"encoding/json"

	"github.com/blue-jay/core/email"
	"github.com/blue-jay/core/jsonconfig"
	"github.com/blue-jay/core/server"
)

// *****************************************************************************
// Application Settings
// *****************************************************************************

// Info structures the application settings.
type Info struct {
	Email email.Info `json:"Email"`
	//	MySQL    mysql.Info        `json:"MySQL"`
	Server   server.Info       `json:"Server"`
	Settings map[string]string `json:"AppSettings"`
	path     string
}

// Path returns the env.json path
func (c *Info) Path() string {
	return c.path
}

// ParseJSON unmarshals bytes to structs
func (c *Info) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}

// New returns a instance of the application settings.
func New(path string) *Info {
	return &Info{
		path: path,
	}
}

// LoadConfig reads the configuration file.
func LoadConfig(configFile string) (*Info, error) {
	// Create a new configuration with the path to the file
	config := New(configFile)

	// Load the configuration file
	err := jsonconfig.Load(configFile, config)

	// Return the configuration
	return config, err
}
