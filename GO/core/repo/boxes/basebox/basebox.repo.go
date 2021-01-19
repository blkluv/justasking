package baseboxrepo

import (
	"justasking/GO/core/model/boxes/basebox"
	"justasking/GO/core/model/phonenumber"
	"justasking/GO/core/startup/flight"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// GetBoxesByAccountId gets all boxes for a specific user
func GetBoxesByAccountId(accountId uuid.UUID) ([]baseboxmodel.BaseBox, error) {
	db := flight.Context(nil, nil).DB

	var boxes []baseboxmodel.BaseBox

	err := db.Raw(`SELECT bb.id, bb.code, bb.original_code, bb.account_id, bb.box_type as box_type_id, bt.name as box_type, bb.theme_id, t.value as theme, 
	bb.is_live, bb.created_at, bb.created_by, bb.updated_at, bb.updated_by, bb.deleted_at
	FROM base_box bb 
	JOIN themes t ON bb.theme_id = t.id 
	JOIN box_type bt ON bt.id = bb.box_type
	WHERE bb.account_id = ? AND bb.deleted_at IS NULL
	order by bb.created_at desc`, accountId).Scan(&boxes).Error

	return boxes, err
}

// GetBoxesByUserId gets all boxes for a specific user on a specific account
func GetBoxesByUserId(userId uuid.UUID, accountId uuid.UUID) ([]baseboxmodel.BaseBox, error) {
	db := flight.Context(nil, nil).DB

	var boxes []baseboxmodel.BaseBox

	err := db.Raw(`SELECT bb.id, bb.code, bb.original_code, bb.account_id, bb.box_type as box_type_id, bt.name as box_type, bb.theme_id, t.value as theme, 
	bb.is_live, bb.created_at, bb.created_by, bb.updated_at, bb.updated_by, bb.deleted_at
	FROM base_box bb 
	JOIN themes t ON bb.theme_id = t.id 
	JOIN box_type bt ON bt.id = bb.box_type
	WHERE bb.account_id = ? AND bb.created_by = ? AND bb.deleted_at IS NULL
	order by bb.created_at desc`, accountId, userId).Scan(&boxes).Error

	return boxes, err
}

// GetActiveBoxesByAccountId gets all active boxes for a specific user
func GetActiveBoxesByAccountId(accountId uuid.UUID) ([]baseboxmodel.BaseBox, error) {
	db := flight.Context(nil, nil).DB

	var boxes []baseboxmodel.BaseBox

	err := db.Raw(`SELECT bb.id, bb.code, bb.original_code, bb.account_id, bb.box_type as box_type_id, bt.name as box_type, bb.theme_id, t.value as theme, 
	bb.is_live, bb.created_at, bb.created_by, bb.updated_at, bb.updated_by, bb.deleted_at
	FROM base_box bb 
	JOIN themes t ON bb.theme_id = t.id 
	JOIN box_type bt ON bt.id = bb.box_type
	WHERE bb.account_id = ? AND bb.is_live = 1 AND bb.deleted_at IS NULL
	order by bb.created_at desc`, accountId).Scan(&boxes).Error

	return boxes, err
}

// GetBaseBoxByCode gets the base box for a specific code
func GetBaseBoxByCode(code string) (baseboxmodel.BaseBox, error) {
	db := flight.Context(nil, nil).DB

	var box baseboxmodel.BaseBox

	err := db.Raw(`SELECT DISTINCT 
	bb.id, 
	bb.code, 
	bb.original_code,
    bb.account_id, 
    bb.box_type as box_type_id, 
    bt.name as box_type, 
    bb.theme_id, t.value as theme, 
	bb.is_live, 
    bb.entry_page_enabled, 
    bb.presentation_page_enabled, 
    bb.login_required, 
    bb.sms_enabled, 
    IF(bb.is_live = false AND bbpn.is_active = true, null, pn.phone_number) as phone_number, 
    bb.created_at, 
    bb.created_by, 
    bb.updated_at, 
    bb.updated_by, 
    bb.deleted_at
	FROM base_box bb 
	JOIN themes t ON bb.theme_id = t.id 
	JOIN box_type bt ON bt.id = bb.box_type
	LEFT JOIN base_box_phone_numbers bbpn ON bb.id = bbpn.base_box_id  AND bbpn.is_active = true
	LEFT JOIN phone_numbers pn ON pn.id = bbpn.phone_number_id
	WHERE bb.code = ? AND bb.deleted_at IS NULL`, code).Scan(&box).Error

	return box, err
}

// GetBaseBoxById gets the base box for a specific code
func GetBaseBoxById(boxId uuid.UUID) (baseboxmodel.BaseBox, error) {
	db := flight.Context(nil, nil).DB

	var box baseboxmodel.BaseBox

	err := db.Raw(`SELECT DISTINCT 
	bb.id, 
	bb.code,
	bb.original_code,
    bb.account_id, 
    bb.box_type as box_type_id, 
    bt.name as box_type, 
    bb.theme_id, t.value as theme, 
	bb.is_live, 
    bb.entry_page_enabled, 
    bb.presentation_page_enabled, 
    bb.login_required, 
    bb.sms_enabled, 
    IF(bb.is_live = false AND bbpn.is_active = true, null, pn.phone_number) as phone_number, 
    bb.created_at, 
    bb.created_by, 
    bb.updated_at, 
    bb.updated_by, 
    bb.deleted_at
	FROM base_box bb 
	JOIN themes t ON bb.theme_id = t.id 
	JOIN box_type bt ON bt.id = bb.box_type
	LEFT JOIN base_box_phone_numbers bbpn ON bb.id = bbpn.base_box_id  AND bbpn.is_active = true
	LEFT JOIN phone_numbers pn ON pn.id = bbpn.phone_number_id
 	WHERE bb.id = ? AND bb.deleted_at IS NULL`, boxId).Scan(&box).Error

	return box, err
}

// InsertBaseBox creates a base box
func InsertBaseBox(baseBox baseboxmodel.BaseBox, tx *gorm.DB) error {
	err := tx.Exec(`INSERT INTO base_box (id, code, original_code, account_id, box_type, theme_id, is_live, entry_page_enabled, presentation_page_enabled, login_required, sms_enabled, created_by) 
						VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		baseBox.ID, baseBox.Code, baseBox.Code, baseBox.AccountId, baseBox.BoxTypeId, baseBox.ThemeId, baseBox.IsLive, baseBox.EntryPageEnabled, baseBox.PresentationPageEnabled, baseBox.LoginRequired, baseBox.SmsEnabled, baseBox.CreatedBy).Error

	return err
}

// BoxCodeExists gets all boxes for a specific code
func BoxCodeExists(code string) (bool, error) {
	db := flight.Context(nil, nil).DB

	var count int
	var exists = false
	var box baseboxmodel.BaseBox

	err := db.Where("code = ? AND deleted_at IS NULL", code).Find(&box).Count(&count).Error

	if count == 1 {
		exists = true
	}

	return exists, err
}

// GetBaseBoxByPhoneNumber gets the base box for a specific phone number
func GetBaseBoxByPhoneNumber(phoneNumber string) (baseboxmodel.BaseBox, error) {
	db := flight.Context(nil, nil).DB

	var box baseboxmodel.BaseBox

	err := db.Raw(`SELECT b.id, b.code, b.original_code, b.account_id, b.box_type as box_type_id, b.theme_id, b.is_live, b.entry_page_enabled, b.presentation_page_enabled, b.login_required, 
	b.sms_enabled, b.created_at, b.created_by, b.updated_at, b.updated_by, b.deleted_at FROM phone_numbers p JOIN base_box_phone_numbers bp ON p.id = bp.phone_number_id 
	JOIN base_box b ON b.id = bp.base_box_id WHERE p.phone_number = ? AND bp.is_active = 1 AND b.deleted_at IS NULL;`, phoneNumber).Scan(&box).Error

	return box, err
}

// ActivateBaseBoxAndAssignNumber sets is_live flag to TRUE and assigns a phone number
func ActivateBaseBoxAndAssignNumber(boxId uuid.UUID, updatedBy uuid.UUID) (phonenumbermodel.PhoneNumber, error) {
	db := flight.Context(nil, nil).DB

	phoneNumber := phonenumbermodel.PhoneNumber{}

	// Wrapping box activation in a transaction. We wouldn't want one of the votes to be created without the other.
	tx := db.Begin()

	if err := db.Exec("CALL AssignPhoneNumber(?, ?)", boxId, updatedBy).Error; err != nil {
		tx.Rollback()
		return phoneNumber, err
	}

	if err := db.Raw(`SELECT p.id, p.sid, p.friendly_name, p.phone_number, p.region, p.iso_country, p.voice, p.sms, p.mms FROM phone_numbers p JOIN base_box_phone_numbers bp ON p.id = bp.phone_number_id WHERE bp.base_box_id = ? AND bp.is_active = 1;`, boxId).Scan(&phoneNumber).Error; err != nil {
		tx.Rollback()
		return phoneNumber, err
	}

	tx.Commit()
	return phoneNumber, nil
}

// ActivateBaseBox activates a base box but does not assign a number
func ActivateBaseBox(boxId uuid.UUID, updatedBy uuid.UUID) error {
	db := flight.Context(nil, nil).DB

	err := db.Exec("UPDATE base_box SET is_live = 1, updated_at = CURRENT_TIMESTAMP, updated_by = ? WHERE id = ?", updatedBy, boxId).Error

	return err
}

// DeactivateBaseBox sets is_live flag to TRUE
func DeactivateBaseBox(boxId uuid.UUID, updatedBy uuid.UUID) error {
	db := flight.Context(nil, nil).DB

	// Wrapping deactivation in a transaction. We wouldn't want one of the votes to be created without the other.
	tx := db.Begin()

	if err := db.Exec("UPDATE base_box SET is_live = 0, updated_at = ?, updated_by = ? WHERE id = ?", time.Now(), updatedBy, boxId).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := db.Exec("UPDATE base_box_phone_numbers SET is_active = 0, updated_at = ?, updated_by = ? WHERE base_box_id = ?", time.Now(), updatedBy, boxId).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

// DeactivateAllBaseBoxesByAccountId deactivates all boxes for an account
func DeactivateAllBaseBoxesByAccountId(accountId uuid.UUID) error {
	db := flight.Context(nil, nil).DB

	err := db.Exec(`UPDATE base_box bb JOIN base_box_phone_numbers bbpn ON bb.id = bbpn.base_box_id 
		SET bb.is_live = 0, bbpn.is_active = 0, bb.updated_at = CURRENT_TIMESTAMP, bbpn.updated_at = CURRENT_TIMESTAMP, bb.updated_by = 'Plan Expiration', bbpn.updated_by = 'Plan Expiration'
		WHERE bb.account_id = ?`, accountId).Error

	return err
}

// GetBaseBoxByAnswerBoxQuestionId gets a base box by an answer box question id
func GetBaseBoxByAnswerBoxQuestionId(questionId uuid.UUID) (baseboxmodel.BaseBox, error) {
	db := flight.Context(nil, nil).DB

	var box baseboxmodel.BaseBox

	err := db.Raw(`SELECT DISTINCT 
		bb.id, 
		bb.code, 
		bb.original_code,
		bb.account_id, 
		bb.box_type as box_type_id, 
		bb.theme_id,
		bb.is_live, 
		bb.entry_page_enabled, 
		bb.presentation_page_enabled, 
		bb.login_required, 
		bb.sms_enabled,
		bb.created_at, 
		bb.created_by, 
		bb.updated_at, 
		bb.updated_by, 
		bb.deleted_at
		FROM base_box bb 
		JOIN answer_box_questions abq ON bb.id = abq.box_id
		WHERE abq.question_id = ? AND bb.deleted_at IS NULL`, questionId).Scan(&box).Error

	return box, err
}

// DeleteBaseBox marks a basebox as deleted
func DeleteBaseBox(boxId uuid.UUID, updatedBy uuid.UUID) error {
	db := flight.Context(nil, nil).DB

	// Wrapping deletion in a transaction. We wouldn't want the box to get deleted without being closed first.
	tx := db.Begin()

	if err := db.Exec("UPDATE base_box SET is_live = 0, updated_at = ?, updated_by = ? WHERE id = ?", time.Now(), updatedBy, boxId).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := db.Exec("UPDATE base_box_phone_numbers SET is_active = 0, updated_at = ?, updated_by = ? WHERE base_box_id = ?", time.Now(), updatedBy, boxId).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := db.Exec(`UPDATE base_box SET code = id, deleted_at = ?, updated_at = ?, updated_by = ? WHERE id = ?`, time.Now().UTC(), time.Now().UTC(), updatedBy, boxId).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
