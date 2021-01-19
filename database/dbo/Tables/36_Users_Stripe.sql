CREATE TABLE IF NOT EXISTS `justasking`.`users_stripe` (
  `user_id`                 CHAR(36)        NOT NULL,
  `account_id`              CHAR(36)        NOT NULL,
  `stripe_user_id`          varchar(50)     NOT NULL,
  `stripe_payment_token`    varchar(50)     NULL,
  `credit_card_last_four`   varchar(4)      NULL,
  `last_payment`            TIMESTAMP       NULL,
  `created_at`              TIMESTAMP       NOT NULL    DEFAULT CURRENT_TIMESTAMP,
  `updated_at`              TIMESTAMP       NULL,
  `deleted_at`              TIMESTAMP       NULL,
  CONSTRAINT `fk_users_stripe_user_id`
    FOREIGN KEY (`user_id`)
    REFERENCES `justasking`.`users` (`id`),
  CONSTRAINT `fk_users_stripe_account_id`
    FOREIGN KEY (`account_id`)
    REFERENCES `justasking`.`accounts` (`id`));
