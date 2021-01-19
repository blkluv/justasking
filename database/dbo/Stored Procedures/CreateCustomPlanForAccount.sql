DROP PROCEDURE IF EXISTS CreateCustomPlanForAccount;

DELIMITER //

CREATE procedure CreateCustomPlanForAccount 
( _accountId varchar(100)
, _userId char(36)
, _licenseCode varchar(256)
, _planName varchar(50)
, _planDescription varchar(250)
, _planPrice int(11)
, _planExpiresInDays int(11)
, _activeBoxes varchar(50)
, _wordcloud varchar(50)
, _questionBox varchar(50)
, _answerBox varchar(50)
, _votesBox varchar(50)
, _toggleResponses varchar(50)
, _sms varchar(50)
, _customCode varchar(50)
, _Delegates varchar(50)
, _Support varchar(50)
, _createdBy varchar(50)
)

  BEGIN
  
    DECLARE EXIT HANDLER FOR SQLEXCEPTION 
    BEGIN
		GET DIAGNOSTICS CONDITION 1 @sqlstate = RETURNED_SQLSTATE, 
		@errno = MYSQL_ERRNO, @text = MESSAGE_TEXT;
		SET @full_error = CONCAT("ERROR ", @errno, " (", @sqlstate, "): ", @text);
		SELECT @full_error;
		ROLLBACK;
    END;
    
    DECLARE exit handler for sqlwarning
	BEGIN
		GET DIAGNOSTICS CONDITION 1 @sqlstate = RETURNED_SQLSTATE, 
		@errno = MYSQL_ERRNO, @text = MESSAGE_TEXT;
		SET @full_error = CONCAT("ERROR ", @errno, " (", @sqlstate, "): ", @text);
		SELECT @full_error;    
		ROLLBACK;
	END;
    
    START TRANSACTION;
    
    SET @planId = uuid();
    SET @customPlanLicenseId = uuid();
    
    SET @activeBoxes_featureId = (SELECT f.id FROM `justasking`.`features` f WHERE f.name = 'Active Boxes'); 
	SET @wordcloud_featureId = (SELECT f.id FROM `justasking`.`features` f WHERE f.name = 'Wordcloud');
	SET @questionBox_featureId = (SELECT f.id FROM `justasking`.`features` f WHERE f.name = 'Question Box');
	SET @answerBox_featureId = (SELECT f.id FROM `justasking`.`features` f WHERE f.name = 'Answer Box');
	SET @votesBox_featureId = (SELECT f.id FROM `justasking`.`features` f WHERE f.name = 'Votes Box');
	SET @toggleResponses_featureId = (SELECT f.id FROM `justasking`.`features` f WHERE f.name = 'Toggle Responses');
	SET @sms_featureId = (SELECT f.id FROM `justasking`.`features` f WHERE f.name = 'SMS');
	SET @customCode_featureId = (SELECT f.id FROM `justasking`.`features` f WHERE f.name = 'Custom Code');
	SET @delegates_featureId = (SELECT f.id FROM `justasking`.`features` f WHERE f.name = 'Delegates');
	SET @support_featureId = (SELECT f.id FROM `justasking`.`features` f WHERE f.name = 'Support');
    
    #Create custom plan 
    INSERT INTO `justasking`.`price_plans`(`id`,`name`,`description`,`display_name`,`price`,`price_description`,`image_path`,`expires_in_days`,`sort_order`,`is_public`,`is_active`,`created_at`,`updated_at`,`deleted_at`)
    VALUES (@planId,_planName,_planDescription,'CUSTOM',_planPrice,CONCAT(_planPrice,'$'),'/assets/graphics/custom.png',_planExpiresInDays,-1,0,1,now(),null,null);
   
    #Add features to custom plan
		#Active Boxes
		INSERT INTO `justasking`.`price_plan_features` (`id`,`plan_id`,`feature_id`,`feature_value`,`created_at`,`updated_at`,`deleted_at`)
		VALUES (uuid(),@planId,@activeBoxes_featureId,_activeBoxes,now(),null,null);
		#Wordcloud
		INSERT INTO `justasking`.`price_plan_features` (`id`,`plan_id`,`feature_id`,`feature_value`,`created_at`,`updated_at`,`deleted_at`)
		VALUES (uuid(),@planId,@wordcloud_featureId,_wordcloud,now(),null,null);
		#Question Box
		INSERT INTO `justasking`.`price_plan_features` (`id`,`plan_id`,`feature_id`,`feature_value`,`created_at`,`updated_at`,`deleted_at`)
		VALUES (uuid(),@planId,@questionBox_featureId,_questionBox,now(),null,null); 
		#Answer Box
		INSERT INTO `justasking`.`price_plan_features` (`id`,`plan_id`,`feature_id`,`feature_value`,`created_at`,`updated_at`,`deleted_at`)
		VALUES (uuid(),@planId,@answerBox_featureId,_answerBox,now(),null,null); 
		#Votes Box
		INSERT INTO `justasking`.`price_plan_features` (`id`,`plan_id`,`feature_id`,`feature_value`,`created_at`,`updated_at`,`deleted_at`)
		VALUES (uuid(),@planId,@votesBox_featureId,_votesBox,now(),null,null); 
		#Toggle Responses
		INSERT INTO `justasking`.`price_plan_features` (`id`,`plan_id`,`feature_id`,`feature_value`,`created_at`,`updated_at`,`deleted_at`)
		VALUES (uuid(),@planId,@toggleResponses_featureId,_toggleResponses,now(),null,null); 
		#SMS
		INSERT INTO `justasking`.`price_plan_features` (`id`,`plan_id`,`feature_id`,`feature_value`,`created_at`,`updated_at`,`deleted_at`)
		VALUES (uuid(),@planId,@sms_featureId,_sms,now(),null,null); 
		#Custom Code
		INSERT INTO `justasking`.`price_plan_features` (`id`,`plan_id`,`feature_id`,`feature_value`,`created_at`,`updated_at`,`deleted_at`)
		VALUES (uuid(),@planId,@customCode_featureId,_customCode,now(),null,null); 
		#Delegates
		INSERT INTO `justasking`.`price_plan_features` (`id`,`plan_id`,`feature_id`,`feature_value`,`created_at`,`updated_at`,`deleted_at`)
		VALUES (uuid(),@planId,@delegates_featureId,_delegates,now(),null,null); 
		#Support
		INSERT INTO `justasking`.`price_plan_features` (`id`,`plan_id`,`feature_id`,`feature_value`,`created_at`,`updated_at`,`deleted_at`)
		VALUES (uuid(),@planId,@support_featureId,_support,now(),null,null); 
 
    #Assign a custom plan license
    INSERT INTO `justasking`.`custom_plan_licenses`
	(`id`,`account_id`,`user_id`,`plan_id`,`license_code`,`is_active`,`created_at`,`created_by`,`updated_at`,`updated_by`,`deleted_at`)
	VALUES (@customPlanLicenseId,_accountId,_userId,@planId,_licenseCode,1,now(),_createdBy,now(),null,null);
    
    SELECT _licenseCode as CustomPlanLicenseKey;
    
    COMMIT;
    
  END //

DELIMITER ;

