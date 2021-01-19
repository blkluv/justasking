CREATE TABLE IF NOT EXISTS `justasking`.`account_invitations` (
  `id`                  CHAR(36)        NOT NULL,
  `account_id`          CHAR(36)        NOT NULL,
  `role_id`             CHAR(36)        NOT NULL,
  `email`               VARCHAR(128)    NOT NULL,
  `invitation_code`     VARCHAR(256)    NOT NULL,
  `is_active`           TINYINT         NOT NULL DEFAULT 1,
  `created_at`          TIMESTAMP       NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by`          varchar(50)     NOT NULL,
  `updated_at`          TIMESTAMP       NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_by`          varchar(50)     NULL,
  `deleted_at`          TIMESTAMP       NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  INDEX `account_invitations_request_code` (`invitation_code`),
  CONSTRAINT `fk_account_invitations_account_id`
    FOREIGN KEY (`account_id`)
    REFERENCES `justasking`.`accounts` (`id`)
    ON DELETE RESTRICT
    ON UPDATE RESTRICT);
  