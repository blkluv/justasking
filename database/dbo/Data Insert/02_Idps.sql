INSERT INTO idps (name)
	SELECT * FROM (SELECT 'google') AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM justasking.idps WHERE name = 'google'
);

INSERT INTO idps (name)
	SELECT * FROM (SELECT 'justasking') AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM justasking.idps WHERE name = 'justasking'
);