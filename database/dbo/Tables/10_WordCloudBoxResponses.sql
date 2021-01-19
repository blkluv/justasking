CREATE TABLE IF NOT EXISTS `justasking`.`word_cloud_box_responses` (
  `box_id` CHAR(36) NOT NULL,
  `response` VARCHAR(50) NOT NULL,
  `is_hidden` TINYINT NOT NULL DEFAULT 0,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(50) NOT NULL,
  `updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_by` varchar(50) NULL,
  `deleted_at` TIMESTAMP NULL DEFAULT NULL,
  CONSTRAINT `fk_word_cloud_box_response_id`
    FOREIGN KEY (`box_id`)
    REFERENCES `justasking`.`base_box` (`id`)
    ON DELETE RESTRICT
    ON UPDATE RESTRICT);
