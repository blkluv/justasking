package userstriperepo

import (
	"justasking/GO/core/model/userstripe"
	"justasking/GO/core/startup/flight"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// InsertUserStripeMapping creates a mapping between a justasking user and stripe customer
func InsertUserStripeMapping(userId uuid.UUID, accountId uuid.UUID, stripeCustomerId string, tx *gorm.DB) error {
	err := tx.Exec(`INSERT INTO users_stripe (user_id, stripe_user_id, account_id)
					SELECT ? as user_id, ? as stripe_user_id, ? as account_id
					WHERE NOT EXISTS (SELECT user_id FROM users_stripe WHERE user_id = ? AND account_id = ?)`,
		userId, stripeCustomerId, accountId, userId, accountId).Error

	return err
}

// GetUserStripeData gets a mapping between a justasking user and stripe customer
func GetUserStripeData(userId uuid.UUID, accountId uuid.UUID) (userstripemodel.UserStripe, error) {
	db := flight.Context(nil, nil).DB

	userStripe := userstripemodel.UserStripe{}

	err := db.Raw(`SELECT user_id, account_id, stripe_user_id, stripe_payment_token, account_id, credit_card_last_four, last_payment, last_payment, created_at, updated_at, deleted_at FROM users_stripe WHERE user_id = ? AND account_id = ?`,
		userId, accountId).Scan(&userStripe).Error

	return userStripe, err
}

// GetUserStripeDataByUserId is used when we're only interested in retrieving the stripe_user_id, because it should be the same across all users_stripe records for a user
func GetUserStripeDataByUserId(userId uuid.UUID) (userstripemodel.UserStripe, error) {
	db := flight.Context(nil, nil).DB

	userStripe := userstripemodel.UserStripe{}

	err := db.Where(`user_id = ?`, userId).First(&userStripe).Error

	return userStripe, err
}

// GetUserStripeDataByAccountId gets a mapping between a justasking user and stripe customer
func GetUserStripeDataByAccountId(accountId uuid.UUID) ([]userstripemodel.UserStripe, error) {
	db := flight.Context(nil, nil).DB

	userStripe := []userstripemodel.UserStripe{}

	err := db.Raw(`SELECT user_id, account_id, stripe_user_id, stripe_payment_token, account_id, credit_card_last_four, last_payment, last_payment, created_at, updated_at, deleted_at FROM users_stripe WHERE account_id = ?`,
		accountId).Scan(&userStripe).Error

	return userStripe, err
}

// UpdateUserStripeDataTx updates the user/stripe customer mapping
func UpdateUserStripeDataTx(userStripe userstripemodel.UserStripe, tx *gorm.DB) error {
	err := tx.Exec("UPDATE users_stripe SET stripe_user_id = ?, stripe_payment_token = ?, credit_card_last_four = ?, last_payment = ?, updated_at = ? WHERE user_id = ? AND account_id = ?",
		userStripe.StripeUserId, userStripe.StripePaymentToken, userStripe.CreditCardLastFour, time.Now(), time.Now(), userStripe.UserId, userStripe.AccountId).Error

	return err
}

// UpdateUserStripeData updates the user/stripe customer mapping
func UpdateUserStripeData(userStripe userstripemodel.UserStripe) error {
	db := flight.Context(nil, nil).DB

	err := db.Exec("UPDATE users_stripe SET stripe_user_id = ?, stripe_payment_token = ?, credit_card_last_four = ?,last_payment = ?, updated_at = ? WHERE user_id = ? AND account_id = ?",
		userStripe.StripeUserId, userStripe.StripePaymentToken, userStripe.CreditCardLastFour, userStripe.LastPayment, time.Now(), userStripe.UserId, userStripe.AccountId).Error

	return err
}
