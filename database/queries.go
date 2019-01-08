package database

import "log"

// CreateOrUpdateUserState sets the state of user in database.
func CreateOrUpdateUserState(db *DB, userID int64, state int) {
	if existsCurrentOrderByID(db, userID) {
		_, err := db.Conn.Exec(`
			UPDATE orders
			SET "state" = $2
			WHERE "user_id" = $1 AND NOT "is_finished"
		`, &userID, &state)
		if err != nil {
			log.Println()
			log.Println(err)
			log.Println()
		}
	} else {
		_, err := db.Conn.Exec(`
			INSERT INTO orders
			("user_id", "firstname", "lastname", "phone", "company", "address", "date", "state")
			VALUES ($1, NULL, NULL, NULL, NULL, NULL, NULL, $2)
		`, &userID, &state)
		if err != nil {
			log.Println()
			log.Println(err)
			log.Println()
		}
	}
}

// GetUserStateByID selects the state of orders.
func GetUserStateByID(db *DB, userID int64) int {
	var state int
	err := db.Conn.QueryRow(`
		SELECT "state" 
		FROM orders 
		WHERE "user_id" = $1 AND NOT "is_finished"
	`, &userID).Scan(&state)

	if err != nil {
		return 0 // StateNull
	}

	return state
}

// SelectFinishedOrderByID selects record of completed order.
func SelectFinishedOrderByID(db *DB, userID int64) (*Order, error) {
	var order Order
	err := db.Conn.QueryRow(`
		SELECT "id", "firstname", "lastname", "phone", "company", "address", "date"
		FROM orders
		WHERE "user_id" = $1 AND "is_finished" AND "date" = (SELECT MAX("date") FROM orders)
	`, &userID,
	).Scan(
		&order.OrderID,
		&order.FirstName,
		&order.LastName,
		&order.Phone,
		&order.Company,
		&order.Address,
		&order.Date,
	)

	if err != nil {
		log.Println()
		log.Println(err)
		log.Println()
		return nil, err
	}

	return &order, nil
}

// SetFinishFlagOrder sets field "is_finished" in order.
func SetFinishFlagOrder(db *DB, userID int64) {
	db.Conn.Exec(`
		UPDATE orders
		SET "is_finished" = true
		WHERE "user_id" = $1 AND NOT "is_finished" AND "date" = (SELECT MAX("date") FROM orders)
	`, &userID)
}

func existsCurrentOrderByID(db *DB, userID int64) bool {
	var id int
	err := db.Conn.QueryRow(`
		SELECT "id"
		FROM orders
		WHERE "user_id" = $1 AND NOT "is_finished"
	`, &userID).Scan(&id)

	if err != nil {
		log.Println()
		log.Println(err)
		log.Println()
		return false
	}

	return true
}
