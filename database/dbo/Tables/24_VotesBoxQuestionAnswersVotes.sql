CREATE TABLE IF NOT EXISTS `justasking`.`votes_box_question_answers_votes` (
  `answer_id` CHAR(36) NOT NULL,
  `question_id` CHAR(36) NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(50) NOT NULL, 
  `updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_by` varchar(50) NULL,
  `deleted_at` TIMESTAMP NULL DEFAULT NULL,
  INDEX `fk_votes_box_answer_vote_id` (`answer_id`),
  CONSTRAINT `fk_votes_box_answer_vote_id`
    FOREIGN KEY (`answer_id`)
    REFERENCES `justasking`.`votes_box_question_answers` (`answer_id`)
    ON DELETE RESTRICT 
    ON UPDATE RESTRICT);