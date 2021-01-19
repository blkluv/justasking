
#User's email
SET @email = 'arieldiaz92@gmail.com';

SET @userId = (SELECT u.id FROM `justasking`.`users` u WHERE u.email = @email);

#to be specified when we talk to customer
SET @accountId = (SELECT a.id FROM `justasking`.`accounts` a WHERE a.owner_id = @userId LIMIT 1);
  
#http://www.unit-conversion.info/texttools/random-string-generator/
SET @licenseCode = 'aouxFWPr6WM0PxpxtNzNOEkIPfAWHO4Ns2h5Yy9PPF6DuY6B0WretYKwZw4KK5EGg1YQfovYxqvHUSUF7SdvEFeMhodLoOEVdQS4FgzWgbC9Y1eTVnyZqi3YSxtFdNtoVGAJIiEtv8ejqEtDqgwA557To7s79WQTEtDyxx5rvbB9dWeyeNnzlLy7QjjN6DAh9QDhdkDP0YhA1tTOEYPP56GtLcXe1GEHYebCrkrEWb9rlmta2nxO8jiYXcoKj9nh';

SET @planName = 'CUSTOM - 2 Active polls',
	@planDescription = 'Mark - info@meadwebdesign.com',
	@planPrice = 75,
	@planExpiresInDays = 365,
	@activeBoxes = '2',
	@wordcloud = 'true',
	@questionBox = 'true',
	@answerBox = 'true',
	@votesBox = 'true',
	@toggleResponses = 'true',
	@sms = 'false',
	@customCode = 'false',
	@delegates = '0',
	@support = 'true',
	@createdBy = 'ariel';
  
CALL `justasking`.`CreateCustomPlanForAccount`(@accountId, @userId, @licenseCode, @planName, @planDescription, @planPrice, @planExpiresInDays, @activeBoxes, @wordcloud, @questionBox, @answerBox, @votesBox, @toggleResponses, @sms, @customCode, @delegates, @support, @createdBy);
