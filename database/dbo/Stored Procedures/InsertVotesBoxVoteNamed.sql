DROP PROCEDURE IF EXISTS InsertVotesBoxVoteNamed;

DELIMITER //

CREATE procedure InsertVotesBoxVoteNamed (answerId varchar(100), questionId varchar(100), createdBy varchar(100))

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
    
    INSERT INTO votes_box_question_answers_votes (answer_id, question_id, created_at, created_by, updated_at, updated_by, deleted_at)
					SELECT * FROM (SELECT answerId as answer_id, questionId as question_id, current_timestamp as created_at, createdBy as created_by, null as updated_at, null as updated_by, null as deleted_at) AS tmp
					WHERE NOT EXISTS (SELECT answer_id FROM votes_box_question_answers_votes WHERE question_id = questionId AND created_by = createdBy);
                    
    SELECT ROW_COUNT() as records_inserted;
    
  END //

DELIMITER ;

