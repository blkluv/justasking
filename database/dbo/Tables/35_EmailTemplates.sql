CREATE TABLE IF NOT EXISTS `justasking`.`email_templates` ( 
  `id` INT NOT NULL,
  `name` VARCHAR(150) NOT NULL,
  `is_active` TINYINT NOT NULL DEFAULT 0,
  `to` VARCHAR(1500) NOT NULL,
  `cc` VARCHAR(500) NOT NULL,
  `bcc` VARCHAR(500) NOT NULL,
  `from` VARCHAR(100) NOT NULL,
  `subject` VARCHAR(500) NOT NULL,
  `body` TEXT NOT NULL,
  `created_at` TIMESTAMP NOT NULL,
  `updated_at` TIMESTAMP NULL,
  `deleted_at` TIMESTAMP NULL,
  PRIMARY KEY (`id`));
  
  