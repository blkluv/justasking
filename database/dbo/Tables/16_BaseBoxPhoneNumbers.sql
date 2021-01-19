CREATE TABLE IF NOT EXISTS `justasking`.`base_box_phone_numbers` (
  `id`               CHAR(36)    NOT NULL,
  `base_box_id`      CHAR(36)    NOT NULL,
  `phone_number_id`  CHAR(36)    NOT NULL,
  `is_active`        TINYINT     NOT NULL DEFAULT 0,
  `created_at`       TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by`       varchar(50) NOT NULL,
  `updated_at`       TIMESTAMP   NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_by`       varchar(50) NULL,
  `deleted_at`       TIMESTAMP   NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_base_box_id`
  FOREIGN KEY (`base_box_id`)
  REFERENCES `justasking`.`base_box` (`id`)
  ON DELETE RESTRICT
  ON UPDATE RESTRICT,
  CONSTRAINT `fk_phone_number_id`
  FOREIGN KEY (`phone_number_id`)
  REFERENCES `justasking`.`phone_numbers` (`id`)
  ON DELETE RESTRICT
  ON UPDATE RESTRICT);