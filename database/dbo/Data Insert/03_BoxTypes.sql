INSERT INTO box_type (name, description)
	SELECT * FROM (SELECT 'wordcloud', 'Box type for word clouds') AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM justasking.box_type WHERE name = 'wordcloud'
);
INSERT INTO box_type (name, description)
	SELECT * FROM (SELECT 'questionbox', 'Box type for people who are accepting questions from their audience') AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM justasking.box_type WHERE name = 'questionbox'
);
INSERT INTO box_type (name, description)
	SELECT * FROM (SELECT 'answerbox', 'Box type for people who are accepting answers from their audience') AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM justasking.box_type WHERE name = 'answerbox'
);
INSERT INTO box_type (name, description)
	SELECT * FROM (SELECT 'votesbox', 'Box type for people who are polling their audience') AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM justasking.box_type WHERE name = 'votesbox'
);

