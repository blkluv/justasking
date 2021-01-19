package accountrepo

import (
	"justasking/GO/common/constants/priceplan"
	"justasking/GO/core/model/account"
	"justasking/GO/core/model/accountinvitation"
	"justasking/GO/core/model/userstripe"
	"justasking/GO/core/startup/flight"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// GetAccount gets an account
func GetAccount(accountId uuid.UUID) (accountmodel.Account, error) {
	db := flight.Context(nil, nil).DB

	account := accountmodel.Account{}

	err := db.Raw(`SELECT id, owner_id, name, is_active, created_at, created_by, updated_at, updated_by, deleted_at
		FROM accounts 
		WHERE id = ?`, accountId).Scan(&account).Error

	return account, err
}

// InsertAccount creates an account in the accounts table
func InsertAccount(account accountmodel.Account, tx *gorm.DB) error {
	err := tx.Create(&account).Error

	return err
}

// InsertDefaultAccountPricePlan is assigns the starter price plan when an account is created
func InsertDefaultAccountPricePlan(accountId uuid.UUID, tx *gorm.DB) error {

	newId, _ := uuid.NewV4()
	err := tx.Exec(`INSERT INTO account_price_plans (id, account_id, plan_id, is_active, created_at) VALUES (?, ?, ?, 1, CURRENT_TIMESTAMP)`, newId, accountId, priceplanconstants.BASIC).Error

	return err
}

// UpdateAccountPricePlan updates an account/priceplan mapping
func UpdateAccountPricePlanTx(accountId uuid.UUID, priceplanId uuid.UUID, endDate *time.Time, tx *gorm.DB) error {
	//set other plans to 0
	err := tx.Exec(`UPDATE account_price_plans SET is_active = 0, updated_at = CURRENT_TIMESTAMP WHERE account_id = ? AND is_active = 1`, accountId).Error

	// insert new plan
	newId, _ := uuid.NewV4()
	err = tx.Exec(`INSERT INTO account_price_plans (id, account_id, plan_id, is_active, period_end, created_at) VALUES (?, ?, ?, 1, ?, CURRENT_TIMESTAMP)`, newId, accountId, priceplanId, endDate).Error

	return err
}

// UpdateAccountPricePlan updates an account/priceplan mapping
func UpdateAccountPricePlan(accountId uuid.UUID, priceplanId uuid.UUID) error {
	db := flight.Context(nil, nil).DB
	var err error

	tx := db.Begin()

	//set other plans to 0
	if err := tx.Exec(`UPDATE account_price_plans SET is_active = 0, updated_at = CURRENT_TIMESTAMP WHERE account_id = ? AND is_active = 1`, accountId).Error; err != nil {
		tx.Rollback()
		return err
	}

	//insert new plan
	newId, _ := uuid.NewV4()
	if err := tx.Exec(`INSERT INTO account_price_plans (id, account_id, plan_id, is_active, created_at) VALUES (?, ?, ?, 1, CURRENT_TIMESTAMP)`, newId, accountId, priceplanId).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return err
}

// GetAccountsByUserId gets accounts for a given user id
func GetAccountsByUserId(id uuid.UUID) ([]accountmodel.Account, error) {
	db := flight.Context(nil, nil).DB

	accounts := []accountmodel.Account{}

	err := db.Raw(`SELECT a.id, a.owner_id, a.name, a.is_active, a.created_by, a.updated_at, a.deleted_at, u.current_account
					FROM justasking.accounts a JOIN justasking.account_users u ON a.id = u.account_id WHERE u.user_id = ? AND u.is_active = 1`, id).Scan(&accounts).Error

	return accounts, err
}

// GetAccountsByEmail gets accounts for a given email address
func GetAccountsByEmail(id uuid.UUID) ([]accountmodel.Account, error) {
	db := flight.Context(nil, nil).DB

	accounts := []accountmodel.Account{}

	err := db.Raw(`SELECT a.id, .owner_id, a.name, a.is_active, a.created_by, a.updated_at, a.deleted_at 
				   FROM justasking.accounts a JOIN justasking.account_users au ON a.id = au.account_id JOIN justasking.users u ON au.user_id = u.id
				   WHERE u.email = ?`, id).Scan(&accounts).Error

	return accounts, err
}

// GetAccountByAnswerBoxEntry returns the account associated with the box associated with the question associated with the entry
func GetAccountByAnswerBoxEntry(entryId uuid.UUID) (accountmodel.Account, error) {
	db := flight.Context(nil, nil).DB

	account := accountmodel.Account{}

	err := db.Raw(`SELECT a.id, a.owner_id, a.name, a.is_active, a.created_at, a.created_by, a.updated_at, a.updated_by, a.deleted_at
		FROM accounts a JOIN base_box b ON a.id = b.account_id JOIN answer_box_questions q ON q.box_id = b.id JOIN answer_box_entries e ON e.question_id = q.question_id
		WHERE e.entry_id = ?`, entryId).Scan(&account).Error

	return account, err
}

// GetAccountByAnswerBoxQuestion returns the account associated with the box associated with the question
func GetAccountByAnswerBoxQuestion(questionId uuid.UUID) (accountmodel.Account, error) {
	db := flight.Context(nil, nil).DB

	account := accountmodel.Account{}

	err := db.Raw(`SELECT a.id, a.owner_id, a.name, a.is_active, a.created_at, a.created_by, a.updated_at, a.updated_by, a.deleted_at
		FROM accounts a JOIN base_box b ON a.id = b.account_id JOIN answer_box_questions q ON q.box_id = b.id 
		WHERE q.question_id = ?`, questionId).Scan(&account).Error

	return account, err
}

// GetAccountByQuestionBoxEntry returns the account associated with the box associated with the question associated with the entry
func GetAccountByQuestionBoxEntry(entryId uuid.UUID) (accountmodel.Account, error) {
	db := flight.Context(nil, nil).DB

	account := accountmodel.Account{}

	err := db.Raw(`SELECT a.id, a.owner_id, a.name, a.is_active, a.created_at, a.created_by, a.updated_at, a.updated_by, a.deleted_at
		FROM accounts a JOIN base_box b ON a.id = b.account_id JOIN question_box_entries q ON q.box_id = b.id
		WHERE q.entry_id = ?`, entryId).Scan(&account).Error

	return account, err
}

// GetAccountByBoxId returns the account associated with the box
func GetAccountByBoxId(boxId uuid.UUID) (accountmodel.Account, error) {
	db := flight.Context(nil, nil).DB

	account := accountmodel.Account{}

	err := db.Raw(`SELECT a.id, a.owner_id, a.name, a.is_active, a.created_at, a.created_by, a.updated_at, a.updated_by, a.deleted_at
		FROM accounts a JOIN base_box b ON a.id = b.account_id
		WHERE b.id = ?`, boxId).Scan(&account).Error

	return account, err
}

// GetAccountByCode returns the account associated with the box
func GetAccountByCode(code string) (accountmodel.Account, error) {
	db := flight.Context(nil, nil).DB

	account := accountmodel.Account{}

	err := db.Raw(`SELECT a.id, a.owner_id, a.name, a.is_active, a.created_at, a.created_by, a.updated_at, a.updated_by, a.deleted_at
		FROM accounts a JOIN base_box b ON a.id = b.account_id
		WHERE b.code = ?`, code).Scan(&account).Error

	return account, err
}

// UpdateAccount is for updating an account
func UpdateAccount(account accountmodel.Account) (accountmodel.Account, error) {
	db := flight.Context(nil, nil).DB

	err := db.Exec(`UPDATE accounts SET name = ? WHERE id = ?`, account.Name, account.Id).Error

	return account, err
}

// UpdateSubscriptionForCancellation updates the account's price plan
func UpdateSubscriptionForCancellation(userStripe userstripemodel.UserStripe) error {
	db := flight.Context(nil, nil).DB
	var err error

	// Wrapping box creation in a transaction. We wouldn't want the price plan to be updated without updating the customer mapping and making them get a new token
	tx := db.Begin()

	//update account priceplan
	planId, _ := uuid.FromString(priceplanconstants.BASIC)
	if err := UpdateAccountPricePlanTx(userStripe.AccountId, planId, nil, tx); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Exec("UPDATE users_stripe SET updated_at = ? WHERE user_id = ?", time.Now(), userStripe.UserId).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return err
}

//GetActiveNonBasicAccounts gets all active accounts which are not on the BASIC plan
func GetActiveNonBasicAccounts() ([]accountmodel.Account, error) {
	db := flight.Context(nil, nil).DB

	accounts := []accountmodel.Account{}

	err := db.Raw(`SELECT a.id, a.owner_id, a.name, a.is_active, a.created_at, a.created_by, a.updated_at, a.updated_by, a.deleted_at 
		FROM account_price_plans app JOIN accounts a ON app.account_id = a.id
		WHERE app.is_active = 1 AND app.plan_id != ?;`, priceplanconstants.BASIC).Scan(&accounts).Error

	return accounts, err
}

// CreateAccountInvite creates an account invite
func CreateAccountInvite(accountInvitation accountinvitationmodel.AccountInvitation, invitationCode string, accountId uuid.UUID, createdBy uuid.UUID) error {
	db := flight.Context(nil, nil).DB

	newId, _ := uuid.NewV4()
	err := db.Exec(`INSERT INTO account_invitations (id, account_id, role_id, email, invitation_code, is_active, created_at, created_by) VALUES (?, ?, ?, ?, ?, 1, CURRENT_TIMESTAMP, ?)`,
		newId, accountId, accountInvitation.RoleId, accountInvitation.Email, invitationCode, createdBy).Error

	return err
}

// GetActiveAccountInvites gets all active invites
func GetActiveAccountInvites(accountId uuid.UUID) ([]accountinvitationmodel.AccountInvitation, error) {
	db := flight.Context(nil, nil).DB
	var accountInvitations []accountinvitationmodel.AccountInvitation

	err := db.Raw(`SELECT ai.id, account_id, role_id, r.name as role_name, email, invitation_code, is_active, ai.created_at, ai.created_by
		FROM account_invitations ai JOIN roles r ON ai.role_id = r.id
		WHERE is_active = 1 AND account_id = ?`, accountId).Scan(&accountInvitations).Error

	return accountInvitations, err
}

// GetInvitation retrieves a pending invitation's data
func GetInvitation(invitationCode string) (accountinvitationmodel.AccountInvitation, error) {
	db := flight.Context(nil, nil).DB
	var accountInvitation accountinvitationmodel.AccountInvitation

	err := db.Raw(`SELECT ai.id, ai.account_id, ai.role_id, ai.email, ai.invitation_code, ai.is_active, ai.created_at, ai.created_by, a.name as account_name
		FROM account_invitations ai JOIN accounts a ON ai.account_id = a.id
		WHERE invitation_code = ?`, invitationCode).Scan(&accountInvitation).Error

	return accountInvitation, err
}

// GetInvitationByAccountAndEmail retrieves a pending invitation's data
func GetInvitationByAccountAndEmail(accountId uuid.UUID, email string) (accountinvitationmodel.AccountInvitation, error) {
	db := flight.Context(nil, nil).DB
	var accountInvitation accountinvitationmodel.AccountInvitation

	err := db.Raw(`
		SELECT ai.id, ai.account_id, ai.role_id, ai.email, ai.invitation_code, ai.is_active, ai.created_at, ai.created_by, a.name as account_name
		FROM account_invitations ai JOIN accounts a ON ai.account_id = a.id
		WHERE ai.is_active = 1
		AND ai.account_id = ?
		AND ai.email = ?`, accountId, email).Scan(&accountInvitation).Error

	return accountInvitation, err
}

// GetInvitationForUpdate retrieves a pending invitation's data using the account ID and email address
func GetInvitationForUpdate(accountInvitation accountinvitationmodel.AccountInvitation) (accountinvitationmodel.AccountInvitation, error) {
	db := flight.Context(nil, nil).DB
	var fullInvitation accountinvitationmodel.AccountInvitation

	err := db.Raw(`SELECT ai.id, ai.account_id, ai.role_id, ai.email, ai.invitation_code, ai.is_active, ai.created_at, ai.created_by, a.name as account_name
		FROM account_invitations ai JOIN accounts a ON ai.account_id = a.id
		WHERE ai.account_id = ? AND ai.email = ?  
		ORDER BY ai.created_at DESC
		LIMIT 1`, accountInvitation.AccountId, accountInvitation.Email).Scan(&fullInvitation).Error

	return fullInvitation, err
}

// CancelInvitation cancels an invitation to an account
func CancelInvitation(invitationCode string, updatedBy uuid.UUID) error {
	db := flight.Context(nil, nil).DB

	err := db.Exec(`UPDATE account_invitations SET is_active = 0, updated_at = ?, updated_by = ? WHERE invitation_code = ?`, time.Now().UTC(), updatedBy, invitationCode).Error

	return err
}
