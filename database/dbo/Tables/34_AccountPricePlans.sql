CREATE TABLE IF NOT EXISTS `justasking`.`account_price_plans` (
  `id`              CHAR(36)        NOT NULL,
  `account_id`      CHAR(36)        NOT NULL,
  `plan_id`         CHAR(36)        NOT NULL,
  `is_active`       TINYINT         NOT NULL DEFAULT 1,
  `period_end` 		TIMESTAMP 		NULL,
  `created_at`      TIMESTAMP       NOT NULL    DEFAULT CURRENT_TIMESTAMP,
  `updated_at`      TIMESTAMP       NULL,
  `deleted_at`      TIMESTAMP       NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_account_price_plan_account_id`
    FOREIGN KEY (`account_id`)
    REFERENCES `justasking`.`accounts` (`id`)
    ON DELETE RESTRICT
    ON UPDATE RESTRICT,
  CONSTRAINT `fk_account_price_plan_plan_id`
    FOREIGN KEY (`plan_id`)
    REFERENCES `justasking`.`price_plans` (`id`)
    ON DELETE RESTRICT
    ON UPDATE RESTRICT);
        