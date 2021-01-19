CREATE TABLE IF NOT EXISTS `justasking`.`support_issues` (
  `issue_id` CHAR(36) NOT NULL,
  `issue` VARCHAR(250) NOT NULL,
  `user_agent` VARCHAR(200) NOT NULL,
  `resolved` TINYINT NOT NULL DEFAULT 0,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(50) NOT NULL,
  `updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_by` varchar(50) NULL,
  `deleted_at` TIMESTAMP NULL DEFAULT NULL,
  CONSTRAINT `fk_support_issue_id`
    FOREIGN KEY (`created_by`)
    REFERENCES `justasking`.`users` (`id`)
    ON DELETE RESTRICT
    ON UPDATE RESTRICT);
