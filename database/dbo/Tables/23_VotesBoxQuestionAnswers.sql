CREATE TABLE IF NOT EXISTS `justasking`.`votes_box_question_answers` (
  `answer_id` CHAR(36) NOT NULL,
  `question_id` CHAR(36) NOT NULL,
  `answer` VARCHAR(250) NOT NULL,
  `sort_order` INT NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(50) NOT NULL,
  `updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_by` varchar(50) NULL,
  `deleted_at` TIMESTAMP NULL DEFAULT NULL,
  PRIMARY KEY (`answer_id`),
  CONSTRAINT `fk_votes_box_question_answer_id`
    FOREIGN KEY (`question_id`)
    REFERENCES `justasking`.`votes_box_questions` (`question_id`)
    ON DELETE RESTRICT
    ON UPDATE RESTRICT);
