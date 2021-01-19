CREATE TABLE IF NOT EXISTS `justasking`.`permissions` (
  `id`          CHAR(36)    NOT NULL,
  `name`        VARCHAR(50) NOT NULL,
  `created_at`  TIMESTAMP   NOT NULL,
  `updated_at`  TIMESTAMP   NULL,
  `deleted_at`  TIMESTAMP   NULL,
  PRIMARY KEY (`id`));
  