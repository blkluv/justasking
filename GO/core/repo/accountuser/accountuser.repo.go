package accountuserrepo

import (
	"justasking/GO/common/constants/role"
	"justasking/GO/core/model/account"
	"justasking/GO/core/model/accountinvitation"
	"justasking/GO/core/model/accountuser"
	"justasking/GO/core/repo/account"
	"justasking/GO/core/repo/userstripe"
	"justasking/GO/core/startup/flight"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// InsertAccountUser creates an account in the accounts table
func InsertAccountUser(accountUser accountusermodel.AccountUser, tx *gorm.DB) error {
	err := tx.Exec(`INSERT INTO account_users (account_id, user_id, role_id, is_active, current_account, token_version, created_at, created_by) 
	VALUES (?, ?, ?, 1, ?, ?, ?, ?)`, accountUser.AccountId, accountUser.UserId, accountUser.RoleId, accountUser.CurrentAccount, accountUser.TokenVersion, time.Now(), accountUser.UserId).Error

	return err
}

// GetUserAccounts gets all account_users records along with role name and account name
func GetUserAccounts(userId uuid.UUID) ([]accountusermodel.AccountUser, error) {
	db := flight.Context(nil, nil).DB

	accountUsers := []accountusermodel.AccountUser{}

	err := db.Raw(`SELECT u.account_id, a.name AS account_name, u.user_id, u.role_id, r.name AS role_name, u.is_active, u.current_account
		FROM account_users u JOIN accounts a ON u.account_id = a.id JOIN roles r ON u.role_id = r.id
		WHERE u.user_id = ? AND u.is_active = 1 AND a.is_active = 1`, userId).Scan(&accountUsers).Error

	return accountUsers, err
}

// UpdateAccountUserRole updates an account user's role
func UpdateAccountUserRole(accountUser accountusermodel.AccountUser, updatedBy uuid.UUID) error {
	db := flight.Context(nil, nil).DB

	newTokenVersion, _ := uuid.NewV4()
	err := db.Exec(`UPDATE account_users SET role_id = ?, token_version = ?, updated_at = ?, updated_by = ? WHERE account_id = ? AND user_id = ?`,
		accountUser.RoleId, newTokenVersion, time.Now().UTC(), updatedBy, accountUser.AccountId, accountUser.UserId).Error

	return err
}

// UpdateAccountUserToken updates an account user's token version
func UpdateAccountUserToken(userId uuid.UUID, accountId uuid.UUID) error {
	db := flight.Context(nil, nil).DB

	newTokenVersion, _ := uuid.NewV4()
	err := db.Exec(`UPDATE account_users SET token_version = ?, updated_at = ?, updated_by = ? WHERE account_id = ? AND user_id = ?`,
		newTokenVersion, time.Now().UTC(), userId, accountId, userId).Error

	return err
}

// GetAccountUser gets an accountUser record
func GetAccountUser(userId uuid.UUID, accountId uuid.UUID) (accountusermodel.AccountUser, error) {
	db := flight.Context(nil, nil).DB

	var accountUser accountusermodel.AccountUser

	err := db.Raw(`SELECT au.account_id, au.user_id, au.role_id, r.name AS role_name, au.is_active, au.current_account, au.token_version, au.created_at, au.created_by, au.updated_at, u.email, au.updated_by, au.deleted_at
		FROM account_users au JOIN users u ON au.user_id = u.id JOIN roles r ON au.role_id = r.id
		WHERE account_id = ? AND user_id = ?`, accountId, userId).Scan(&accountUser).Error

	return accountUser, err
}

// TransferOwnership transfers ownership from one user to another
func TransferOwnership(oldOwnerId uuid.UUID, newOwnerId uuid.UUID, accountId uuid.UUID, newOwnerStripeCustomerId string, newAccountName string, needsNewAccount bool, stripeCustomerId string) error {
	db := flight.Context(nil, nil).DB
	var err error

	tx := db.Begin()

	//set old owner to admin
	oldOwnerTokenVersion, _ := uuid.NewV4()
	if err := tx.Exec(`UPDATE account_users SET role_id = ?, token_version = ?, updated_at = ?, updated_by = ? WHERE account_id = ? AND user_id = ?`, roleconstants.ADMIN, oldOwnerTokenVersion, time.Now().UTC(), oldOwnerId, accountId, oldOwnerId).Error; err != nil {
		tx.Rollback()
		return err
	}

	//set new owner to owner
	newOwnerTokenVersion, _ := uuid.NewV4()
	if err := tx.Exec(`UPDATE account_users SET role_id = ?, token_version = ?, updated_at = ?, updated_by = ? WHERE account_id = ? AND user_id = ?`, roleconstants.OWNER, newOwnerTokenVersion, time.Now().UTC(), oldOwnerId, accountId, newOwnerId).Error; err != nil {
		tx.Rollback()
		return err
	}

	//update ownerid of oldowner's account record
	if err := tx.Exec(`UPDATE accounts SET owner_id = ?, updated_at = ?, updated_by = ? WHERE id = ?`, newOwnerId, time.Now().UTC(), oldOwnerId, accountId).Error; err != nil {
		tx.Rollback()
		return err
	}

	//add stripe mapping for
	if err := userstriperepo.InsertUserStripeMapping(newOwnerId, accountId, newOwnerStripeCustomerId, tx); err != nil {
		tx.Rollback()
		return err
	}

	if needsNewAccount {
		//create new account for oldOwner
		var account accountmodel.Account
		account.Id, _ = uuid.NewV4()
		account.OwnerId = oldOwnerId
		account.Name = newAccountName
		account.IsActive = true
		account.CreatedAt = time.Now().UTC()
		account.CreatedBy = oldOwnerId.String()

		if err := tx.Create(&account).Error; err != nil {
			tx.Rollback()
			return err
		}

		//insert account user for newly created account
		var accountUser accountusermodel.AccountUser
		accountUser.AccountId = account.Id
		accountUser.UserId = oldOwnerId
		accountUser.RoleId, _ = uuid.FromString(roleconstants.OWNER)
		accountUser.IsActive = true
		accountUser.CurrentAccount = false
		accountUser.CreatedAt = time.Now()
		accountUser.CreatedBy = oldOwnerId.String()

		if err := InsertAccountUser(accountUser, tx); err != nil {
			tx.Rollback()
			return err
		}

		if err := accountrepo.InsertDefaultAccountPricePlan(account.Id, tx); err != nil {
			tx.Rollback()
			return err
		}

		if err := userstriperepo.InsertUserStripeMapping(oldOwnerId, account.Id, stripeCustomerId, tx); err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()

	return err
}

// RemoveUserFromAccount removes a user from an account
func RemoveUserFromAccount(userId uuid.UUID, accountId uuid.UUID, updatedBy uuid.UUID, needsCurrentAccountUpdated bool) error {
	db := flight.Context(nil, nil).DB
	var err error

	tx := db.Begin()

	newToken, _ := uuid.NewV4()
	if err := tx.Exec(`UPDATE account_users SET is_active = 0, current_account = 0, token_version = ?, updated_at = ?, updated_by = ? WHERE user_id = ? AND account_id = ?`, newToken, time.Now().UTC(), updatedBy, userId, accountId).Error; err != nil {
		tx.Rollback()
		return err
	}

	if needsCurrentAccountUpdated {

		// update current_account to first active accountUser record for this user
		if err := tx.Exec(`UPDATE account_users SET current_account = 1, updated_at = ?, updated_by = ? WHERE user_id = ? AND is_active = 1 LIMIT 1`, time.Now().UTC(), updatedBy, userId).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()

	return err
}

// UpdateUserForPlanExpiration updates a user's current account and token_version for an account that recently expired
func UpdateUserForPlanExpiration(userId uuid.UUID, accountId uuid.UUID, needsCurrentAccountUpdated bool) error {
	db := flight.Context(nil, nil).DB
	var err error

	tx := db.Begin()

	newToken, _ := uuid.NewV4()
	if err := tx.Exec(`UPDATE account_users SET current_account = 0, token_version = ?, updated_at = ?, updated_by = ? WHERE user_id = ? AND account_id = ?`, newToken, time.Now().UTC(), "SyncDomain", userId, accountId).Error; err != nil {
		tx.Rollback()
		return err
	}

	if needsCurrentAccountUpdated {

		// update current_account to first active accountUser record for this user
		if err := tx.Exec(`UPDATE account_users SET current_account = 1, updated_at = ?, updated_by = ? WHERE user_id = ? AND is_active = 1 LIMIT 1`, time.Now().UTC(), "SyncDomain", userId).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()

	return err
}

// UpdateCurrentAccount updates the user's current account
func UpdateCurrentAccount(userId uuid.UUID, accountId uuid.UUID) error {
	db := flight.Context(nil, nil).DB
	var err error

	tx := db.Begin()

	//set accounts to 0
	if err := tx.Exec(`UPDATE account_users SET current_account = 0 WHERE user_id = ?`, userId).Error; err != nil {
		tx.Rollback()
		return err
	}

	//update new current_account to 1
	if err := tx.Exec(`UPDATE account_users SET current_account = 1 WHERE user_id = ? and account_id = ?`, userId, accountId).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return err
}

// GetActiveAccountUsers gets all active users for an account
func GetActiveAccountUsers(accountId uuid.UUID) ([]accountusermodel.AccountUser, error) {
	db := flight.Context(nil, nil).DB
	var accountUsers []accountusermodel.AccountUser

	err := db.Raw(`SELECT account_id, user_id, role_id, r.name as role_name, is_active, au.created_at, au.created_by, au.updated_at, au.updated_by, au.deleted_at
		FROM justasking.account_users au JOIN roles r ON au.role_id = r.id
		WHERE is_active = 1 AND account_id = ?`, accountId).Scan(&accountUsers).Error

	return accountUsers, err
}

// RedeemAccountInvitation deactivates an invitation and adds a user to an account
func RedeemAccountInvitation(invitation accountinvitationmodel.AccountInvitation, userId uuid.UUID, userAccountExists bool) error {
	db := flight.Context(nil, nil).DB
	var err error

	tx := db.Begin()

	//update account invite
	if err := tx.Exec(`UPDATE account_invitations SET is_active = 0 WHERE invitation_code = ? `, invitation.InvitationCode).Error; err != nil {
		tx.Rollback()
		return err
	}

	//update old current_account value
	if err := tx.Exec(`UPDATE account_users SET current_account = 0 WHERE user_id = ?`, userId).Error; err != nil {
		tx.Rollback()
		return err
	}

	if userAccountExists {
		//update old account_user value
		if err := tx.Exec(`UPDATE account_users SET current_account = 1, is_active = 1, role_id = ? WHERE user_id = ? AND account_id = ?`, invitation.RoleId, userId, invitation.AccountId).Error; err != nil {
			tx.Rollback()
			return err
		}
	} else {
		//insert account user
		var accountUser accountusermodel.AccountUser
		accountUser.AccountId = invitation.AccountId
		accountUser.UserId = userId
		accountUser.RoleId = invitation.RoleId
		accountUser.IsActive = true
		accountUser.CurrentAccount = true
		accountUser.CreatedAt = time.Now()
		accountUser.CreatedBy = userId.String()
		accountUser.TokenVersion, _ = uuid.NewV4()

		if err := InsertAccountUser(accountUser, tx); err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return err
}

// GetCurrentAccount gets a user's current AccountUser record
func GetCurrentAccount(userId uuid.UUID) (accountusermodel.AccountUser, error) {
	db := flight.Context(nil, nil).DB
	var accountUser accountusermodel.AccountUser

	err := db.Raw(`SELECT account_id, user_id, role_id, is_active, current_account, created_at, token_version, created_by, updated_at, updated_by, deleted_at
		FROM justasking.account_users WHERE is_active = 1 AND current_account = 1 AND user_id = ?`, userId).Scan(&accountUser).Error

	return accountUser, err
}
