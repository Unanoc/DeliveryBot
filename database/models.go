package database

import (
	"fmt"
	"time"
)

// Order is a struct for order confirmation.
type Order struct {
	OrderID    int
	FirstName  string
	LastName   string
	Phone      string
	Company    string
	Address    string
	Date       time.Time
	State      int
	IsFinished bool
}

func (o Order) String() string {
	return fmt.Sprintf("Номер заказа: %d\nИмя: %s\n, Фамилия: %s\n, Телефон: %s\n, Компания: %s\n, Адрес: %s\n, Дата: %s",
		o.OrderID,
		o.FirstName,
		o.LastName,
		o.Phone,
		o.Company,
		o.Address,
		o.Date.Format("2006-01-02 15:04:05"),
	)
}
