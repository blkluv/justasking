package striperepo

import (
	"justasking/GO/core/model/priceplan"
	"justasking/GO/core/model/userstripe"
	"justasking/GO/core/repo/account"
	"justasking/GO/core/repo/userstripe"
	"justasking/GO/core/startup/flight"
	"time"

	"github.com/jinzhu/gorm"
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"github.com/stripe/stripe-go/customer"
)

// CreateStripeCustomer creates a customer in stripe
func CreateStripeCustomer(userEmail string, stripeKey string, tx *gorm.DB) (*stripe.Customer, error) {
	stripe.Key = stripeKey
	params := &stripe.CustomerParams{
		Email: userEmail,
	}
	customer, err := customer.New(params)

	return customer, err
}

// UpdateSubscription updates a subscription in stripe and our system
func UpdateSubscription(planData priceplanmodel.PricePlan, stripeKey string, stripeData userstripemodel.UserStripe, endDate *time.Time, customPlan bool) error {
	db := flight.Context(nil, nil).DB
	var err error

	stripe.Key = stripeKey

	chargeParams := &stripe.ChargeParams{
		Amount:   uint64(planData.Price) * 100, //stripe takes the amount in cents
		Currency: "usd",
		Customer: stripeData.StripeUserId,
	}
	_, err = charge.New(chargeParams)

	if err == nil {

		// Wrapping update in a transaction. We wouldn't want our data and the stripe data to be out of sync
		tx := db.Begin()

		//update user/stripe mapping
		if err := userstriperepo.UpdateUserStripeDataTx(stripeData, tx); err != nil {
			tx.Rollback()
			return err
		}

		//update account/priceplan mapping
		if err := accountrepo.UpdateAccountPricePlanTx(stripeData.AccountId, planData.Id, endDate, tx); err != nil {
			tx.Rollback()
			return err
		}

		//update token version for all users on this account
		if err := tx.Exec(`UPDATE account_users SET token_version = UUID(), updated_at = ? WHERE account_id = ?`, time.Now().UTC(), stripeData.AccountId).Error; err != nil {
			tx.Rollback()
			return err
		}

		if customPlan {
			//update custom plan mapping
			if err := tx.Exec(`UPDATE custom_plan_licenses SET is_active = 0, updated_at = ?, updated_by = ? WHERE user_id = ? AND account_id = ?`, time.Now().UTC(), stripeData.UserId, stripeData.UserId, stripeData.AccountId).Error; err != nil {
				tx.Rollback()
				return err
			}
		}

		tx.Commit()
	}

	return err
}

// UpdateCreditCard adds or updates a credit card for a stripe customer
func UpdateCreditCard(stripeData userstripemodel.UserStripe, cardString string, lastFour string, stripeKey string) (*stripe.Customer, error) {
	db := flight.Context(nil, nil).DB
	stripe.Key = stripeKey

	// Wrapping update in a transaction. We wouldn't want our data and the stripe data to be out of sync
	tx := db.Begin()

	stripeData.StripePaymentToken = cardString
	stripeData.CreditCardLastFour = lastFour
	//update user/stripe mapping
	if err := userstriperepo.UpdateUserStripeDataTx(stripeData, tx); err != nil {
		tx.Rollback()
		return nil, err
	}
	tokenParams := &stripe.SourceParams{
		Token: cardString,
	}
	customerParams := &stripe.CustomerParams{
		Source: tokenParams,
	}
	customer, err := customer.Update(
		stripeData.StripeUserId,
		customerParams,
	)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return customer, err
}
