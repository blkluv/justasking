package applogsmodel

import (
	"time"
)

// AppLog is the model
type AppLog struct {
	ID          uint8
	LogType     string
	SourceName  string
	Message     string
	MachineName string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
