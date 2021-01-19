
DELETE FROM justasking.custom_plan_licenses;

/*ADD CUSTOM PLAN LICENSE*/

SELECT @user_id := (SELECT u.id FROM justasking.users u where u.email = 'arieldiaz92@gmail.com' LIMIT 1);
SELECT @account_id := (SELECT  a.id FROM justasking.accounts a where a.owner_id = @user_id AND a.is_active = 1 LIMIT 1);
SELECT @plan_id := (SELECT pp.id FROM justasking.price_plans pp where pp.name = 'PREMIUM-WEEK' LIMIT 1);
  
INSERT INTO `justasking`.`custom_plan_licenses` (`id`,`account_id`,`user_id`,`plan_id`,`license_code`,`is_active`,`created_at`,`created_by`,`updated_at`,`updated_by`,`deleted_at`)
VALUES
(uuid(), 
@account_id, 
@user_id, 
@plan_id,
'w9LxyDugHWNKbiLJeR5BXQTELMMCj0J7PpxfnY0TuV9jSXNhOi2r4DSc5WuDk5f191iYt6mJqca3fgvBPIQqh68o0zFi3XCh1T9nW4RqcOrCX1yGHfLeMqqznqSLid1sq4lvseKL3p01FHFlTGgbhMO2rIWBUNxWEnBwSNJHHospzQ3dtuzuh1gj05FzpEJwXy0YM5EsKytM0dZiSNgOmjyYKYEay78plZl8S4Ll8yhVI36AhQ5hcMnSKQNJbSEC',
1,NOW(),'ummmm',null,null,null);

SELECT * FROM justasking.custom_plan_licenses;

