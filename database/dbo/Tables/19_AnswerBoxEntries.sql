CREATE TABLE IF NOT EXISTS `justasking`.`answer_box_entries` (
  `entry_id` CHAR(36) NOT NULL,
  `question_id` CHAR(36) NOT NULL,
  `entry` VARCHAR(2000) NOT NULL,
  `is_hidden` TINYINT NOT NULL DEFAULT 0,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(50) NOT NULL,
  `updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_by` varchar(50) NULL,
  `deleted_at` TIMESTAMP NULL DEFAULT NULL,
  PRIMARY KEY (`entry_id`),
  CONSTRAINT `fk_answer_box_entry_id`
    FOREIGN KEY (`question_id`)
    REFERENCES `justasking`.`answer_box_questions` (`question_id`)
    ON DELETE RESTRICT
    ON UPDATE RESTRICT);
