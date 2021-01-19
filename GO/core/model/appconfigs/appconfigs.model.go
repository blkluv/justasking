package appconfigsmodel

import (
	"time"
)

// AppConfig is the model for the app_configs table
type AppConfig struct {
	ID          uint8
	ConfigType  string
	ConfigCode  string
	ConfigValue string
	Comments    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
