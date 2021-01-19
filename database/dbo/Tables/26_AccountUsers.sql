CREATE TABLE IF NOT EXISTS `justasking`.`account_users` (
  `account_id`          CHAR(36)        NOT NULL,
  `user_id`             CHAR(36)        NOT NULL,
  `role_id`             CHAR(36)        NOT NULL,
  `is_active`           TINYINT         NOT NULL DEFAULT 1,
  `current_account` 	TINYINT		 	NOT NULL DEFAULT 0,
  `token_version` 		char(36) 		NOT NULL,
  `created_at`          TIMESTAMP       NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by`          varchar(50)     NOT NULL,
  `updated_at`          TIMESTAMP       NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_by`          varchar(50)     NULL,
  `deleted_at`          TIMESTAMP       NULL DEFAULT NULL,
  CONSTRAINT `fk_account_users_account_id`
    FOREIGN KEY (`account_id`)
    REFERENCES `justasking`.`accounts` (`id`)
    ON DELETE RESTRICT
    ON UPDATE RESTRICT,
  CONSTRAINT `fk_account_users_user_id`
    FOREIGN KEY (`user_id`)
    REFERENCES `justasking`.`users` (`id`)
    ON DELETE RESTRICT
    ON UPDATE RESTRICT);
