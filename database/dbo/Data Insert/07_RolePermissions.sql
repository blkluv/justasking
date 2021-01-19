#Owner AccessBilling
INSERT INTO role_permissions (role_id, permission_id, permission_value)
	SELECT * FROM (SELECT 'db9577c2-a7e1-11e7-9ec3-305a3a07203e', '68d15d89-a7e5-11e7-9ec3-305a3a07203e', true) AS tmp
		WHERE NOT EXISTS (
			SELECT role_id FROM role_permissions WHERE role_id = 'db9577c2-a7e1-11e7-9ec3-305a3a07203e' AND permission_id = '68d15d89-a7e5-11e7-9ec3-305a3a07203e'
);
#Owner SeeAllBoxes
INSERT INTO role_permissions (role_id, permission_id, permission_value)
	SELECT * FROM (SELECT 'db9577c2-a7e1-11e7-9ec3-305a3a07203e', '9985866b-a7e5-11e7-9ec3-305a3a07203e', true) AS tmp
		WHERE NOT EXISTS (
			SELECT role_id FROM role_permissions WHERE role_id = 'db9577c2-a7e1-11e7-9ec3-305a3a07203e' AND permission_id = '9985866b-a7e5-11e7-9ec3-305a3a07203e'
);
#Owner ManageUsers
INSERT INTO role_permissions (role_id, permission_id, permission_value)
	SELECT * FROM (SELECT 'db9577c2-a7e1-11e7-9ec3-305a3a07203e', 'f98d4be7-a7e5-11e7-9ec3-305a3a07203e', true) AS tmp
		WHERE NOT EXISTS (
			SELECT role_id FROM role_permissions WHERE role_id = 'db9577c2-a7e1-11e7-9ec3-305a3a07203e' AND permission_id = 'f98d4be7-a7e5-11e7-9ec3-305a3a07203e'
);
#Owner AccessAccountDetails
INSERT INTO role_permissions (role_id, permission_id, permission_value)
	SELECT * FROM (SELECT 'db9577c2-a7e1-11e7-9ec3-305a3a07203e', '0664a2ea-a7e6-11e7-9ec3-305a3a07203e', true) AS tmp
		WHERE NOT EXISTS (
			SELECT role_id FROM role_permissions WHERE role_id = 'db9577c2-a7e1-11e7-9ec3-305a3a07203e' AND permission_id = '0664a2ea-a7e6-11e7-9ec3-305a3a07203e'
);
#Owner CloseAccount
INSERT INTO role_permissions (role_id, permission_id, permission_value)
	SELECT * FROM (SELECT 'db9577c2-a7e1-11e7-9ec3-305a3a07203e', '25a27218-a7e7-11e7-9ec3-305a3a07203e', true) AS tmp
		WHERE NOT EXISTS (
			SELECT role_id FROM role_permissions WHERE role_id = 'db9577c2-a7e1-11e7-9ec3-305a3a07203e' AND permission_id = '25a27218-a7e7-11e7-9ec3-305a3a07203e'
);
#Owner ManageOwners
INSERT INTO role_permissions (role_id, permission_id, permission_value)
	SELECT * FROM (SELECT 'db9577c2-a7e1-11e7-9ec3-305a3a07203e', '333b293f-0d3e-11e8-8372-fe15f964686c', true) AS tmp
		WHERE NOT EXISTS (
			SELECT role_id FROM role_permissions WHERE role_id = 'db9577c2-a7e1-11e7-9ec3-305a3a07203e' AND permission_id = '333b293f-0d3e-11e8-8372-fe15f964686c'
);
#Owner EditAccountName
INSERT INTO role_permissions (role_id, permission_id, permission_value)
	SELECT * FROM (SELECT 'db9577c2-a7e1-11e7-9ec3-305a3a07203e', 'fad60991-0d3e-11e8-8372-fe15f964686c', true) AS tmp
		WHERE NOT EXISTS (
			SELECT role_id FROM role_permissions WHERE role_id = 'db9577c2-a7e1-11e7-9ec3-305a3a07203e' AND permission_id = 'fad60991-0d3e-11e8-8372-fe15f964686c'
);
#END------------------------------------------------------------------------------------------------------------------------------------------------------------

#BEGIN------------------------------------------------------------------------------------------------------------------------------------------------------------
#Admin AccessBilling
INSERT INTO role_permissions (role_id, permission_id, permission_value)
	SELECT * FROM (SELECT 'ff2fd925-a7e1-11e7-9ec3-305a3a07203e', '68d15d89-a7e5-11e7-9ec3-305a3a07203e', false) AS tmp
		WHERE NOT EXISTS (
			SELECT role_id FROM role_permissions WHERE role_id = 'ff2fd925-a7e1-11e7-9ec3-305a3a07203e' AND permission_id = '68d15d89-a7e5-11e7-9ec3-305a3a07203e'
);
#Admin SeeAllBoxes
INSERT INTO role_permissions (role_id, permission_id, permission_value)
	SELECT * FROM (SELECT 'ff2fd925-a7e1-11e7-9ec3-305a3a07203e', '9985866b-a7e5-11e7-9ec3-305a3a07203e', true) AS tmp
		WHERE NOT EXISTS (
			SELECT role_id FROM role_permissions WHERE role_id = 'ff2fd925-a7e1-11e7-9ec3-305a3a07203e' AND permission_id = '9985866b-a7e5-11e7-9ec3-305a3a07203e'
);
#Admin ManageUsers
INSERT INTO role_permissions (role_id, permission_id, permission_value)
	SELECT * FROM (SELECT 'ff2fd925-a7e1-11e7-9ec3-305a3a07203e', 'f98d4be7-a7e5-11e7-9ec3-305a3a07203e', true) AS tmp
		WHERE NOT EXISTS (
			SELECT role_id FROM role_permissions WHERE role_id = 'ff2fd925-a7e1-11e7-9ec3-305a3a07203e' AND permission_id = 'f98d4be7-a7e5-11e7-9ec3-305a3a07203e'
);
#Admin AccessAccountDetails
INSERT INTO role_permissions (role_id, permission_id, permission_value)
	SELECT * FROM (SELECT 'ff2fd925-a7e1-11e7-9ec3-305a3a07203e', '0664a2ea-a7e6-11e7-9ec3-305a3a07203e', true) AS tmp
		WHERE NOT EXISTS (
			SELECT role_id FROM role_permissions WHERE role_id = 'ff2fd925-a7e1-11e7-9ec3-305a3a07203e' AND permission_id = '0664a2ea-a7e6-11e7-9ec3-305a3a07203e'
);
#Admin CloseAccount
INSERT INTO role_permissions (role_id, permission_id, permission_value)
	SELECT * FROM (SELECT 'ff2fd925-a7e1-11e7-9ec3-305a3a07203e', '25a27218-a7e7-11e7-9ec3-305a3a07203e', false) AS tmp
		WHERE NOT EXISTS (
			SELECT role_id FROM role_permissions WHERE role_id = 'ff2fd925-a7e1-11e7-9ec3-305a3a07203e' AND permission_id = '25a27218-a7e7-11e7-9ec3-305a3a07203e'
);
#Admin ManageOwners
INSERT INTO role_permissions (role_id, permission_id, permission_value)
	SELECT * FROM (SELECT 'ff2fd925-a7e1-11e7-9ec3-305a3a07203e', '333b293f-0d3e-11e8-8372-fe15f964686c', false) AS tmp
		WHERE NOT EXISTS (
			SELECT role_id FROM role_permissions WHERE role_id = 'ff2fd925-a7e1-11e7-9ec3-305a3a07203e' AND permission_id = '333b293f-0d3e-11e8-8372-fe15f964686c'
);
#Admin EditAccountName
INSERT INTO role_permissions (role_id, permission_id, permission_value)
	SELECT * FROM (SELECT 'ff2fd925-a7e1-11e7-9ec3-305a3a07203e', 'fad60991-0d3e-11e8-8372-fe15f964686c', false) AS tmp
		WHERE NOT EXISTS (
			SELECT role_id FROM role_permissions WHERE role_id = 'ff2fd925-a7e1-11e7-9ec3-305a3a07203e' AND permission_id = 'fad60991-0d3e-11e8-8372-fe15f964686c'
);
#END------------------------------------------------------------------------------------------------------------------------------------------------------------

#BEGIN------------------------------------------------------------------------------------------------------------------------------------------------------------
#Presenter AccessBilling
INSERT INTO role_permissions (role_id, permission_id, permission_value)
	SELECT * FROM (SELECT '3113af99-a7e2-11e7-9ec3-305a3a07203e', '68d15d89-a7e5-11e7-9ec3-305a3a07203e', false) AS tmp
		WHERE NOT EXISTS (
			SELECT role_id FROM role_permissions WHERE role_id = '3113af99-a7e2-11e7-9ec3-305a3a07203e' AND permission_id = '68d15d89-a7e5-11e7-9ec3-305a3a07203e'
);
#Presenter SeeAllBoxes
INSERT INTO role_permissions (role_id, permission_id, permission_value)
	SELECT * FROM (SELECT '3113af99-a7e2-11e7-9ec3-305a3a07203e', '9985866b-a7e5-11e7-9ec3-305a3a07203e', false) AS tmp
		WHERE NOT EXISTS (
			SELECT role_id FROM role_permissions WHERE role_id = '3113af99-a7e2-11e7-9ec3-305a3a07203e' AND permission_id = '9985866b-a7e5-11e7-9ec3-305a3a07203e'
);
#Presenter ManageUsers
INSERT INTO role_permissions (role_id, permission_id, permission_value)
	SELECT * FROM (SELECT '3113af99-a7e2-11e7-9ec3-305a3a07203e', 'f98d4be7-a7e5-11e7-9ec3-305a3a07203e', false) AS tmp
		WHERE NOT EXISTS (
			SELECT role_id FROM role_permissions WHERE role_id = '3113af99-a7e2-11e7-9ec3-305a3a07203e' AND permission_id = 'f98d4be7-a7e5-11e7-9ec3-305a3a07203e'
);
#Presenter AccessAccountDetails
INSERT INTO role_permissions (role_id, permission_id, permission_value)
	SELECT * FROM (SELECT '3113af99-a7e2-11e7-9ec3-305a3a07203e', '0664a2ea-a7e6-11e7-9ec3-305a3a07203e', false) AS tmp
		WHERE NOT EXISTS (
			SELECT role_id FROM role_permissions WHERE role_id = '3113af99-a7e2-11e7-9ec3-305a3a07203e' AND permission_id = '0664a2ea-a7e6-11e7-9ec3-305a3a07203e'
);
#Presenter CloseAccount
INSERT INTO role_permissions (role_id, permission_id, permission_value)
	SELECT * FROM (SELECT '3113af99-a7e2-11e7-9ec3-305a3a07203e', '25a27218-a7e7-11e7-9ec3-305a3a07203e', false) AS tmp
		WHERE NOT EXISTS (
			SELECT role_id FROM role_permissions WHERE role_id = '3113af99-a7e2-11e7-9ec3-305a3a07203e' AND permission_id = '25a27218-a7e7-11e7-9ec3-305a3a07203e'
);
#Presenter ManageOwners
INSERT INTO role_permissions (role_id, permission_id, permission_value)
	SELECT * FROM (SELECT '3113af99-a7e2-11e7-9ec3-305a3a07203e', '333b293f-0d3e-11e8-8372-fe15f964686c', false) AS tmp
		WHERE NOT EXISTS (
			SELECT role_id FROM role_permissions WHERE role_id = '3113af99-a7e2-11e7-9ec3-305a3a07203e' AND permission_id = '333b293f-0d3e-11e8-8372-fe15f964686c'
);
#Presenter EditAccountName
INSERT INTO role_permissions (role_id, permission_id, permission_value)
	SELECT * FROM (SELECT '3113af99-a7e2-11e7-9ec3-305a3a07203e', 'fad60991-0d3e-11e8-8372-fe15f964686c', false) AS tmp
		WHERE NOT EXISTS (
			SELECT role_id FROM role_permissions WHERE role_id = '3113af99-a7e2-11e7-9ec3-305a3a07203e' AND permission_id = 'fad60991-0d3e-11e8-8372-fe15f964686c'
);
#END------------------------------------------------------------------------------------------------------------------------------------------------------------

