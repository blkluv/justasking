SELECT 
bb.code,
bb.is_live,
bt.name,
a.name,
u.email,
a.created_at as 'account created on',
bb.created_at as 'box created on'
FROM justasking.base_box bb
JOIN justasking.accounts a ON a.id = bb.account_id
JOIN justasking.box_type bt ON bb.box_type = bt.id
JOIN justasking.users u ON a.owner_id = u.id
WHERE 1=1 
Order by bb.created_at desc
