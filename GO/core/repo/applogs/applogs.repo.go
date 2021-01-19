package applogsrepo

import (
	"justasking/GO/core/model/applogs"
	"justasking/GO/core/startup/flight"
)

// InsertLog creates a record in the log table
func InsertLog(appLog applogsmodel.AppLog) {
	db := flight.Context(nil, nil).DB
	db.Create(&appLog)
}
