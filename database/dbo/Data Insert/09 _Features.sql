INSERT INTO features (id, name, description)
	SELECT * FROM (SELECT '0f534226-afb4-11e7-a739-305a3a07203e', 'Active Boxes', 'Maximum number of poll boxes that a registered user may have Open, accepting entries, at any given time.') AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM features WHERE name = 'Active Boxes'
);

INSERT INTO features (id, name, description)
	SELECT * FROM (SELECT '1754f6b2-afb4-11e7-a739-305a3a07203e', 'Wordcloud', 'Type of poll box where participants'' entries compose an image with words, in which the size of each word indicates its frequency.') AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM features WHERE name = 'Wordcloud'
);


INSERT INTO features (id, name, description)
	SELECT * FROM (SELECT '2b813934-afb4-11e7-a739-305a3a07203e', 'Question Box', 'Type of poll box where participants submit entries which can be upvoted and downvoted by other participants.') AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM features WHERE name = 'Question Box'
);


INSERT INTO features (id, name, description)
	SELECT * FROM (SELECT '2fc98dcf-afb4-11e7-a739-305a3a07203e', 'Answer Box', 'Type of poll box where participants can answer multiple questions at a time and box owner can activate/deactivate questions on the spot.') AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM features WHERE name = 'Answer Box'
);


INSERT INTO features (id, name, description)
	SELECT * FROM (SELECT '387ec1be-afb4-11e7-a739-305a3a07203e', 'Votes Box', 'Type of poll box where participants can vote from a selection of one or more multiple choise options.') AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM features WHERE name = 'Votes Box'
);


INSERT INTO features (id, name, description)
	SELECT * FROM (SELECT '3e93ce5f-afb4-11e7-a739-305a3a07203e', 'Toggle Responses', 'Poll box owner may toggle specific responses'' visibility at their discretion.') AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM features WHERE name = 'Toggle Responses'
);


INSERT INTO features (id, name, description)
	SELECT * FROM (SELECT '44402b56-afb4-11e7-a739-305a3a07203e', 'SMS', 'Participants are allowed to submit entries for poll boxes through SMS.') AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM features WHERE name = 'SMS'
);


INSERT INTO features (id, name, description)
	SELECT * FROM (SELECT '48ca8ac8-afb4-11e7-a739-305a3a07203e', 'Custom Code', 'Registered user is able to specify the code that will be used to access the poll box.') AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM features WHERE name = 'Custom Code'
);


INSERT INTO features (id, name, description)
	SELECT * FROM (SELECT '4d7f6826-afb4-11e7-a739-305a3a07203e', 'Delegates', 'Number of users that will share access to the justasking.io account.') AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM features WHERE name = 'Delegates'
);


INSERT INTO features (id, name, description)
	SELECT * FROM (SELECT '538512cc-afb4-11e7-a739-305a3a07203e', 'Support', 'Justasking.io support options.') AS tmp
		WHERE NOT EXISTS (
			SELECT name FROM features WHERE name = 'Support'
);
