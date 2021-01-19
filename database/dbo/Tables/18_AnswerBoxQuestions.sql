CREATE TABLE IF NOT EXISTS `justasking`.`answer_box_questions` (
  `question_id` CHAR(36) NOT NULL,
  `box_id` CHAR(36) NOT NULL,
  `question` VARCHAR(128) NOT NULL,
  `is_active` TINYINT NOT NULL DEFAULT 0,
  `sort_order` INT NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(50) NOT NULL,
  `updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_by` varchar(50) NULL,
  `deleted_at` TIMESTAMP NULL DEFAULT NULL,
  PRIMARY KEY (`question_id`),
  CONSTRAINT `fk_answer_box_question_id`
    FOREIGN KEY (`box_id`)
    REFERENCES `justasking`.`base_box` (`id`)
    ON DELETE RESTRICT
    ON UPDATE RESTRICT);
