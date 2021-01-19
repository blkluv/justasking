CREATE TABLE IF NOT EXISTS `justasking`.`accounts` (
  `id`                  CHAR(36)        NOT NULL,
  `owner_id`            CHAR(36)        NOT NULL,
  `name`                VARCHAR(50)    	NULL,
  `is_active`           TINYINT         NOT NULL DEFAULT 1,
  `created_at`          TIMESTAMP       NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by`          varchar(50)     NOT NULL,
  `updated_at`          TIMESTAMP       NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_by`          varchar(50)     NULL,
  `deleted_at`          TIMESTAMP       NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_accounts_owner_id`
    FOREIGN KEY (`owner_id`)
    REFERENCES `justasking`.`users` (`id`)
    ON DELETE RESTRICT
    ON UPDATE RESTRICT);
