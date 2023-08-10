CREATE TABLE IF NOT EXISTS `justasking`.`roles` (
  `id`          CHAR(36)    NOT NULL,
  `name`        VARCHAR(50) NOT NULL,
  `created_at`  TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at`  TIMESTAMP   NULL,
  `deleted_at`  TIMESTAMP   NULL,
  PRIMARY KEY (`id`));
  