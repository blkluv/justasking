package themerepo

import (
	"justasking/GO/core/model/theme"
	"justasking/GO/core/startup/flight"
)

// GetTheme gets all boxes for a specific user
func GetTheme(themeId int) (thememodel.Theme, error) {
	db := flight.Context(nil, nil).DB

	var theme thememodel.Theme

	err := db.Where("id = ?", themeId).Find(&theme).Error

	return theme, err
}

// GetThemes gets all boxes for a specific user
func GetThemes() ([]thememodel.Theme, error) {
	db := flight.Context(nil, nil).DB

	var theme []thememodel.Theme

	err := db.Find(&theme).Error

	return theme, err
}
