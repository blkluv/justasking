CREATE TABLE IF NOT EXISTS `justasking`.`password_reset_requests` (
  `id` CHAR(36) NOT NULL,
  `user_id` CHAR(36) NOT NULL,
  `reset_code` VARCHAR(256)  NOT NULL,
  `is_active` TINYINT NOT NULL DEFAULT 1,
  `expires_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` TIMESTAMP NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  INDEX `password_reset_request_code` (`reset_code`),
  CONSTRAINT `fk_password_reset_id`
    FOREIGN KEY (`user_id`)
    REFERENCES `justasking`.`users` (`id`)
    ON DELETE RESTRICT
    ON UPDATE RESTRICT);
  