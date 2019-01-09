package database

import (
	"fmt"
	"log"
	"time"
)

// CreateOrUpdateUserState changes the state of user in database.
func CreateOrUpdateUserState(db *DB, userID int64, state int) {
	if _, err := db.Conn.Exec(sqlInsertUserState, &userID, &state); err != nil {
		db.Conn.Exec(sqlUpdateUserState, &userID, &state)
	}
}

// GetUserStateByID selects the state of orders.
func GetUserStateByID(db *DB, userID int64) (state int) {
	db.Conn.QueryRow(sqlSelectUserStateByUserID, &userID).Scan(&state)
	return // if user's state is absent, return 0 by default
}

// CreateOrder creates an order.
func CreateOrder(db *DB, userID int64) {
	orderDate := time.Now().Format("2006-01-02 15:04:05")
	db.Conn.Exec(sqlInsertOrder, &userID, &orderDate)
}

// UpdateOrder updates order.
func UpdateOrder(db *DB, userID int64, field, value string) {
	sqlQuery := fmt.Sprintf("UPDATE orders SET %s = '%s' WHERE user_id = $1 AND order_date = (SELECT MAX(order_date) FROM orders)", field, value)

	if _, err := db.Conn.Exec(sqlQuery, &userID); err != nil {
		log.Println("1")
		log.Println(err)
		log.Println()
	}
}

// SelectOrderByID selects an order by user_id from db.
func SelectOrderByID(db *DB, userID int64) (*Order, error) {
	var order Order
	err := db.Conn.QueryRow(`
		SELECT "id", "firstname", "lastname", "phone", "company", "address", "delivery_date", "order_date"
		FROM orders
		WHERE "user_id" = $1 AND "order_date" = (SELECT MAX("order_date") FROM orders)
	`, &userID,
	).Scan(
		&order.OrderID,
		&order.FirstName,
		&order.LastName,
		&order.Phone,
		&order.Company,
		&order.Address,
		&order.DeliveryDate,
		&order.OrderDate,
	)

	if err != nil {
		log.Println("2")
		log.Println(err)
		log.Println()
		return nil, err
	}

	return &order, nil
}
