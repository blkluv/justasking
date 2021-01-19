package authenticationrepo

import (
	"justasking/GO/core/model/account"
	"justasking/GO/core/model/accountuser"
	"justasking/GO/core/model/authentication"
	"justasking/GO/core/model/idpjustasking"
	"justasking/GO/core/model/idpmapping"
	"justasking/GO/core/model/user"
	"justasking/GO/core/repo/account"
	"justasking/GO/core/repo/accountuser"
	"justasking/GO/core/repo/stripe"
	"justasking/GO/core/repo/userstripe"
	"justasking/GO/core/startup/flight"
	"time"
)

// GetGoogleUserBySub gets a user from the idp_google table
func GetGoogleUserBySub(idpID string) (authenticationmodel.IdpGoogle, error) {
	db := flight.Context(nil, nil).DB

	var googleUser authenticationmodel.IdpGoogle
	err := db.Where("sub = ?", idpID).First(&googleUser).Error

	return googleUser, err
}

// CreateGoogleUser adds a user to the idp_google table
func CreateGoogleUser(idpGoogle authenticationmodel.IdpGoogle, justAskingUser usermodel.User, idpMapping idpmappingmodel.IdpMapping, account accountmodel.Account, accountUser accountusermodel.AccountUser, stripeKey string) (*usermodel.User, error) {
	db := flight.Context(nil, nil).DB

	// Wrapping user creation in a transaction. We wouldn't want the Google user to be created without the JustAsking user
	tx := db.Begin()

	if err := tx.Create(&idpGoogle).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	justAskingUser.LastLoggedInAt = time.Now().UTC()
	if err := tx.Create(&justAskingUser).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	idpMapping.UserId = justAskingUser.ID
	if err := tx.Create(&idpMapping).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := accountrepo.InsertAccount(account, tx); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := accountuserrepo.InsertAccountUser(accountUser, tx); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := accountrepo.InsertDefaultAccountPricePlan(account.Id, tx); err != nil {
		tx.Rollback()
		return nil, err
	}

	stripeCustomer, err := striperepo.CreateStripeCustomer(justAskingUser.Email, stripeKey, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := userstriperepo.InsertUserStripeMapping(justAskingUser.ID, account.Id, stripeCustomer.ID, tx); err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return &justAskingUser, nil
}

// CreateIdpJustAskingUser creates a record in the idp_justasking table
func CreateIdpJustAskingUser(idpJustAskingUser idpjustaskingmodel.IdpJustAsking, justAskingUser usermodel.User, idpMapping idpmappingmodel.IdpMapping, account accountmodel.Account, accountUser accountusermodel.AccountUser, stripeKey string) (*usermodel.User, error) {
	db := flight.Context(nil, nil).DB

	// Wrapping user creation in a transaction. We wouldn't want the idpJustAsking user to be created without the regular JustAsking user
	tx := db.Begin()

	if err := tx.Create(&idpJustAskingUser).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	justAskingUser.LastLoggedInAt = time.Now().UTC()
	if err := tx.Create(&justAskingUser).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	idpMapping.UserId = justAskingUser.ID
	if err := tx.Create(&idpMapping).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := accountrepo.InsertAccount(account, tx); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := accountuserrepo.InsertAccountUser(accountUser, tx); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := accountrepo.InsertDefaultAccountPricePlan(account.Id, tx); err != nil {
		tx.Rollback()
		return nil, err
	}

	stripeCustomer, err := striperepo.CreateStripeCustomer(justAskingUser.Email, stripeKey, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := userstriperepo.InsertUserStripeMapping(justAskingUser.ID, account.Id, stripeCustomer.ID, tx); err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return &justAskingUser, nil
}
