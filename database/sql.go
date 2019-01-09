package database

// CreateOrUpdateUserState //
const sqlInsertUserState = `
	INSERT INTO states
	("user_id", "state")
	VALUES ($1, $2)
`
const sqlUpdateUserState = `
	UPDATE states
	SET "state" = $2
	WHERE "user_id" = $1
`

// GetUserStateByID //
const sqlSelectUserStateByUserID = `
	SELECT "state" 
	FROM states 
	WHERE "user_id" = $1
`

// CreateOrder //
const sqlInsertOrder = `
	INSERT INTO orders
	("user_id", "firstname", "lastname", "phone", "company", "address", "delivery_date", "order_date")
	VALUES ($1, NULL, NULL, NULL, NULL, NULL, NULL, $2)
`
