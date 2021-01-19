INSERT INTO themes (name, value)
	SELECT * FROM (SELECT 'Red', 'core-red') AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM justasking.themes WHERE value = 'core-red'
);
INSERT INTO themes (name, value)
	SELECT * FROM (SELECT 'Pink', 'core-pink') AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM justasking.themes WHERE value = 'core-pink'
);
INSERT INTO themes (name, value)
	SELECT * FROM (SELECT 'Purple', 'core-purple') AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM justasking.themes WHERE value = 'core-purple'
);
INSERT INTO themes (name, value)
	SELECT * FROM (SELECT 'Deep-purple', 'core-deep-purple') AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM justasking.themes WHERE value = 'core-deep-purple'
);
INSERT INTO themes (name, value)
	SELECT * FROM (SELECT 'Indigo', 'core-indigo') AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM justasking.themes WHERE value = 'core-indigo'
);
INSERT INTO themes (name, value)
	SELECT * FROM (SELECT 'Blue', 'core-blue') AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM justasking.themes WHERE value = 'core-blue'
);
INSERT INTO themes (name, value)
	SELECT * FROM (SELECT 'Light-blue', 'core-light-blue') AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM justasking.themes WHERE value = 'core-light-blue'
);
INSERT INTO themes (name, value)
	SELECT * FROM (SELECT 'Cyan', 'core-cyan') AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM justasking.themes WHERE value = 'core-cyan'
);
INSERT INTO themes (name, value)
	SELECT * FROM (SELECT 'Teal', 'core-teal') AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM justasking.themes WHERE value = 'core-teal'
);
INSERT INTO themes (name, value)
	SELECT * FROM (SELECT 'Green', 'core-green') AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM justasking.themes WHERE value = 'core-green'
);
INSERT INTO themes (name, value)
	SELECT * FROM (SELECT 'Light-green', 'core-light-green') AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM justasking.themes WHERE value = 'core-light-green'
);
INSERT INTO themes (name, value)
	SELECT * FROM (SELECT 'Lime', 'core-lime') AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM justasking.themes WHERE value = 'core-lime'
);
INSERT INTO themes (name, value)
	SELECT * FROM (SELECT 'Yellow', 'core-yellow') AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM justasking.themes WHERE value = 'core-yellow'
);
INSERT INTO themes (name, value)
	SELECT * FROM (SELECT 'Amber', 'core-amber') AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM justasking.themes WHERE value = 'core-amber'
);
INSERT INTO themes (name, value)
	SELECT * FROM (SELECT 'Orange', 'core-orange') AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM justasking.themes WHERE value = 'core-orange'
);
INSERT INTO themes (name, value)
	SELECT * FROM (SELECT 'Deep-orange', 'core-deep-orange') AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM justasking.themes WHERE value = 'core-deep-orange'
);
INSERT INTO themes (name, value)
	SELECT * FROM (SELECT 'Brown', 'core-brown') AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM justasking.themes WHERE value = 'core-brown'
);
INSERT INTO themes (name, value)
	SELECT * FROM (SELECT 'Grey', 'core-grey') AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM justasking.themes WHERE value = 'core-grey'
);
INSERT INTO themes (name, value)
	SELECT * FROM (SELECT 'Blue-grey', 'core-blue-grey') AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM justasking.themes WHERE value = 'core-blue-grey'
);
