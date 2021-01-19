// Package boot handles the initialization of the web components.
package boot

import (
	"github.com/jinzhu/gorm"
	//To be used at a later time?
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"fmt"
	"justasking/GO/core/startup/env"
	"justasking/GO/core/startup/flight"
)

// RegisterServices sets up all the web components.
func RegisterServices(config *env.Info) {

	// Connect to the MySQL database
	mysqlDB, err := gorm.Open(config.Settings["DatabaseType"], config.Settings["ConnectionString"])

	if err != nil {
		fmt.Println(err)
	}

	// Store the variables in flight
	flight.StoreConfig(*config)

	// Store the database connection in flight
	flight.StoreDB(mysqlDB)
}
