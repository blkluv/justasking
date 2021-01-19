CREATE TABLE IF NOT EXISTS `justasking`.`question_box_entries` (
  `entry_id` CHAR(36) NOT NULL,
  `box_id` CHAR(36) NOT NULL,
  `question` VARCHAR(128) NOT NULL,
  `is_hidden` TINYINT NOT NULL DEFAULT 0,
  `is_favorite` TINYINT NOT NULL DEFAULT 0,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(50) NOT NULL,
  `updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_by` varchar(50) NULL,
  `deleted_at` TIMESTAMP NULL DEFAULT NULL,
  PRIMARY KEY (`entry_id`),
  CONSTRAINT `fk_question_box_entry_id`
    FOREIGN KEY (`box_id`)
    REFERENCES `justasking`.`base_box` (`id`)
    ON DELETE RESTRICT
    ON UPDATE RESTRICT);
