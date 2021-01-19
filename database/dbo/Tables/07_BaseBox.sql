CREATE TABLE IF NOT EXISTS `justasking`.`base_box` (
  `id` CHAR(36) NOT NULL,
  `code` VARCHAR(50) NULL,
  `original_code` VARCHAR(50) NULL,
  `account_id` CHAR(36) NOT NULL,
  `box_type` INT NOT NULL,
  `theme_id` INT NOT NULL,
  `is_live` TINYINT NOT NULL DEFAULT 0,
  `entry_page_enabled` TINYINT NOT NULL DEFAULT 1,
  `presentation_page_enabled` TINYINT NOT NULL DEFAULT 1,
  `login_required` TINYINT NOT NULL DEFAULT 0,
  `sms_enabled` TINYINT NOT NULL DEFAULT 1,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(50) NOT NULL,
  `updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_by` varchar(50) NULL,
  `deleted_at` TIMESTAMP NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `code_UNIQUE` (`code` ASC),
  CONSTRAINT `fk_base_box_account_id`
    FOREIGN KEY (`account_id`)
    REFERENCES `justasking`.`accounts` (`id`)
    ON DELETE RESTRICT
    ON UPDATE RESTRICT);