package database

const (
	// CreateOrUpdateUser
	sqlInsertUser = `
	INSERT INTO users
	("id", "firstname", "lastname", "age", "sex")
	VALUES ($1, $2, $3, $4, $5)
`
	sqlUpdateUser = `
	UPDATE users
	SET "firstname" = $2, 
		"lastname" = $3,
		"age" = $4, 
		"sex" = $5
	WHERE "id" = $1
`

	// CreateOrUpdateUserState
	sqlInsertUserState = `
	INSERT INTO states
	("user_id", "state")
	VALUES ($1, $2)
`
	sqlUpdateUserState = `
	UPDATE states
	SET "state" = $2
	WHERE "user_id" = $1
`

	// GetUserStateByID
	sqlSelectUserStateByUserID = `
	SELECT "state" 
	FROM states 
	WHERE "user_id" = $1
`

	// CreateOrder
	sqlInsertOrder = `
	INSERT INTO orders
	("user_id", "order_date")
	VALUES ($1, $2)
`

	// SelectOrderByID
	sqlSelectOrderByID = `
	SELECT "id", "firstname", "lastname", "phone", "company", "address", "delivery_date", "order_date"
	FROM orders
	WHERE "user_id" = $1 AND "order_date" = (SELECT MAX("order_date") FROM orders WHERE "user_id" = $1)
`

	// DeleteOrder
	sqlDeleteOrder = `
	DELETE 
	FROM orders 
	WHERE "user_id" = $1 AND "order_date" = (SELECT MAX("order_date") FROM orders WHERE "user_id" = $1)
`
)

// UpdateOrder
var sqlUpdateOrder = "UPDATE orders SET %s = '%s' WHERE user_id = $1 AND order_date = (SELECT MAX(order_date) FROM orders WHERE user_id = $1)"
