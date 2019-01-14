package database

import (
	"fmt"
	"time"
)

// Order is a struct for order confirmation.
type Order struct {
	OrderID      int
	FirstName    string
	LastName     string
	Phone        string
	Company      string
	Address      string
	DeliveryDate string
	OrderDate    time.Time
}

func (o Order) String() string {
	return fmt.Sprintf(
		"Номер заказа: %d\nИмя: %s\nФамилия: %s\nТелефон: %s\n"+
			"Компания: %s\nАдрес: %s\nДата доставки: %s\nДата заказа: %s",
		o.OrderID,
		o.FirstName,
		o.LastName,
		o.Phone,
		o.Company,
		o.Address,
		o.DeliveryDate,
		o.OrderDate.Format("2006-01-02 15:04:05"),
	)
}
