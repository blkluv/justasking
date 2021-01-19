INSERT INTO app_configs (config_type, config_code, config_value, comments)
	SELECT * FROM (SELECT 'syncservice', 'PhoneNumbersThreshold', '5', 'Number of extra phone numbers we should have at all times.') AS tmp
		WHERE NOT EXISTS (
			SELECT config_type, config_code FROM justasking.app_configs WHERE config_type = 'syncservice' AND config_code = 'PhoneNumbersThreshold'
);
INSERT INTO app_configs (config_type, config_code, config_value, comments)
	SELECT * FROM (SELECT 'syncservice', 'RunPlanExpirationSync', 'true', 'Feature toggle for automatically canceling expired plans') AS tmp
		WHERE NOT EXISTS (
			SELECT config_type, config_code FROM justasking.app_configs WHERE config_type = 'syncservice' AND config_code = 'RunPlanExpirationSync'
);
INSERT INTO app_configs (config_type, config_code, config_value, comments)
	SELECT * FROM (SELECT 'syncservice', 'RunPhoneNumbersSync', 'false', 'Feature toggle for automatically syncing phone numbers from twilio') AS tmp
		WHERE NOT EXISTS (
			SELECT config_type, config_code FROM justasking.app_configs WHERE config_type = 'syncservice' AND config_code = 'RunPhoneNumbersSync'
);
INSERT INTO app_configs (config_type, config_code, config_value, comments)
	SELECT * FROM (SELECT 'twilio', 'TwilioLiveAccountSid', '{{TwilioLiveAccountSid}}', 'Twilio account Sid') AS tmp
		WHERE NOT EXISTS (
			SELECT config_type, config_code FROM justasking.app_configs WHERE config_type = 'twilio' AND config_code = 'TwilioLiveAccountSid'
);
INSERT INTO app_configs (config_type, config_code, config_value, comments)
	SELECT * FROM (SELECT 'twilio', 'TwilioTestAccountSid', '{{TwilioTestAccountSid}}', 'Twilio account Sid') AS tmp
		WHERE NOT EXISTS (
			SELECT config_type, config_code FROM justasking.app_configs WHERE config_type = 'twilio' AND config_code = 'TwilioTestAccountSid'
);
INSERT INTO app_configs (config_type, config_code, config_value, comments)
	SELECT * FROM (SELECT 'twilio', 'TwilioLiveAuthToken', '{{TwilioLiveAuthToken}}', 'Twilio auth token') AS tmp
		WHERE NOT EXISTS (
			SELECT config_type, config_code FROM justasking.app_configs WHERE config_type = 'twilio' AND config_code = 'TwilioLiveAuthToken'
);
INSERT INTO app_configs (config_type, config_code, config_value, comments)
	SELECT * FROM (SELECT 'twilio', 'TwilioTestAuthToken', '{{TwilioTestAuthToken}}', 'Twilio auth token') AS tmp
		WHERE NOT EXISTS (
			SELECT config_type, config_code FROM justasking.app_configs WHERE config_type = 'twilio' AND config_code = 'TwilioTestAuthToken'
);
INSERT INTO app_configs (config_type, config_code, config_value, comments)
	SELECT * FROM (SELECT 'twilio', 'TwilioTestModeFlag', '{{TwilioTestModeFlag}}', 'Twilio auth token') AS tmp
		WHERE NOT EXISTS (
			SELECT config_type, config_code FROM justasking.app_configs WHERE config_type = 'twilio' AND config_code = 'TwilioTestModeFlag'
);
INSERT INTO app_configs (config_type, config_code, config_value, comments)
	SELECT * FROM (SELECT 'twilio', 'TwilioApiBaseUrl', 'https://api.twilio.com/', 'Twilio api base url') AS tmp
		WHERE NOT EXISTS (
			SELECT config_type, config_code FROM justasking.app_configs WHERE config_type = 'twilio' AND config_code = 'TwilioApiBaseUrl'
);
INSERT INTO app_configs (config_type, config_code, config_value, comments)
	SELECT * FROM (SELECT 'twilio', 'TwilioAvailableNumbersUri', '2010-04-01/Accounts/{{TwilioLiveAccountSid}}/AvailablePhoneNumbers/US/Local.json', 'Twilio uri for getting available numbers') AS tmp
		WHERE NOT EXISTS (
			SELECT config_type, config_code FROM justasking.app_configs WHERE config_type = 'twilio' AND config_code = 'TwilioAvailableNumbersUri'
);
INSERT INTO app_configs (config_type, config_code, config_value, comments)
	SELECT * FROM (SELECT 'twilio', 'TwilioOurPhoneNumbersUri', '2010-04-01/Accounts/{{TwilioLiveAccountSid}}/IncomingPhoneNumbers.json', 'Twilio uri for getting numbers from our account') AS tmp
		WHERE NOT EXISTS (
			SELECT config_type, config_code FROM justasking.app_configs WHERE config_type = 'twilio' AND config_code = 'TwilioOurPhoneNumbersUri'
);
INSERT INTO app_configs (config_type, config_code, config_value, comments)
	SELECT * FROM (SELECT 'twilio', 'TwilioPurchasePhoneNumberUri', '2010-04-01/Accounts/{TwilioAccountSid}/IncomingPhoneNumbers', 'Twilio uri for purchasing a phone number') AS tmp
		WHERE NOT EXISTS (
			SELECT config_type, config_code FROM justasking.app_configs WHERE config_type = 'twilio' AND config_code = 'TwilioPurchasePhoneNumberUri'
);
INSERT INTO app_configs (config_type, config_code, config_value, comments)
	SELECT * FROM (SELECT 'twilio', 'TwilioReleasePhoneNumberUri', '/2010-04-01/Accounts/{TwilioAccountSid}/IncomingPhoneNumbers/{PhoneNumberSid}', 'Twilio uri for releasing a phone number') AS tmp
		WHERE NOT EXISTS (
			SELECT config_type, config_code FROM justasking.app_configs WHERE config_type = 'twilio' AND config_code = 'TwilioReleasePhoneNumberUri'
);
INSERT INTO app_configs (config_type, config_code, config_value, comments)
	SELECT * FROM (SELECT 'twilio', 'TwilioSmsUrl', '{{TwilioSmsUrl}}', 'endpoint that twilio will call when an sms message comes in') AS tmp
		WHERE NOT EXISTS (
			SELECT config_type, config_code FROM justasking.app_configs WHERE config_type = 'twilio' AND config_code = 'TwilioSmsUrl'
);
INSERT INTO app_configs (config_type, config_code, config_value, comments)
	SELECT * FROM (SELECT 'justasking', 'RealTimeHubBaseUrl', '{{RealTimeHubBaseUrl}}', 'Base Url for real time hub') AS tmp
		WHERE NOT EXISTS (
			SELECT config_type, config_code FROM justasking.app_configs WHERE config_type = 'justasking' AND config_code = 'RealTimeHubBaseUrl'
);
INSERT INTO app_configs (config_type, config_code, config_value, comments)
	SELECT * FROM (SELECT 'justasking', 'JustAskingWebBaseUrl', '{{JustAskingWebBaseUrl}}', 'Base Url for real time hub') AS tmp
		WHERE NOT EXISTS (
			SELECT config_type, config_code FROM justasking.app_configs WHERE config_type = 'justasking' AND config_code = 'JustAskingWebBaseUrl'
);
INSERT INTO app_configs (config_type, config_code, config_value, comments)
	SELECT * FROM (SELECT 'justasking', 'MinimumPasswordLength', '12', 'Minimum length for justasking passwords') AS tmp
		WHERE NOT EXISTS (
			SELECT config_type, config_code FROM justasking.app_configs WHERE config_type = 'justasking' AND config_code = 'MinimumPasswordLength'
);
INSERT INTO app_configs (config_type, config_code, config_value, comments)
	SELECT * FROM (SELECT 'justasking', 'MaximumPasswordLength', '32', 'Minimum length for justasking passwords') AS tmp
		WHERE NOT EXISTS (
			SELECT config_type, config_code FROM justasking.app_configs WHERE config_type = 'justasking' AND config_code = 'MaximumPasswordLength'
);
INSERT INTO app_configs (config_type, config_code, config_value, comments)
	SELECT * FROM (SELECT 'stripe', 'StripeSecretKey', '{{StripeSecretKey}}', 'Secret key for Stripe Api') AS tmp
		WHERE NOT EXISTS (
			SELECT config_type, config_code FROM justasking.app_configs WHERE config_type = 'stripe' AND config_code = 'StripeSecretKey'
);
INSERT INTO app_configs (config_type, config_code, config_value, comments)
	SELECT * FROM (SELECT 'recaptcha', 'ReCaptchaUrl', 'https://www.google.com/recaptcha/api/siteverify', 'URL for verifying recaptcha tokens') AS tmp
		WHERE NOT EXISTS (
			SELECT config_type, config_code FROM justasking.app_configs WHERE config_type = 'recaptcha' AND config_code = 'ReCaptchaUrl'
);
INSERT INTO app_configs (config_type, config_code, config_value, comments)
	SELECT * FROM (SELECT 'recaptcha', 'ReCaptchaSecretKey', '6LdToUAUAAAAADVvCh_MDvWlTcu-pw27Zp-FM1Sz', 'Secret key for calling recaptcha endpoint') AS tmp
		WHERE NOT EXISTS (
			SELECT config_type, config_code FROM justasking.app_configs WHERE config_type = 'recaptcha' AND config_code = 'ReCaptchaSecretKey'
);