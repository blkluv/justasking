package userrepo

import (
	"justasking/GO/core/model/account"
	"justasking/GO/core/model/accountuser"
	"justasking/GO/core/model/idpjustasking"
	"justasking/GO/core/model/passwordresetrequest"
	"justasking/GO/core/model/user"
	"justasking/GO/core/repo/account"
	"justasking/GO/core/repo/accountuser"
	"justasking/GO/core/startup/flight"
	"time"

	uuid "github.com/satori/go.uuid"
)

// GetUserById gets a user from the users table by Id
func GetUserById(userId uuid.UUID) (usermodel.User, error) {
	user := usermodel.User{}
	account := accountmodel.Account{}
	db := flight.Context(nil, nil).DB

	//get user data
	err := db.Where("id = ? AND is_active = 1", userId).Find(&user).Error

	if err == nil {
		//get account id
		err = db.Raw(`SELECT a.id, a.owner_id, a.name, a.is_active, a.created_at, a.created_by, a.updated_at, a.updated_by, a.deleted_at 
			FROM accounts a JOIN account_users au ON a.id = au.account_id WHERE user_id = ? AND au.current_account = 1 AND au.is_active = 1`, user.ID).Scan(&account).Error
		if err == nil {
			user.Account = account
		}
	}

	return user, err
}

// GetSimpleUserRecord gets the actual user record, without getting plan, role, or account data
func GetSimpleUserRecord(userId uuid.UUID) (usermodel.User, error) {
	user := usermodel.User{}
	db := flight.Context(nil, nil).DB

	//get user data
	err := db.Where("id = ?", userId).Find(&user).Error

	return user, err
}

// GetUserByEmail gets a user from the users table by email
func GetUserByEmail(email string) (usermodel.User, error) {
	user := usermodel.User{}
	account := accountmodel.Account{}
	db := flight.Context(nil, nil).DB

	//get user data
	err := db.Where("email = ?", email).Find(&user).Error

	if err == nil {
		//get account id
		err = db.Raw(`SELECT a.id, a.owner_id, a.name, a.is_active, a.created_at, a.created_by, a.updated_at, a.updated_by, a.deleted_at 
			FROM accounts a JOIN account_users au ON a.id = au.account_id WHERE user_id = ?`, user.ID).Scan(&account).Error
		if err == nil {
			user.Account = account
		}
	}

	return user, err
}

// GetJustAskingIdpUser gets a record from the id_justasking table
func GetJustAskingIdpUser(email string) (idpjustaskingmodel.IdpJustAsking, error) {
	user := idpjustaskingmodel.IdpJustAsking{}

	db := flight.Context(nil, nil).DB

	//get user data
	err := db.Where("email = ?", email).Find(&user).Error

	return user, err
}

// InsertUser creates a user in the users table
func InsertUser(justAskingUser usermodel.User, account accountmodel.Account, accountUser accountusermodel.AccountUser) error {
	db := flight.Context(nil, nil).DB

	// Wrapping user creation in a transaction. We wouldn't want the user to be created without the account and accountuser record
	tx := db.Begin()

	if err := tx.Create(&justAskingUser).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := accountrepo.InsertAccount(account, tx); err != nil {
		tx.Rollback()
		return err
	}

	if err := accountuserrepo.InsertAccountUser(accountUser, tx); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

// GetJustAskingUserByGoogleSub gets a user based on a Google sub
func GetJustAskingUserByGoogleSub(sub string) (usermodel.User, error) {
	db := flight.Context(nil, nil).DB

	var user usermodel.User
	err := db.Raw(`SELECT users.id, first_name, last_name, users.email, last_logged_in_at, users.created_at, users.updated_at, users.deleted_at 
	FROM idp_google JOIN idp_users ON idp_google.sub = idp_users.sub JOIN users ON users.id = idp_users.user_id 
	WHERE idp_google.sub = ? AND users.is_active = 1`, sub).Scan(&user).Error

	return user, err
}

// UpdateUserLastLogin updates the last_logged_in field for a user
func UpdateUserLastLogin(userId uuid.UUID) error {
	db := flight.Context(nil, nil).DB

	err := db.Exec("UPDATE users SET last_logged_in_at = ? WHERE id = ?",
		time.Now(), userId).Error

	return err
}

// CreatePasswordResetRequest creates a new password reset request
func CreatePasswordResetRequest(request passwordresetrequestmodel.PasswordResetRequest) error {
	db := flight.Context(nil, nil).DB

	err := db.Create(&request).Error

	return err
}

// InvalidateOldPasswordResetRequests marks all previous password reset requests for a user as invalid
func InvalidateOldPasswordResetRequests(resetRequest passwordresetrequestmodel.PasswordResetRequest) error {
	db := flight.Context(nil, nil).DB

	err := db.Exec("UPDATE password_reset_requests SET is_active = 0 WHERE reset_code = ?", resetRequest.ResetCode).Error

	return err
}

// GetUserByResetCode returns a user attached to the reset cod passed in
func GetUserByResetCode(resetRequest passwordresetrequestmodel.PasswordResetRequest) (usermodel.User, error) {
	db := flight.Context(nil, nil).DB

	var user usermodel.User
	err := db.Raw(`SELECT u.id, first_name, last_name, email, image_url, last_logged_in_at 
		FROM users u JOIN password_reset_requests prr ON u.id = prr.user_id
		WHERE prr.reset_code = ?`, resetRequest.ResetCode).Scan(&user).Error

	return user, err
}

// GetResetRequest returns a password reset request
func GetResetRequest(resetRequest passwordresetrequestmodel.PasswordResetRequest) (passwordresetrequestmodel.PasswordResetRequest, error) {
	db := flight.Context(nil, nil).DB

	err := db.Raw(`SELECT id, user_id, reset_code, is_active, expires_at FROM password_reset_requests WHERE reset_code = ?`, resetRequest.ResetCode).Scan(&resetRequest).Error

	return resetRequest, err
}

// UpdatePassword updates a user's password
func UpdatePassword(resetRequest passwordresetrequestmodel.PasswordResetRequest, password string) error {
	db := flight.Context(nil, nil).DB

	// Wrapping user creation in a transaction. We wouldn't want the password to be updated without also updating the request
	tx := db.Begin()

	if err := tx.Exec(`UPDATE password_reset_requests SET is_active = 0 WHERE reset_code = ?`, resetRequest.ResetCode).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Exec(`UPDATE idp_justasking set password = ? WHERE sub = ?`, password, resetRequest.UserId).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
