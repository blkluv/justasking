SELECT
	pp.name,
    f.name,
    ppf.feature_value
FROM justasking.price_plan_features ppf
JOIN  justasking.features f ON  ppf.feature_id = f.id
JOIN  justasking.price_plans pp ON ppf.plan_id = pp.id
