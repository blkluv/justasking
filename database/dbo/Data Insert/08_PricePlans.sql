INSERT INTO price_plans (id, name, display_name, description, price, price_description, image_path, expires_in_days, is_public, is_active, sort_order)
	SELECT * FROM (SELECT '1158c16d-afb2-11e7-a739-305a3a07203e', 'BASIC' as name, 'BASIC' as display_name, 'You will not be charged for using this plan.', 0, '$0', '/assets/graphics/starter.png', 0 as expires_in_days, 1 as is_public, 1, 0 as sort_order) AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM price_plans WHERE name = 'BASIC'
);

UPDATE price_plans SET expires_in_days = 0 WHERE id = '1158c16d-afb2-11e7-a739-305a3a07203e';
UPDATE price_plans SET sort_order = 0 WHERE id = '1158c16d-afb2-11e7-a739-305a3a07203e';

INSERT INTO price_plans (id, name, display_name, description, price, price_description, image_path, expires_in_days, is_public, is_active, sort_order)
	SELECT * FROM (SELECT '7634512a-1781-11e8-b176-54ee75ba93ea', 'PREMIUM-WEEK' as name, 'PREMIUM' as display_name, 'You will be charged once, and will be able to use PREMIUM features for one week.', 100, '$100', '/assets/graphics/crowd.png', 7, 1 as is_public, 1, 1 as sort_order) AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM price_plans WHERE name = 'PREMIUM-WEEK'
);

UPDATE price_plans SET expires_in_days = 7 WHERE id = '7634512a-1781-11e8-b176-54ee75ba93ea';
UPDATE price_plans SET sort_order = 1 WHERE id = '7634512a-1781-11e8-b176-54ee75ba93ea';

INSERT INTO price_plans (id, name, display_name, description, price, price_description, image_path, expires_in_days, is_public, is_active, sort_order)
	SELECT * FROM (SELECT '414169eb-afb2-11e7-a739-305a3a07203e', 'PREMIUM-MONTH' as name, 'PREMIUM' as display_name, 'You will be charged once, and will be able to use PREMIUM features for one month.', 150, '$150', '/assets/graphics/crowd.png', 30, 1 as is_public, 1, 2 as sort_order) AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM price_plans WHERE name = 'PREMIUM-MONTH'
);

UPDATE price_plans SET expires_in_days = 30 WHERE id = '414169eb-afb2-11e7-a739-305a3a07203e';
UPDATE price_plans SET sort_order = 2 WHERE id = '414169eb-afb2-11e7-a739-305a3a07203e';

INSERT INTO price_plans (id, name, display_name, description, price, price_description, image_path, expires_in_days, is_public, is_active, sort_order)
	SELECT * FROM (SELECT '414169eb-afb2-11e7-a739-305a3a07203f', 'PREMIUM-YEAR' as name, 'PREMIUM' as display_name, 'You will be charged once, and will be able to use PREMIUM features for one year.', 750, '$750', '/assets/graphics/crowd.png', 365, 1 as is_public, 1, 3 as sort_order) AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM price_plans WHERE name = 'PREMIUM-YEAR'
);

UPDATE price_plans SET expires_in_days = 365 WHERE id = '414169eb-afb2-11e7-a739-305a3a07203f';
UPDATE price_plans SET sort_order = 3 WHERE id = '414169eb-afb2-11e7-a739-305a3a07203f';
