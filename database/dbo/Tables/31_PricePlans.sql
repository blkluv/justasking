CREATE TABLE IF NOT EXISTS `justasking`.`price_plans` (
  `id`                          CHAR(36)        NOT NULL,
  `name`                        VARCHAR(50)     NOT NULL,
  `description`                 VARCHAR(250)    NOT NULL,
  `display_name`                VARCHAR(50)     NOT NULL,
  `price`                       INT             NOT NULL,
  `price_description`           VARCHAR(50)     NOT NULL,
  `image_path`                  VARCHAR(120)    NOT NULL,
  `expires_in_days` 			INT 			NOT NULL,
  `is_public` 					TINYINT 		NOT NULL DEFAULT 1,
  `sort_order` 					INT 			NOT NULL,
  `is_active`                   TINYINT         NOT NULL DEFAULT 1,
  `created_at`                  TIMESTAMP       NOT NULL    DEFAULT CURRENT_TIMESTAMP,
  `updated_at`                  TIMESTAMP       NULL,
  `deleted_at`                  TIMESTAMP       NULL,
  PRIMARY KEY (`id`));
  