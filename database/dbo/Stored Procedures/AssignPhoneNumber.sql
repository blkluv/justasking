DROP PROCEDURE IF EXISTS AssignPhoneNumber;

DELIMITER //

CREATE procedure AssignPhoneNumber (boxId varchar(100), updatedBy varchar(50))

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
		UPDATE base_box SET is_live = 1, updated_at = CURRENT_TIMESTAMP, updated_by = updatedBy WHERE id = boxId;
		
        IF EXISTS (SELECT * FROM base_box_phone_numbers WHERE base_box_id = boxId AND is_active = 1) THEN SELECT ''; #do nothing in this case
        ELSE
			IF EXISTS (SELECT updated_at FROM base_box_phone_numbers WHERE is_active = 0 AND base_box_id = boxId AND updated_at >= CURRENT_TIMESTAMP - INTERVAL 2 MINUTE) THEN UPDATE base_box_phone_numbers SET is_active = 1, updated_at = CURRENT_TIMESTAMP WHERE base_box_id = boxId AND updated_at >= CURRENT_TIMESTAMP - INTERVAL 2 MINUTE ORDER BY created_at DESC LIMIT 1;
			ELSE INSERT INTO base_box_phone_numbers (id, base_box_id, phone_number_id, is_active, created_at, created_by, updated_at) 
								SELECT UUID(), boxId, (SELECT id FROM phone_numbers AS p 
														WHERE p.is_active = 1 AND NOT EXISTS (SELECT * FROM justasking.base_box_phone_numbers WHERE phone_number_id = p.id AND (is_active = 1 OR updated_at >= CURRENT_TIMESTAMP - INTERVAL 2 MINUTE))
								LIMIT 1), 1, CURRENT_TIMESTAMP, updatedBy, CURRENT_TIMESTAMP;
			END IF;
		END IF;
    COMMIT;
    
  END //

DELIMITER ;

