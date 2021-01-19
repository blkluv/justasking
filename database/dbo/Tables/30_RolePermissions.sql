CREATE TABLE IF NOT EXISTS `justasking`.`role_permissions` (
  `role_id`             CHAR(36)    NOT NULL,
  `permission_id`       CHAR(36)    NOT NULL,
  `permission_value`    TINYINT     NOT NULL,
  `created_at`          TIMESTAMP   NOT NULL,
  `updated_at`          TIMESTAMP   NULL,
  `deleted_at`          TIMESTAMP   NULL,
  CONSTRAINT `fk_role_permissions_role_id`
    FOREIGN KEY (`role_id`)
    REFERENCES `justasking`.`roles` (`id`)
    ON DELETE RESTRICT
    ON UPDATE RESTRICT,
  CONSTRAINT `fk_role_permissions_permission_id`
    FOREIGN KEY (`permission_id`)
    REFERENCES `justasking`.`permissions` (`id`)
    ON DELETE RESTRICT
    ON UPDATE RESTRICT);
  