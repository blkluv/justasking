CREATE TABLE IF NOT EXISTS `justasking`.`feature_requests` (
  `feature_request_id` CHAR(36) NOT NULL,
  `feature_request` VARCHAR(250) NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(50) NOT NULL,
  CONSTRAINT `fk_feature_request_id`
    FOREIGN KEY (`created_by`)
    REFERENCES `justasking`.`users` (`id`)
    ON DELETE RESTRICT
    ON UPDATE RESTRICT);
