CREATE TABLE IF NOT EXISTS `justasking`.`custom_plan_licenses` (
  `id`                  CHAR(36)        NOT NULL,
  `account_id`          CHAR(36)        NOT NULL,
  `user_id`             CHAR(36)        NOT NULL,
  `plan_id`             CHAR(36)        NOT NULL,
  `license_code`        VARCHAR(256)    NOT NULL,
  `is_active`           TINYINT         NOT NULL DEFAULT 1,
  `created_at`          TIMESTAMP       NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by`          varchar(50)     NOT NULL,
  `updated_at`          TIMESTAMP       NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_by`          varchar(50)     NULL,
  `deleted_at`          TIMESTAMP       NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  INDEX `custom_plan_licenses_license_code` (`license_code`),
  CONSTRAINT `fk_custom_plan_licenses_account_id`
    FOREIGN KEY (`account_id`)
    REFERENCES `justasking`.`accounts` (`id`)
    ON DELETE RESTRICT
    ON UPDATE RESTRICT,
    CONSTRAINT `fk_custom_plan_licenses_user_id`
    FOREIGN KEY (`user_id`)
    REFERENCES `justasking`.`users` (`id`)
    ON DELETE RESTRICT
    ON UPDATE RESTRICT,
    CONSTRAINT `fk_custom_plan_licenses_plan_id`
    FOREIGN KEY (`plan_id`)
    REFERENCES `justasking`.`price_plans` (`id`)
    ON DELETE RESTRICT
    ON UPDATE RESTRICT);
  