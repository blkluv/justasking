INSERT INTO permissions (id, name)
	SELECT * FROM (SELECT '68d15d89-a7e5-11e7-9ec3-305a3a07203e', 'AccessBilling') AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM permissions WHERE name = 'AccessBilling'
);
INSERT INTO permissions (id, name)
	SELECT * FROM (SELECT '9985866b-a7e5-11e7-9ec3-305a3a07203e', 'SeeAllBoxes') AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM permissions WHERE name = 'SeeAllBoxes'
);
INSERT INTO permissions (id, name)
	SELECT * FROM (SELECT 'f98d4be7-a7e5-11e7-9ec3-305a3a07203e', 'ManageUsers') AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM permissions WHERE name = 'ManageUsers'
);
INSERT INTO permissions (id, name)
	SELECT * FROM (SELECT '0664a2ea-a7e6-11e7-9ec3-305a3a07203e', 'AccessAccountDetails') AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM permissions WHERE name = 'AccessAccountDetails'
);
INSERT INTO permissions (id, name)
	SELECT * FROM (SELECT '25a27218-a7e7-11e7-9ec3-305a3a07203e', 'CloseAccount') AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM permissions WHERE name = 'CloseAccount'
);
INSERT INTO permissions (id, name)
	SELECT * FROM (SELECT '333b293f-0d3e-11e8-8372-fe15f964686c', 'ManageOwners') AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM permissions WHERE name = 'ManageOwners'
);
INSERT INTO permissions (id, name)
	SELECT * FROM (SELECT 'fad60991-0d3e-11e8-8372-fe15f964686c', 'EditAccountName') AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM permissions WHERE name = 'EditAccountName'
);

