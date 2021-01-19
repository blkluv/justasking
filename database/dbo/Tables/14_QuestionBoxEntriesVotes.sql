CREATE TABLE IF NOT EXISTS `justasking`.`question_box_entries_votes` (
  `entry_id` CHAR(36) NOT NULL,
  `vote_type` varchar(8) NOT NULL,
  `vote_value` INT NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(50) NOT NULL,
  `updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_by` varchar(50) NULL,
  `deleted_at` TIMESTAMP NULL DEFAULT NULL,
  INDEX `fk_question_box_entry_vote_id` (`entry_id`),
  CONSTRAINT `fk_question_box_entry_vote_id`
    FOREIGN KEY (`entry_id`)
    REFERENCES `justasking`.`question_box_entries` (`entry_id`)
    ON DELETE RESTRICT
    ON UPDATE RESTRICT);