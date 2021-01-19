package appconfigsdomain

import (
	"fmt"
	"justasking/GO/common/operationresult"
	"justasking/GO/core/domain/applogs"
	"justasking/GO/core/model/appconfigs"
	"justasking/GO/core/repo/appconfigs"
)

var domainName = "AppConfigsDomain"

// GetAppConfig calls the repo to get a singl app config
func GetAppConfig(configType, configCode string) (appconfigsmodel.AppConfig, *operationresult.OperationResult) {
	functionName := "GetAppConfig"
	result := operationresult.New()

	retVal, err := appconfigsrepo.GetAppConfig(configType, configCode)
	if err != nil {
		msg := fmt.Sprintf("%v [%v] %v [%v] %v [%v]", "Error getting Config Value for ConfigType:", configType, "ConfigCode:", configCode, "Error:", err.Error())
		//result.SetError(msg, err)
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, msg, false)
	}

	return retVal, result
}

// GetAppConfigs calls the repo to get all app configs
func GetAppConfigs(configType string) (map[string]string, *operationresult.OperationResult) {
	functionName := "GetAppConfigs"
	result := operationresult.New()

	retVal, err := appconfigsrepo.GetAppConfigs(configType)
	if err != nil {
		msg := fmt.Sprintf("%v [%v] %v [%v]", "Error getting Config Value for ConfigType:", configType, "Error:", err.Error())
		//result.SetError(msg, err)
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, msg, false)
	}

	return retVal, result
}
