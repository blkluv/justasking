package appconfigsrepo

import (
	"justasking/GO/core/model/appconfigs"
	"justasking/GO/core/startup/flight"
)

// GetAppConfig gets a single config value
func GetAppConfig(configType, configCode string) (appconfigsmodel.AppConfig, error) {
	db := flight.Context(nil, nil).DB

	var config appconfigsmodel.AppConfig
	err := db.Where("config_type = ? AND config_code = ?", configType, configCode).Find(&config).Error

	return config, err
}

// GetAppConfigs gets all config values of a given type
func GetAppConfigs(configType string) (map[string]string, error) {
	db := flight.Context(nil, nil).DB

	var configs map[string]string
	configs = make(map[string]string)

	config := []appconfigsmodel.AppConfig{}
	err := db.Where("config_type = ?", configType).Find(&config).Error

	for _, row := range config {
		configs[row.ConfigCode] = row.ConfigValue
	}

	return configs, err
}
