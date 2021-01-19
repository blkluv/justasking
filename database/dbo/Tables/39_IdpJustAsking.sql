CREATE TABLE IF NOT EXISTS `justasking`.`idp_justasking` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `email` VARCHAR(128) NOT NULL UNIQUE,
  `password` VARCHAR(256)  NOT NULL,
  `sub` CHAR(36) NOT NULL,
  `name` VARCHAR(256)  NULL,
  `phone_number` VARCHAR(25)  NULL,
  `image_url` VARCHAR(512)  NULL,
  `given_name` VARCHAR(128)  NULL,
  `family_name` VARCHAR(128)  NULL,
  `created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` TIMESTAMP NULL DEFAULT NULL,
  PRIMARY KEY (`id`));
  