CREATE TABLE IF NOT EXISTS `justasking`.`price_plan_features` (
  `id`              CHAR(36)        NOT NULL,
  `plan_id`         CHAR(36)        NOT NULL,
  `feature_id`      CHAR(36)        NOT NULL,
  `feature_value`   VARCHAR(50)     NOT NULL,
  `created_at`      TIMESTAMP       NOT NULL    DEFAULT CURRENT_TIMESTAMP,
  `updated_at`      TIMESTAMP       NULL,
  `deleted_at`      TIMESTAMP       NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_plan_features_plan_id`
    FOREIGN KEY (`plan_id`)
    REFERENCES `justasking`.`price_plans` (`id`)
    ON DELETE RESTRICT
    ON UPDATE RESTRICT,
  CONSTRAINT `fk_plan_features_feature_id`
    FOREIGN KEY (`feature_id`)
    REFERENCES `justasking`.`features` (`id`)
    ON DELETE RESTRICT
    ON UPDATE RESTRICT);
  