package model

import "time"

type Order struct {
	OrderId     int        `db:"order_id"`
	CustomerId  int        `db:"id_customer"`
	Total       int        `db:"total"`
	TableNumber int        `db:"id_meja"`
	Payment     string     `db:"payment"`
	Employee    string     `db:"id_karyawan"`
	OrderTime   *time.Time `db:"time_order"`
	PaymentTime *time.Time `db:"payment_time"`
}

func NewOrder(orderId, customerId, total, tableNumber int, payment, employee string, timeOrder, paymentTime *time.Time) Order {
	return Order{
		OrderId:     orderId,
		CustomerId:  customerId,
		Total:       total,
		TableNumber: tableNumber,
		Payment:     payment,
		Employee:    employee,
		OrderTime:   timeOrder,
		PaymentTime: paymentTime,
	}
}
