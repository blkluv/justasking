/*BEGIN BASIC******************************************************/
INSERT INTO price_plan_features (id, plan_id, feature_id, feature_value )
	SELECT * FROM (SELECT UUID(), '1158c16d-afb2-11e7-a739-305a3a07203e', '0f534226-afb4-11e7-a739-305a3a07203e', '1') AS tmp
		WHERE NOT EXISTS (
			SELECT id FROM price_plan_features WHERE plan_id = '1158c16d-afb2-11e7-a739-305a3a07203e' AND feature_id = '0f534226-afb4-11e7-a739-305a3a07203e'
);

INSERT INTO price_plan_features (id, plan_id, feature_id, feature_value )
	SELECT * FROM (SELECT UUID(), '1158c16d-afb2-11e7-a739-305a3a07203e', '1754f6b2-afb4-11e7-a739-305a3a07203e', 'true') AS tmp
		WHERE NOT EXISTS (
			SELECT id FROM price_plan_features WHERE plan_id = '1158c16d-afb2-11e7-a739-305a3a07203e' AND feature_id = '1754f6b2-afb4-11e7-a739-305a3a07203e'
);

INSERT INTO price_plan_features (id, plan_id, feature_id, feature_value )
	SELECT * FROM (SELECT UUID(), '1158c16d-afb2-11e7-a739-305a3a07203e', '2b813934-afb4-11e7-a739-305a3a07203e', 'true') AS tmp
		WHERE NOT EXISTS (
			SELECT id FROM price_plan_features WHERE plan_id = '1158c16d-afb2-11e7-a739-305a3a07203e' AND feature_id = '2b813934-afb4-11e7-a739-305a3a07203e'
);

INSERT INTO price_plan_features (id, plan_id, feature_id, feature_value )
	SELECT * FROM (SELECT UUID(), '1158c16d-afb2-11e7-a739-305a3a07203e', '2fc98dcf-afb4-11e7-a739-305a3a07203e', 'true') AS tmp
		WHERE NOT EXISTS (
			SELECT id FROM price_plan_features WHERE plan_id = '1158c16d-afb2-11e7-a739-305a3a07203e' AND feature_id = '2fc98dcf-afb4-11e7-a739-305a3a07203e'
);

INSERT INTO price_plan_features (id, plan_id, feature_id, feature_value )
	SELECT * FROM (SELECT UUID(), '1158c16d-afb2-11e7-a739-305a3a07203e', '387ec1be-afb4-11e7-a739-305a3a07203e', 'true') AS tmp
		WHERE NOT EXISTS (
			SELECT id FROM price_plan_features WHERE plan_id = '1158c16d-afb2-11e7-a739-305a3a07203e' AND feature_id = '387ec1be-afb4-11e7-a739-305a3a07203e'
);

INSERT INTO price_plan_features (id, plan_id, feature_id, feature_value )
	SELECT * FROM (SELECT UUID(), '1158c16d-afb2-11e7-a739-305a3a07203e', '3e93ce5f-afb4-11e7-a739-305a3a07203e', 'true') AS tmp
		WHERE NOT EXISTS (
			SELECT id FROM price_plan_features WHERE plan_id = '1158c16d-afb2-11e7-a739-305a3a07203e' AND feature_id = '3e93ce5f-afb4-11e7-a739-305a3a07203e'
);

INSERT INTO price_plan_features (id, plan_id, feature_id, feature_value )
	SELECT * FROM (SELECT UUID(), '1158c16d-afb2-11e7-a739-305a3a07203e', '44402b56-afb4-11e7-a739-305a3a07203e', 'false') AS tmp
		WHERE NOT EXISTS (
			SELECT id FROM price_plan_features WHERE plan_id = '1158c16d-afb2-11e7-a739-305a3a07203e' AND feature_id = '44402b56-afb4-11e7-a739-305a3a07203e'
);

INSERT INTO price_plan_features (id, plan_id, feature_id, feature_value )
	SELECT * FROM (SELECT UUID(), '1158c16d-afb2-11e7-a739-305a3a07203e', '48ca8ac8-afb4-11e7-a739-305a3a07203e', 'false') AS tmp
		WHERE NOT EXISTS (
			SELECT id FROM price_plan_features WHERE plan_id = '1158c16d-afb2-11e7-a739-305a3a07203e' AND feature_id = '48ca8ac8-afb4-11e7-a739-305a3a07203e'
);

INSERT INTO price_plan_features (id, plan_id, feature_id, feature_value )
	SELECT * FROM (SELECT UUID(), '1158c16d-afb2-11e7-a739-305a3a07203e', '4d7f6826-afb4-11e7-a739-305a3a07203e', '0') AS tmp
		WHERE NOT EXISTS (
			SELECT id FROM price_plan_features WHERE plan_id = '1158c16d-afb2-11e7-a739-305a3a07203e' AND feature_id = '4d7f6826-afb4-11e7-a739-305a3a07203e'
);

INSERT INTO price_plan_features (id, plan_id, feature_id, feature_value )
	SELECT * FROM (SELECT UUID(), '1158c16d-afb2-11e7-a739-305a3a07203e', '538512cc-afb4-11e7-a739-305a3a07203e', 'false') AS tmp
		WHERE NOT EXISTS (
			SELECT id FROM price_plan_features WHERE plan_id = '1158c16d-afb2-11e7-a739-305a3a07203e' AND feature_id = '538512cc-afb4-11e7-a739-305a3a07203e'
);

UPDATE price_plan_features SET feature_value = 'false' WHERE plan_id = '1158c16d-afb2-11e7-a739-305a3a07203e' AND feature_id = '538512cc-afb4-11e7-a739-305a3a07203e';
/*END BASIC******************************************************/

/*BEGIN PREMIUM-WEEK******************************************************/
INSERT INTO price_plan_features (id, plan_id, feature_id, feature_value )
	SELECT * FROM (SELECT UUID(), '7634512a-1781-11e8-b176-54ee75ba93ea', '0f534226-afb4-11e7-a739-305a3a07203e', '5') AS tmp
		WHERE NOT EXISTS (
			SELECT id FROM price_plan_features WHERE plan_id = '7634512a-1781-11e8-b176-54ee75ba93ea' AND feature_id = '0f534226-afb4-11e7-a739-305a3a07203e'
);

INSERT INTO price_plan_features (id, plan_id, feature_id, feature_value )
	SELECT * FROM (SELECT UUID(), '7634512a-1781-11e8-b176-54ee75ba93ea', '1754f6b2-afb4-11e7-a739-305a3a07203e', 'true') AS tmp
		WHERE NOT EXISTS (
			SELECT id FROM price_plan_features WHERE plan_id = '7634512a-1781-11e8-b176-54ee75ba93ea' AND feature_id = '1754f6b2-afb4-11e7-a739-305a3a07203e'
);


INSERT INTO price_plan_features (id, plan_id, feature_id, feature_value )
	SELECT * FROM (SELECT UUID(), '7634512a-1781-11e8-b176-54ee75ba93ea', '2b813934-afb4-11e7-a739-305a3a07203e', 'true') AS tmp
		WHERE NOT EXISTS (
			SELECT id FROM price_plan_features WHERE plan_id = '7634512a-1781-11e8-b176-54ee75ba93ea' AND feature_id = '2b813934-afb4-11e7-a739-305a3a07203e'
);


INSERT INTO price_plan_features (id, plan_id, feature_id, feature_value )
	SELECT * FROM (SELECT UUID(), '7634512a-1781-11e8-b176-54ee75ba93ea', '2fc98dcf-afb4-11e7-a739-305a3a07203e', 'true') AS tmp
		WHERE NOT EXISTS (
			SELECT id FROM price_plan_features WHERE plan_id = '7634512a-1781-11e8-b176-54ee75ba93ea' AND feature_id = '2fc98dcf-afb4-11e7-a739-305a3a07203e'
);

INSERT INTO price_plan_features (id, plan_id, feature_id, feature_value )
	SELECT * FROM (SELECT UUID(), '7634512a-1781-11e8-b176-54ee75ba93ea', '387ec1be-afb4-11e7-a739-305a3a07203e', 'true') AS tmp
		WHERE NOT EXISTS (
			SELECT id FROM price_plan_features WHERE plan_id = '7634512a-1781-11e8-b176-54ee75ba93ea' AND feature_id = '387ec1be-afb4-11e7-a739-305a3a07203e'
);


INSERT INTO price_plan_features (id, plan_id, feature_id, feature_value )
	SELECT * FROM (SELECT UUID(), '7634512a-1781-11e8-b176-54ee75ba93ea', '3e93ce5f-afb4-11e7-a739-305a3a07203e', 'true') AS tmp
		WHERE NOT EXISTS (
			SELECT id FROM price_plan_features WHERE plan_id = '7634512a-1781-11e8-b176-54ee75ba93ea' AND feature_id = '3e93ce5f-afb4-11e7-a739-305a3a07203e'
);


INSERT INTO price_plan_features (id, plan_id, feature_id, feature_value )
	SELECT * FROM (SELECT UUID(), '7634512a-1781-11e8-b176-54ee75ba93ea', '44402b56-afb4-11e7-a739-305a3a07203e', 'true') AS tmp
		WHERE NOT EXISTS (
			SELECT id FROM price_plan_features WHERE plan_id = '7634512a-1781-11e8-b176-54ee75ba93ea' AND feature_id = '44402b56-afb4-11e7-a739-305a3a07203e'
);

INSERT INTO price_plan_features (id, plan_id, feature_id, feature_value )
	SELECT * FROM (SELECT UUID(), '7634512a-1781-11e8-b176-54ee75ba93ea', '48ca8ac8-afb4-11e7-a739-305a3a07203e', 'true') AS tmp
		WHERE NOT EXISTS (
			SELECT id FROM price_plan_features WHERE plan_id = '7634512a-1781-11e8-b176-54ee75ba93ea' AND feature_id = '48ca8ac8-afb4-11e7-a739-305a3a07203e'
);


INSERT INTO price_plan_features (id, plan_id, feature_id, feature_value )
	SELECT * FROM (SELECT UUID(), '7634512a-1781-11e8-b176-54ee75ba93ea', '4d7f6826-afb4-11e7-a739-305a3a07203e', '5') AS tmp
		WHERE NOT EXISTS (
			SELECT id FROM price_plan_features WHERE plan_id = '7634512a-1781-11e8-b176-54ee75ba93ea' AND feature_id = '4d7f6826-afb4-11e7-a739-305a3a07203e'
);


INSERT INTO price_plan_features (id, plan_id, feature_id, feature_value )
	SELECT * FROM (SELECT UUID(), '7634512a-1781-11e8-b176-54ee75ba93ea', '538512cc-afb4-11e7-a739-305a3a07203e', 'true') AS tmp
		WHERE NOT EXISTS (
			SELECT id FROM price_plan_features WHERE plan_id = '7634512a-1781-11e8-b176-54ee75ba93ea' AND feature_id = '538512cc-afb4-11e7-a739-305a3a07203e'
);
/*END PREMIUM-MONTH******************************************************/

/*BEGIN PREMIUM-MONTH******************************************************/
INSERT INTO price_plan_features (id, plan_id, feature_id, feature_value )
	SELECT * FROM (SELECT UUID(), '414169eb-afb2-11e7-a739-305a3a07203e', '0f534226-afb4-11e7-a739-305a3a07203e', '5') AS tmp
		WHERE NOT EXISTS (
			SELECT id FROM price_plan_features WHERE plan_id = '414169eb-afb2-11e7-a739-305a3a07203e' AND feature_id = '0f534226-afb4-11e7-a739-305a3a07203e'
);

INSERT INTO price_plan_features (id, plan_id, feature_id, feature_value )
	SELECT * FROM (SELECT UUID(), '414169eb-afb2-11e7-a739-305a3a07203e', '1754f6b2-afb4-11e7-a739-305a3a07203e', 'true') AS tmp
		WHERE NOT EXISTS (
			SELECT id FROM price_plan_features WHERE plan_id = '414169eb-afb2-11e7-a739-305a3a07203e' AND feature_id = '1754f6b2-afb4-11e7-a739-305a3a07203e'
);

INSERT INTO price_plan_features (id, plan_id, feature_id, feature_value )
	SELECT * FROM (SELECT UUID(), '414169eb-afb2-11e7-a739-305a3a07203e', '2b813934-afb4-11e7-a739-305a3a07203e', 'true') AS tmp
		WHERE NOT EXISTS (
			SELECT id FROM price_plan_features WHERE plan_id = '414169eb-afb2-11e7-a739-305a3a07203e' AND feature_id = '2b813934-afb4-11e7-a739-305a3a07203e'
);


INSERT INTO price_plan_features (id, plan_id, feature_id, feature_value )
	SELECT * FROM (SELECT UUID(), '414169eb-afb2-11e7-a739-305a3a07203e', '2fc98dcf-afb4-11e7-a739-305a3a07203e', 'true') AS tmp
		WHERE NOT EXISTS (
			SELECT id FROM price_plan_features WHERE plan_id = '414169eb-afb2-11e7-a739-305a3a07203e' AND feature_id = '2fc98dcf-afb4-11e7-a739-305a3a07203e'
);

INSERT INTO price_plan_features (id, plan_id, feature_id, feature_value )
	SELECT * FROM (SELECT UUID(), '414169eb-afb2-11e7-a739-305a3a07203e', '387ec1be-afb4-11e7-a739-305a3a07203e', 'true') AS tmp
		WHERE NOT EXISTS (
			SELECT id FROM price_plan_features WHERE plan_id = '414169eb-afb2-11e7-a739-305a3a07203e' AND feature_id = '387ec1be-afb4-11e7-a739-305a3a07203e'
);


INSERT INTO price_plan_features (id, plan_id, feature_id, feature_value )
	SELECT * FROM (SELECT UUID(), '414169eb-afb2-11e7-a739-305a3a07203e', '3e93ce5f-afb4-11e7-a739-305a3a07203e', 'true') AS tmp
		WHERE NOT EXISTS (
			SELECT id FROM price_plan_features WHERE plan_id = '414169eb-afb2-11e7-a739-305a3a07203e' AND feature_id = '3e93ce5f-afb4-11e7-a739-305a3a07203e'
);


INSERT INTO price_plan_features (id, plan_id, feature_id, feature_value )
	SELECT * FROM (SELECT UUID(), '414169eb-afb2-11e7-a739-305a3a07203e', '44402b56-afb4-11e7-a739-305a3a07203e', 'true') AS tmp
		WHERE NOT EXISTS (
			SELECT id FROM price_plan_features WHERE plan_id = '414169eb-afb2-11e7-a739-305a3a07203e' AND feature_id = '44402b56-afb4-11e7-a739-305a3a07203e'
);

INSERT INTO price_plan_features (id, plan_id, feature_id, feature_value )
	SELECT * FROM (SELECT UUID(), '414169eb-afb2-11e7-a739-305a3a07203e', '48ca8ac8-afb4-11e7-a739-305a3a07203e', 'true') AS tmp
		WHERE NOT EXISTS (
			SELECT id FROM price_plan_features WHERE plan_id = '414169eb-afb2-11e7-a739-305a3a07203e' AND feature_id = '48ca8ac8-afb4-11e7-a739-305a3a07203e'
);


INSERT INTO price_plan_features (id, plan_id, feature_id, feature_value )
	SELECT * FROM (SELECT UUID(), '414169eb-afb2-11e7-a739-305a3a07203e', '4d7f6826-afb4-11e7-a739-305a3a07203e', '5') AS tmp
		WHERE NOT EXISTS (
			SELECT id FROM price_plan_features WHERE plan_id = '414169eb-afb2-11e7-a739-305a3a07203e' AND feature_id = '4d7f6826-afb4-11e7-a739-305a3a07203e'
);


INSERT INTO price_plan_features (id, plan_id, feature_id, feature_value )
	SELECT * FROM (SELECT UUID(), '414169eb-afb2-11e7-a739-305a3a07203e', '538512cc-afb4-11e7-a739-305a3a07203e', 'true') AS tmp
		WHERE NOT EXISTS (
			SELECT id FROM price_plan_features WHERE plan_id = '414169eb-afb2-11e7-a739-305a3a07203e' AND feature_id = '538512cc-afb4-11e7-a739-305a3a07203e'
);
/*END PREMIUM-MONTH******************************************************/

/*BEGIN PREMIUM-YEAR******************************************************/
INSERT INTO price_plan_features (id, plan_id, feature_id, feature_value )
	SELECT * FROM (SELECT UUID(), '414169eb-afb2-11e7-a739-305a3a07203f', '0f534226-afb4-11e7-a739-305a3a07203e', '5') AS tmp
		WHERE NOT EXISTS (
			SELECT id FROM price_plan_features WHERE plan_id = '414169eb-afb2-11e7-a739-305a3a07203f' AND feature_id = '0f534226-afb4-11e7-a739-305a3a07203e'
);

INSERT INTO price_plan_features (id, plan_id, feature_id, feature_value )
	SELECT * FROM (SELECT UUID(), '414169eb-afb2-11e7-a739-305a3a07203f', '1754f6b2-afb4-11e7-a739-305a3a07203e', 'true') AS tmp
		WHERE NOT EXISTS (
			SELECT id FROM price_plan_features WHERE plan_id = '414169eb-afb2-11e7-a739-305a3a07203f' AND feature_id = '1754f6b2-afb4-11e7-a739-305a3a07203e'
);

INSERT INTO price_plan_features (id, plan_id, feature_id, feature_value )
	SELECT * FROM (SELECT UUID(), '414169eb-afb2-11e7-a739-305a3a07203f', '2b813934-afb4-11e7-a739-305a3a07203e', 'true') AS tmp
		WHERE NOT EXISTS (
			SELECT id FROM price_plan_features WHERE plan_id = '414169eb-afb2-11e7-a739-305a3a07203f' AND feature_id = '2b813934-afb4-11e7-a739-305a3a07203e'
);

INSERT INTO price_plan_features (id, plan_id, feature_id, feature_value )
	SELECT * FROM (SELECT UUID(), '414169eb-afb2-11e7-a739-305a3a07203f', '2fc98dcf-afb4-11e7-a739-305a3a07203e', 'true') AS tmp
		WHERE NOT EXISTS (
			SELECT id FROM price_plan_features WHERE plan_id = '414169eb-afb2-11e7-a739-305a3a07203f' AND feature_id = '2fc98dcf-afb4-11e7-a739-305a3a07203e'
);

INSERT INTO price_plan_features (id, plan_id, feature_id, feature_value )
	SELECT * FROM (SELECT UUID(), '414169eb-afb2-11e7-a739-305a3a07203f', '387ec1be-afb4-11e7-a739-305a3a07203e', 'true') AS tmp
		WHERE NOT EXISTS (
			SELECT id FROM price_plan_features WHERE plan_id = '414169eb-afb2-11e7-a739-305a3a07203f' AND feature_id = '387ec1be-afb4-11e7-a739-305a3a07203e'
);

INSERT INTO price_plan_features (id, plan_id, feature_id, feature_value )
	SELECT * FROM (SELECT UUID(), '414169eb-afb2-11e7-a739-305a3a07203f', '3e93ce5f-afb4-11e7-a739-305a3a07203e', 'true') AS tmp
		WHERE NOT EXISTS (
			SELECT id FROM price_plan_features WHERE plan_id = '414169eb-afb2-11e7-a739-305a3a07203f' AND feature_id = '3e93ce5f-afb4-11e7-a739-305a3a07203e'
);

INSERT INTO price_plan_features (id, plan_id, feature_id, feature_value )
	SELECT * FROM (SELECT UUID(), '414169eb-afb2-11e7-a739-305a3a07203f', '44402b56-afb4-11e7-a739-305a3a07203e', 'true') AS tmp
		WHERE NOT EXISTS (
			SELECT id FROM price_plan_features WHERE plan_id = '414169eb-afb2-11e7-a739-305a3a07203f' AND feature_id = '44402b56-afb4-11e7-a739-305a3a07203e'
);

INSERT INTO price_plan_features (id, plan_id, feature_id, feature_value )
	SELECT * FROM (SELECT UUID(), '414169eb-afb2-11e7-a739-305a3a07203f', '48ca8ac8-afb4-11e7-a739-305a3a07203e', 'true') AS tmp
		WHERE NOT EXISTS (
			SELECT id FROM price_plan_features WHERE plan_id = '414169eb-afb2-11e7-a739-305a3a07203f' AND feature_id = '48ca8ac8-afb4-11e7-a739-305a3a07203e'
);

INSERT INTO price_plan_features (id, plan_id, feature_id, feature_value )
	SELECT * FROM (SELECT UUID(), '414169eb-afb2-11e7-a739-305a3a07203f', '4d7f6826-afb4-11e7-a739-305a3a07203e', '5') AS tmp
		WHERE NOT EXISTS (
			SELECT id FROM price_plan_features WHERE plan_id = '414169eb-afb2-11e7-a739-305a3a07203f' AND feature_id = '4d7f6826-afb4-11e7-a739-305a3a07203e'
);

INSERT INTO price_plan_features (id, plan_id, feature_id, feature_value )
	SELECT * FROM (SELECT UUID(), '414169eb-afb2-11e7-a739-305a3a07203f', '538512cc-afb4-11e7-a739-305a3a07203e', 'true') AS tmp
		WHERE NOT EXISTS (
			SELECT id FROM price_plan_features WHERE plan_id = '414169eb-afb2-11e7-a739-305a3a07203f' AND feature_id = '538512cc-afb4-11e7-a739-305a3a07203e'
);
/*END PREMIUM-YEAR******************************************************/