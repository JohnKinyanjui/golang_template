package properties_service

const createPropertyQuery = `
	INSERT INTO properties (
		id, user_id, title, description, price, address, city, state, zip_code,
		bedrooms, bathrooms, square_feet, property_type, images
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
	) RETURNING *
`

const getPropertyQuery = `
	SELECT * FROM properties WHERE id = $1
`

const getUserPropertiesQuery = `
	SELECT * FROM properties WHERE user_id = $1 ORDER BY created_at DESC
`

const updatePropertyQuery = `
	UPDATE properties SET
		title = $2, description = $3, price = $4, address = $5, city = $6,
		state = $7, zip_code = $8, bedrooms = $9, bathrooms = $10,
		square_feet = $11, property_type = $12, status = $13, images = $14,
		updated_at = current_timestamp
	WHERE id = $1 AND user_id = $15 RETURNING *
`

const deletePropertyQuery = `
	DELETE FROM properties WHERE id = $1 AND user_id = $2
`

const getUserPropertyCountQuery = `
	SELECT COUNT(*) FROM properties WHERE user_id = $1 AND status = 'active'
`
