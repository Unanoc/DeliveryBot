package database

// CreateOrUpdateUserState
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

// GetUserStateByID
const sqlSelectUserStateByUserID = `
	SELECT "state" 
	FROM states 
	WHERE "user_id" = $1
`

// CreateOrder
const sqlInsertOrder = `
	INSERT INTO orders
	("user_id", "order_date")
	VALUES ($1, $2)
`

// UpdateOrder
var sqlUpdateOrder = "UPDATE orders SET %s = '%s' WHERE user_id = $1 AND order_date = (SELECT MAX(order_date) FROM orders WHERE user_id = $1)"

// SelectOrderByID
const sqlSelectOrderByID = `
	SELECT "id", "firstname", "lastname", "phone", "company", "address", "delivery_date", "order_date"
	FROM orders
	WHERE "user_id" = $1 AND "order_date" = (SELECT MAX("order_date") FROM orders WHERE "user_id" = $1)
`

// DeleteOrder
const sqlDeleteOrder = `
	DELETE 
	FROM orders 
	WHERE "user_id" = $1 AND "order_date" = (SELECT MAX("order_date") FROM orders WHERE "user_id" = $1)
`
