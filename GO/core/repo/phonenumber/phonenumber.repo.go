package phonenumberrepo

import (
	"justasking/GO/common/clients/twilio"
	"justasking/GO/core/model/phonenumber"
	"justasking/GO/core/model/twilio/phonenumber"
	"justasking/GO/core/startup/flight"
	"time"

	uuid "github.com/satori/go.uuid"
)

// InsertPhoneNumber inserts a phone number into the phone numbers table
func InsertPhoneNumber(phoneNumber twiliophonenumbermodel.TwilioPhoneNumber) error {
	db := flight.Context(nil, nil).DB

	newId, _ := uuid.NewV4()
	err := db.Exec(`INSERT INTO phone_numbers (id, sid, friendly_name, phone_number, region, iso_country, voice, sms, mms, is_active, created_at)
		SELECT ? as id, ? as sid, ? as friendly_name, ? as phone_number, ? as region, ? as iso_country, ? as voice, ? as sms, ? as mms, ? as is_active, ? as created_at 
		WHERE NOT EXISTS (SELECT id FROM phone_numbers WHERE sid = ?);`,
		newId, phoneNumber.Sid, phoneNumber.FriendlyName, phoneNumber.PhoneNumber, phoneNumber.Region, "US", phoneNumber.Capabilities.Voice,
		phoneNumber.Capabilities.Sms, phoneNumber.Capabilities.Mms, true, time.Now(), phoneNumber.Sid).Error

	return err
}

// GetAllActiveJustAskingPhoneNumbers gets all active phone numbers from our database
func GetAllActiveJustAskingPhoneNumbers() ([]phonenumbermodel.PhoneNumber, error) {
	db := flight.Context(nil, nil).DB

	phoneNumbers := []phonenumbermodel.PhoneNumber{}

	err := db.Raw(`SELECT id, sid, friendly_name, phone_number, region, iso_country, sms, mms, is_active, created_at, updated_at, deleted_at
		FROM phone_numbers WHERE is_active = 1`).Scan(&phoneNumbers).Error

	return phoneNumbers, err
}

// ReleasePhoneNumber finds a phone number which can be safely released, then releases and disables it
func ReleasePhoneNumber(twilioConfigs map[string]string) (phonenumbermodel.PhoneNumber, error) {
	db := flight.Context(nil, nil).DB
	var err error
	var justAskingPhoneNumber phonenumbermodel.PhoneNumber

	// find an active number which is not currently assigned to a box and is not within the cooldown window
	err = db.Raw(`SELECT pn.id, pn.sid, pn.friendly_name, pn.phone_number, pn.region, pn.iso_country, pn.voice, pn.sms, pn.mms, pn.is_active, pn.created_at, pn.updated_at, pn.deleted_at 
		FROM phone_numbers pn LEFT JOIN base_box_phone_numbers bbpn ON pn.id = bbpn.phone_number_id
		WHERE pn.is_active = 1 AND ((bbpn.is_active = 0 OR bbpn.is_active is NULL) AND (bbpn.updated_at < CURRENT_TIMESTAMP - INTERVAL 2 MINUTE OR bbpn.updated_at is NULL))
		ORDER BY bbpn.updated_at DESC LIMIT 1;`).Scan(&justAskingPhoneNumber).Error

	if err != nil {
		return justAskingPhoneNumber, err
	}

	tx := db.Begin()

	//update the is_active flag for this number
	if err := tx.Exec(`UPDATE phone_numbers SET is_active = 0 WHERE sid = ?`, justAskingPhoneNumber.Sid).Error; err != nil {
		tx.Rollback()
		return justAskingPhoneNumber, err
	}

	err = twilioclient.ReleasePhoneNumber(twilioConfigs, justAskingPhoneNumber.Sid)
	if err != nil {
		tx.Rollback()
		return justAskingPhoneNumber, err
	}

	tx.Commit()
	return justAskingPhoneNumber, err
}
