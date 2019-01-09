package database

import (
	"fmt"
	"time"
)

// CreateOrUpdateUserState changes the state of user in database.
func CreateOrUpdateUserState(db *DB, userID int64, state int) {
	if _, err := db.Conn.Exec(sqlInsertUserState, &userID, &state); err != nil {
		db.Conn.Exec(sqlUpdateUserState, &userID, &state)
	}
}

// GetUserStateByID selects the state of users.
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
	sqlQuery := fmt.Sprintf(sqlUpdateOrder, field, value)
	db.Conn.Exec(sqlQuery, &userID)
}

// SelectOrderByID selects an order by user_id from table.
func SelectOrderByID(db *DB, userID int64) (*Order, error) {
	var order Order
	err := db.Conn.QueryRow(sqlSelectOrderByID, &userID).Scan(
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
		return nil, err
	}

	return &order, nil
}

// DeleteOrder deletes a row from table "orders".
func DeleteOrder(db *DB, userID int64) {
	db.Conn.Exec(sqlDeleteOrder, userID)
}
