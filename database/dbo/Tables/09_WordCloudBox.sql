CREATE TABLE IF NOT EXISTS `justasking`.`word_cloud_box` (
  `box_id` CHAR(36) NOT NULL,
  `header` VARCHAR(256) NOT NULL,
  `default_word` VARCHAR(50) NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(50) NOT NULL,
  `updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_by` varchar(50) NULL,
  `deleted_at` TIMESTAMP NULL DEFAULT NULL,
  PRIMARY KEY (`box_id`),
  CONSTRAINT `fk_word_cloud_box_id`
    FOREIGN KEY (`box_id`)
    REFERENCES `justasking`.`base_box` (`id`)
    ON DELETE RESTRICT
    ON UPDATE RESTRICT);
