package applogsrepo

import (
	applogsmodel "github.com/chande/justasking/core/model/applogs"
	"github.com/chande/justasking/core/startup/flight"
)

// InsertLog creates a record in the log table
func InsertLog(appLog applogsmodel.AppLog) {
	db := flight.Context(nil, nil).DB
	db.Create(&appLog)
}
