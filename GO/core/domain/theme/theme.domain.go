package themedomain

import (
	"fmt"

	"github.com/chande/justasking/common/operationresult"
	applogsdomain "github.com/chande/justasking/core/domain/applogs"
	thememodel "github.com/chande/justasking/core/model/theme"
	themerepo "github.com/chande/justasking/core/repo/theme"
)

var domainName = "ThemeDomain"

// GetTheme gets all boxes for a specific user
func GetTheme(themeId int) (thememodel.Theme, *operationresult.OperationResult) {
	functionName := "GetTheme"
	result := operationresult.New()

	theme, err := themerepo.GetTheme(themeId)

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting theme with id [%v]. Error: [%v]", themeId, msg), false)
	}

	return theme, result
}

// GetThemes gets all boxes for a specific user
func GetThemes() ([]thememodel.Theme, *operationresult.OperationResult) {
	functionName := "GetThemes"
	result := operationresult.New()

	themes, err := themerepo.GetThemes()

	if err != nil {
		msg := err.Error()
		result = operationresult.CreateErrorResult(msg, err)
		applogsdomain.LogError(domainName, functionName, fmt.Sprintf("Error getting all themes. Error: [%v]", msg), false)
	}

	return themes, result
}
