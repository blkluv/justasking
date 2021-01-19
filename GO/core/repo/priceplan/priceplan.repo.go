package priceplanrepo

import (
	"justasking/GO/core/model/customplanlicense"
	"justasking/GO/core/model/priceplan"
	"justasking/GO/core/startup/flight"

	uuid "github.com/satori/go.uuid"
)

// GetPublicPricePlans gets all price plans features
func GetPublicPricePlans() ([]priceplanmodel.PricePlan, error) {
	db := flight.Context(nil, nil).DB

	priceplans := []priceplanmodel.PricePlan{}

	err := db.Raw(`SELECT p.id,p.name,p.display_name,p.description,p.price,p.price_description,p.image_path,p.is_active,p.is_public,p.expires_in_days,p.sort_order,p.created_at,p.updated_at,p.deleted_at
				   FROM price_plans p WHERE is_public = 1 AND is_active = 1 ORDER BY sort_order ASC`).Scan(&priceplans).Error
	if err != nil {
		return priceplans, err
	}

	return priceplans, err
}

// GetPricePlanDetailsByPlanName gets price plan feature details
func GetPricePlanDetailsByPlanName(planName string) ([]priceplanmodel.PricePlan, error) {
	db := flight.Context(nil, nil).DB

	priceplan := []priceplanmodel.PricePlan{}

	err := db.Raw(`SELECT p.id, p.name, p.display_name, p.description, p.price, p.price_description, p.image_path, p.expires_in_days, f.name as feature_name, f.description as feature_description, 
		ppf.feature_value FROM price_plans p JOIN price_plan_features ppf ON p.id = ppf.plan_id JOIN features f ON f.id = ppf.feature_id 
		WHERE p.name = ? AND is_active = 1`, planName).Scan(&priceplan).Error
	if err != nil {
		return priceplan, err
	}

	return priceplan, err
}

// GetPricePlanDetailsByPlanId gets price plan featur details
func GetPricePlanDetailsByPlanId(guid uuid.UUID) ([]priceplanmodel.PricePlan, error) {
	db := flight.Context(nil, nil).DB

	priceplan := []priceplanmodel.PricePlan{}

	err := db.Raw(`SELECT p.id, p.name, p.display_name, p.description, p.price, p.price_description, p.image_path, p.expires_in_days, f.name as feature_name, f.description as feature_description, ppf.feature_value 
		FROM price_plans p JOIN price_plan_features ppf ON p.id = ppf.plan_id JOIN features f ON f.id = ppf.feature_id 
		WHERE p.id = ? AND is_active = 1`, guid).Scan(&priceplan).Error
	if err != nil {
		return priceplan, err
	}

	return priceplan, err
}

// GetPricePlanDetailsByAccountId gets price plan feature details for an account
func GetPricePlanDetailsByAccountId(accountId uuid.UUID) ([]priceplanmodel.PricePlan, error) {
	db := flight.Context(nil, nil).DB

	priceplan := []priceplanmodel.PricePlan{}

	err := db.Raw(`SELECT p.id, p.name, p.display_name, p.description, p.price, p.price_description, p.image_path, p.expires_in_days, f.name as feature_name, f.description as feature_description, ppf.feature_value, app.period_end 
		FROM price_plans p JOIN price_plan_features ppf ON p.id = ppf.plan_id JOIN features f ON f.id = ppf.feature_id JOIN account_price_plans app ON app.plan_id = p.id
		WHERE app.account_id = ? AND app.is_active = 1`, accountId).Scan(&priceplan).Error
	if err != nil {
		return priceplan, err
	}

	return priceplan, err
}

// GetPricePlanDetailsByBoxId gets price plan feature details for an account
func GetPricePlanDetailsByBoxId(boxId uuid.UUID) ([]priceplanmodel.PricePlan, error) {
	db := flight.Context(nil, nil).DB

	priceplan := []priceplanmodel.PricePlan{}

	err := db.Raw(`SELECT p.id, p.name, p.display_name, p.description, p.price, p.price_description, p.image_path, p.expires_in_days, f.name as feature_name, f.description as feature_description, ppf.feature_value 
		FROM price_plans p 
        JOIN price_plan_features ppf ON p.id = ppf.plan_id 
        JOIN features f ON f.id = ppf.feature_id 
        JOIN account_price_plans app ON app.plan_id = p.id
        JOIN base_box bb ON bb.account_id = app.account_id
		WHERE bb.id = ? AND app.is_active = 1;`, boxId).Scan(&priceplan).Error
	if err != nil {
		return priceplan, err
	}

	return priceplan, err
}

// GetPricePlanDetailsByUserId gets price plan feature details for an account
func GetPricePlanDetailsByUserId(userId uuid.UUID) ([]priceplanmodel.PricePlan, error) {
	db := flight.Context(nil, nil).DB

	priceplan := []priceplanmodel.PricePlan{}

	err := db.Raw(`SELECT p.id, p.name, p.display_name, p.description, p.price, p.price_description, p.image_path, p.expires_in_days, f.name as feature_name, f.description as feature_description, ppf.feature_value 
		FROM price_plans p JOIN price_plan_features ppf ON p.id = ppf.plan_id JOIN features f ON f.id = ppf.feature_id JOIN account_price_plans app ON app.plan_id = p.id 
        JOIN accounts a ON app.account_id = a.id JOIN account_users au ON au.account_id = a.id
		WHERE au.user_id = ? AND app.is_active = 1`, userId).Scan(&priceplan).Error
	if err != nil {
		return priceplan, err
	}

	return priceplan, err
}

// GetPricePlanByLicenseCode gets custom price plan details by license code
func GetPricePlanByLicenseCode(licenseCode string) ([]priceplanmodel.PricePlan, error) {
	db := flight.Context(nil, nil).DB

	priceplan := []priceplanmodel.PricePlan{}

	err := db.Raw(`SELECT p.id, p.name, p.display_name, p.description, p.price, p.price_description, p.image_path, p.expires_in_days, f.name as feature_name, f.description as feature_description, ppf.feature_value 
		FROM price_plans p JOIN price_plan_features ppf ON p.id = ppf.plan_id JOIN features f ON f.id = ppf.feature_id JOIN custom_plan_licenses cpl ON cpl.plan_id = p.id
		WHERE cpl.license_code = ? AND cpl.is_active=1 AND p.is_active = 1`, licenseCode).Scan(&priceplan).Error
	if err != nil {
		return priceplan, err
	}

	return priceplan, err
}

// GetCustomPlanLicense gets custom plan license details
func GetCustomPlanLicense(licenseCode string) (customplanlicensemodel.CustomPlanLicense, error) {
	db := flight.Context(nil, nil).DB

	var planLicense customplanlicensemodel.CustomPlanLicense

	err := db.Raw(`SELECT id, account_id, user_id, plan_id, license_code, is_active, created_at, created_by, updated_at, updated_by, deleted_at
		 FROM custom_plan_licenses WHERE license_code = ?`, licenseCode).Scan(&planLicense).Error

	return planLicense, err
}
