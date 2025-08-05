package subscriptions_service

const getUserSubscriptionQuery = `
	SELECT * FROM subscriptions WHERE user_id = $1 AND status = 'active' ORDER BY created_at DESC LIMIT 1
`

const createSubscriptionQuery = `
	INSERT INTO subscriptions (user_id, tier, max_properties, features) VALUES ($1, $2, $3, $4)
`

const updateSubscriptionQuery = `
	UPDATE subscriptions SET tier = $2, max_properties = $3, features = $4, updated_at = current_timestamp WHERE id = $1 RETURNING *
`
