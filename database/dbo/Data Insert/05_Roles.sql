INSERT INTO roles (id, name)
	SELECT * FROM (SELECT 'db9577c2-a7e1-11e7-9ec3-305a3a07203e', 'Owner') AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM roles WHERE name = 'Owner'
);
INSERT INTO roles (id, name)
	SELECT * FROM (SELECT 'ff2fd925-a7e1-11e7-9ec3-305a3a07203e', 'Admin') AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM roles WHERE name = 'Admin'
);
INSERT INTO roles (id, name)
	SELECT * FROM (SELECT '3113af99-a7e2-11e7-9ec3-305a3a07203e', 'Presenter') AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM roles WHERE name = 'Presenter'
);